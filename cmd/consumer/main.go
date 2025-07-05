package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"telem.kmani/internal/subscriber"
	"telem.kmani/internal/utils"
)

func main() {
	// 1. Connect to MQTT broker
	utils.InitClient()
	client := utils.ConnectMQTT("tcp://localhost:1883", "consumer")
	subscriber.Subscribe(client)

	// 3. Block forever so your subscriber doesn't exit immediately
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	log.Println("Consumer running... Press Ctrl+C to exit.")
	<-sigs

	client.Disconnect(250)
	log.Println("Consumer stopped.")
}
