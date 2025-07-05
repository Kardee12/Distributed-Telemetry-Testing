package utils

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func ConnectMQTT(brokerURL, clientID string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(brokerURL)
	opts.SetClientID(clientID)
	clientInstance := mqtt.NewClient(opts)
	if token := clientInstance.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	log.Println("Connected to broker", brokerURL)
	return clientInstance
}
