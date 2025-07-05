package models

import (
	"math/rand"
	"time"
)

type DeviceConfig struct {
	DeviceID     string   `json:"deviceId"`
	DeviceName   string   `json:"deviceName"`
	Model        string   `json:"model"`
	InstallDate  string   `json:"installDate"`
	Firmware     string   `json:"firmware"`
	Tags         []string `json:"tags"` //["indoor", "lab"]
	LastModified string   `json:"lastModified"`
}

func GenerateDeviceConfigRandom(deviceID string) DeviceConfig {
	models := []string{"SENSOR_X", "SENSOR_Y", "SENSOR_PRO"}
	tags := [][]string{
		{"indoor", "lab"},
		{"outdoor", "field"},
		{"test", "experimental"},
	}

	config := DeviceConfig{
		DeviceID:     deviceID,
		DeviceName:   "device-" + deviceID,
		Model:        models[rand.Intn(len(models))],
		InstallDate:  time.Now().AddDate(-rand.Intn(3), 0, 0).Format(time.RFC3339), // up to 3 years old
		Firmware:     "v" + string(rune(rand.Intn(5)+'1')) + ".0",
		Tags:         tags[rand.Intn(len(tags))],
		LastModified: time.Now().Format(time.RFC3339),
	}

	return config
}
