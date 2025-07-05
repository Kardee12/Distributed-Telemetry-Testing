package subscriber

import (
	"context"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strings"
	"telem.kmani/internal/models"
	"telem.kmani/internal/storage"
)

func Subscribe(client mqtt.Client) {
	topics := []string{
		"telemetry/+",
		"events/+",
		"device-config/+",
	}
	for _, topic := range topics {
		if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}
		log.Println("Subscribed to topic:", topic)
	}
}

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message on topic %s\n", msg.Topic())

	ctx := context.Background()
	switch {
	case strings.HasPrefix(msg.Topic(), "telemetry/"):
		var telemetry models.Telemetry
		if err := json.Unmarshal(msg.Payload(), &telemetry); err != nil {
			log.Println("Failed to decode telemetry:", err)
			return
		}
		if err := storage.StoreDoc(ctx, telemetry); err != nil {
			log.Println("Failed to store telemetry:", err)
		} else {
			log.Println("Indexed telemetry.")
		}
	case strings.HasPrefix(msg.Topic(), "events/"):
		var event models.Event
		if err := json.Unmarshal(msg.Payload(), &event); err != nil {
			log.Println("Failed to decode event:", err)
			return
		}
		if err := storage.StoreDoc(ctx, event); err != nil {
			log.Println("Failed to store event:", err)
		} else {
			log.Println("Indexed event.")
		}
	case strings.HasPrefix(msg.Topic(), "device-config/"):
		var config models.DeviceConfig
		if err := json.Unmarshal(msg.Payload(), &config); err != nil {
			log.Println("Failed to decode device config:", err)
			return
		}
		if err := storage.StoreDoc(ctx, config); err != nil {
			log.Println("Failed to store device config:", err)
		} else {
			log.Println("Indexed device config.")
		}

	default:
		log.Println("Unknown topic:", msg.Topic())
	}
}
