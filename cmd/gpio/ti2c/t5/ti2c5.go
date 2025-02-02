package main

import (
	"fmt"
	"log"
	"time"

	"github.com/stianeustrup/go-rpio/v4"
	"golang.org/x/exp/io/i2c"
)

const (
	RELAY_HAT_I2C_ADDR = 0x20 // Adresa I2C relé HAT
	RELAY_COUNT        = 8    // Počet relé
)

func main() {
	// Otvorenie I2C zbernice
	dev, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, RELAY_HAT_I2C_ADDR)
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	// Inicializácia RPIO pre prístup k GPIO pinom
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()

	// Vytvorenie slice pre uloženie stavu relé
	relays := make([]bool, RELAY_COUNT)

	// Hlavná slučka programu
	for {
		// Prechod cez všetky relé
		for i := 0; i < RELAY_COUNT; i++ {
			// Zmena stavu relé
			relays[i] = !relays[i]

			// Zápis stavu relé do I2C zbernice
			err := writeRelayState(dev, relays)
			if err != nil {
				log.Fatal(err)
			}

			// Výpis stavu relé do konzoly
			fmt.Printf("Relay %d: %t\n", i+1, relays[i])

			// Čakanie 1 sekundu
			time.Sleep(time.Second)
		}
	}
}

// Funkcia pre zápis stavu relé do I2C zbernice
func writeRelayState(dev *i2c.Device, relays []bool) error {
	// Vytvorenie bajtu pre zápis do I2C zbernice
	var data byte
	for i := 0; i < RELAY_COUNT; i++ {
		if relays[i] {
			data |= 1 << i
		}
	}

	// Zápis bajtu do I2C zbernice
	_, err := dev.Write([]byte{data})
	return err
}
