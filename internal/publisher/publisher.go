package publisher

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"math/rand"
	"telem.kmani/internal/models"
	"time"
)

func RunPublisherLoop(client mqtt.Client, duration time.Duration, frequency int, deviceID string) {
	interval := time.Second / time.Duration(frequency)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	done := time.After(duration)
	numMessages := 0

	for {
		select {
		case <-ticker.C:
			telemetry := models.GenerateTelemetryRandom(deviceID)
			payload, _ := json.Marshal(telemetry)
			token := client.Publish("telemetry/"+deviceID, 0, false, payload)
			token.Wait()
			log.Printf("Published Telemetry for %s", deviceID)

			if rand.Intn(20) == 0 { // ~5% chance
				event := models.GenerateEventRandom(deviceID)
				eventPayload, _ := json.Marshal(event)
				token := client.Publish("events/"+deviceID, 0, false, eventPayload)
				token.Wait()
				log.Printf("Published Event for %s", deviceID)
			}
			if numMessages%30 == 0 { // every ~30 telemetry ticks
				config := models.GenerateDeviceConfigRandom(deviceID)
				configPayload, _ := json.Marshal(config)
				token := client.Publish("device-config/"+deviceID, 0, false, configPayload)
				token.Wait()
				log.Printf("Published DeviceConfig for %s", deviceID)
			}
			numMessages++

		case <-done:
			log.Printf("Device %s done publishing!", deviceID)
			return
		}
	}
}
