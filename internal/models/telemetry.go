package models

import (
	"math/rand"
	"time"
)

type Sensors struct {
	Temperature float64 `json:"temperature"`
	Pressure    float64 `json:"pressure"`
	Humidity    float64 `json:"humidity"`
	Vibration   float64 `json:"vibration"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Telemetry struct {
	DeviceID   string   `json:"deviceId"`
	DeviceName string   `json:"deviceName"`
	Mode       string   `json:"mode"` // e.g., "IDLE", "ACTIVE", "SLEEP"
	Fault      bool     `json:"fault"`
	Sensors    Sensors  `json:"sensors"`
	Location   Location `json:"location"`
	Battery    float64  `json:"battery"` // %
	CreatedAt  string   `json:"createdAt"`
	UpdatedAt  string   `json:"updatedAt"`
}

func (t *Telemetry) GetDeviceId() string {
	return t.DeviceID
}
func (t *Telemetry) GetDeviceName() string {
	return t.DeviceName
}

func GenerateTelemetryRandom(deviceID string) Telemetry {
	deviceName := "device-" + deviceID
	sensors := Sensors{
		Temperature: rand.Float64()*50 + 10, // 10°C to 60°C
		Pressure:    rand.Float64()*5 + 1,   // 1 bar to 6 bar
		Humidity:    rand.Float64()*50 + 20, // 20% to 70%
		Vibration:   rand.Float64() * 5,     // arbitrary units
	}
	location := Location{
		Lat: 37.7 + rand.Float64()*0.1,
		Lon: -122.5 + rand.Float64()*0.1,
	}
	modes := []string{"IDLE", "ACTIVE", "SLEEP"}
	mode := modes[rand.Intn(len(modes))]
	telemetry := Telemetry{
		DeviceID:   deviceID,
		DeviceName: deviceName,
		Mode:       mode,
		Fault:      rand.Intn(10) == 0, // 10% chance of fault
		Sensors:    sensors,
		Location:   location,
		Battery:    rand.Float64() * 100, // %
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}

	return telemetry
}
