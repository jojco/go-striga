package pkg3

// Zoznam teplomerov a umiestnenie je v súbore configw1.json

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	//"strconv"
	"strings"
	"time"
)

const w1DevicesDir = "/sys/bus/w1/devices/"

type W1Device struct {
	SensorID string `json:"sensorid"`
	Path     string `json:"path"`
	Location string `json:"location"` // umiestnenie teplomera(UK,TUV,VT a pod.)
}

type Config struct {
	Devices []W1Device `json:"devices"`
}

// ************************************************************
// vytvorenie databázy teplomerov na rýchly prístup
// ************************************************************
func InitTeplomery() {
	//Načítanie zoznamu teplomerov zo súboru json do databázy
	config, err := loadConfig("config_w1.json")
	if err != nil {
		fmt.Println("Chyba pri načítaní konfigurácie:", err)
		return
	}
	// Tlač obsahu súboru config_w1.json
	for _, device := range config.Devices {
		fmt.Println("obsah súboru config_w1.json") //
		fmt.Println("SensorID:", device.SensorID)
		fmt.Println("Path:", device.Path)
		fmt.Println("Location:", device.Location)
		fmt.Println("---") // Oddeľovač pre lepšiu čitateľnosť
	}
	// Otvorenie alebo vytvorenie databázy SQLite3
	db, err := sql.Open("sqlite3", "./config_w1.db")
	if err != nil {
		log.Fatalf("Chyba pri otvorení databázy: %v", err)
	}
	//defer db.Close() // zabezpečí zatvorenie db po ukončení funkcie ale ja chcem aby bola prístupná počas chodu programu
	// Vytvorenie tabuľky, ak neexistuje
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS config_w1 (
					sensorid TEXT,
					path TEXT,
					location TEXT
			)
	`)

	if err != nil {
		log.Fatalf("Chyba pri vytváraní tabuľky: %v", err)
	}
	// Vloženie dát z JSON do databázy
	for _, device := range config.Devices {
		_, err = db.Exec(
			"INSERT INTO config_w1 (sensorid, path, location) VALUES (?, ?, ?)",
			device.SensorID, device.Path, device.Location,
		)
		if err != nil {
			log.Printf("Chyba pri vkladaní dát: %v", err)
		}
	}

	fmt.Println("Dáta úspešne uložené do databázy.")

}

func loadConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

// ********************************************************************
// readTemperature reads the temperature from the specified w1 device
// ********************************************************************

func ReadTemperature(location string) (string, float32, time.Time, error) {

	return "28-44444444", 55.0, time.Now(), nil
}

// ***************************************************************
// funkciu môžeš zavolať na vyhľadanie pripojených teplomerov
// určené pre servisný mód
// ***************************************************************
func NajdiTeplomer() {
	//List all w1 devices
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
		fmt.Println("Device ID:", device.SensorID)
		fmt.Println()
	}
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
				SensorID: name,
				Path:     devicePath,
			})
		}
	}
	return w1Devices, nil

}
