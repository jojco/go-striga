package main

import (
	"fmt"
	"log"
	"time"

	"github.com/d2r2/go-i2c"
)

const (
	// I2C adresa 8-relé HAT
	relayHatAddress = 0x26 // Prednastavená adresa pre tento modul
	// Register, do ktorého zapisujeme na ovládanie relé
	relayRegister = 0x00
)

func main() {
	// Otvor I2C zariadenie (0x1 pre /dev/i2c-1 na Raspberry Pi)
	dev, err := i2c.NewI2C(relayHatAddress, 1) // 1 je číslo I2C busu (na Raspberry Pi je to zvyčajne 1)
	if err != nil {
		log.Fatalf("Chyba pri otváraní zariadenia: %v\n", err)
	}
	defer dev.Close()

	// Vytvoríme buffer na uchovávanie stavu relé (8 bitov - pre 8 relé)
	// Počiatočne budú všetky relé vypnuté (0x00 znamená všetky relé vypnuté)
	relayState := byte(0x00)

	// Striedavé zapínanie a vypínanie všetkých relé
	for {
		// Zapni všetky relé
		relayState = byte(0xFF)
		err = setRelayState(dev, relayState)
		if err != nil {
			log.Fatalf("Chyba pri zapínaní relé: %v\n", err)
		}
		fmt.Println("Všetky relé zapnuté")

		// Počkajte 2 sekundy
		time.Sleep(2 * time.Second)

		// Vypni všetky relé
		relayState = byte(0x00)
		err = setRelayState(dev, relayState)
		if err != nil {
			log.Fatalf("Chyba pri vypínaní relé: %v\n", err)
		}
		fmt.Println("Všetky relé vypnuté")

		// Počkajte 2 sekundy
		time.Sleep(2 * time.Second)
	}
}

// setRelayState nastaví stav všetkých relé na zariadení cez I2C
// relayState je 1 byte, kde každý bit reprezentuje jedno relé (1 = zapnuté, 0 = vypnuté)
func setRelayState(dev *i2c.I2C, state byte) error {
	// Zapíš stav do zariadenia (I2C zápis)
	err := dev.WriteRegU8(relayRegister, state)
	if err != nil {
		return fmt.Errorf("chyba pri zápise do I2C: %v", err)
	}
	return nil
}
