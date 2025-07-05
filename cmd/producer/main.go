package main

import (
	"sync"
	"telem.kmani/internal/publisher"
	"telem.kmani/internal/utils"
	"time"
)

func main() {
	client := utils.ConnectMQTT("tcp://localhost:1883", "unique-string")
	devices := [...]string{"device1", "device2", "device3", "device4", "device5"}
	var wg sync.WaitGroup
	for _, device := range devices {
		wg.Add(1)
		go func(device string) {
			defer wg.Done()
			publisher.RunPublisherLoop(client, time.Minute*2, 100, device)
		}(device)
	}
	wg.Wait()
}
