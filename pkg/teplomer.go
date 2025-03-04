package teplota

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const w1DevicesDir = "/sys/bus/w1/devices/"

var Tpole [16]float32

func teplomer() {
	// List all w1 devices
	devices, err := listW1Devices()
	if err != nil {
		log.Fatal("Error listing w1 devices:", err)
	}

	// Print information about each device
	if len(devices) == 0 {
		log.Fatal("No w1 devices found")
	}

	fmt.Println("Found w1 devices:")
	for _, device := range devices {
		fmt.Println("Device path:", device.Path)
		fmt.Println("Device ID:", device.ID)
		fmt.Println()
	}

	// Choose a specific w1 device to read temperature from
	sensorID := devices[1].ID // You can choose the first device for simplicity, or prompt the user to select one
	fmt.Println("Reading temperature from:", sensorID)

	// Read temperature from the selected device continuously
	for {
		temp, err := readTemperature(sensorID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Temperature from %s: %.2f°C\n", sensorID, temp)

		time.Sleep(2 * time.Second) // Read temperature every 2 seconds
	}

}

// W1Device represents a w1 device with its ID and full path.
type W1Device struct {
	ID   string // Device ID
	Path string // Full path of the device directory
}

// listW1Devices lists all w1 devices available in the /sys/bus/w1/devices/ directory.
func listW1Devices() ([]W1Device, error) {
	dir, err := os.Open(w1DevicesDir)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	deviceNames, err := dir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	var w1Devices []W1Device
	for _, name := range deviceNames {
		// Check if the device name starts with "28-" (indicating a sensor device)
		if strings.HasPrefix(name, "28-") {
			devicePath := filepath.Join(w1DevicesDir, name)
			w1Devices = append(w1Devices, W1Device{
				ID:   name,
				Path: devicePath,
			})
		}
	}
	return w1Devices, nil
}

// readTemperature reads the temperature from the specified w1 device.
func readTemperature(sensorID string) (float64, error) {
	filename := filepath.Join(w1DevicesDir, sensorID, "temperature")
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	tempStr := string(data)
	tempValue, err := strconv.ParseFloat(strings.TrimSpace(tempStr), 64)
	if err != nil {
		return 0, err
	}

	return tempValue / 1000.0, nil
}
