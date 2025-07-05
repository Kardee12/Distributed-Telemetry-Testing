package models

import (
	"math/rand"
	"time"
)

type Event struct {
	DeviceID    string `json:"deviceId"`
	EventType   string `json:"eventType"` // e.g., "FAULT", "REBOOT", "UPDATE"
	Description string `json:"description"`
	Severity    string `json:"severity"` // e.g., "INFO", "WARNING", "CRITICAL"
	Timestamp   string `json:"timestamp"`
}

func GenerateEventRandom(deviceID string) Event {
	eventTypes := []string{"FAULT", "REBOOT", "UPDATE"}
	severities := []string{"INFO", "WARNING", "CRITICAL"}

	event := Event{
		DeviceID:    deviceID,
		EventType:   eventTypes[rand.Intn(len(eventTypes))],
		Description: "Simulated event for " + deviceID,
		Severity:    severities[rand.Intn(len(severities))],
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	return event
}
