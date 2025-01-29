package main

import (
	"fmt"
	"log"
	"time"

	"github.com/d2r2/go-i2c"
)

func main() {
	// Otvorí I2C zariadenie na adrese 0x20 (alebo podľa potreby uprav adresu)
	i2c, err := i2c.NewI2C(0x20, 1) // 0x20 je I2C adresa, 1 je číslo zariadenia na Raspberry Pi (i2c-1)
	if err != nil {
		log.Fatalf("Chyba pri otváraní I2C: %v", err)
	}
	defer i2c.Close()

	// Predpokladajme, že relé sú pripojené k 8 bitovým výstupom (0-7) na PCF8574
	// Ak sa relé zapínajú/vypínajú na základe výstupných bitov, môžeme nastaviť tieto bity.
	var relayState byte = 0x00 // Všetky relé vypnuté

	// Cyklus na zapínanie a vypínanie relé
	for {
		// Zapneme všetky relé
		relayState = 0xFF
		err := i2c.WriteByte(relayState)
		if err != nil {
			log.Fatalf("Chyba pri zapísaní na I2C: %v", err)
		}
		fmt.Println("Relé zapnuté")
		time.Sleep(1 * time.Second)

		// Vypneme všetky relé
		relayState = 0x00
		err = i2c.WriteByte(relayState)
		if err != nil {
			log.Fatalf("Chyba pri zapísaní na I2C: %v", err)
		}
		fmt.Println("Relé vypnuté")
		time.Sleep(1 * time.Second)
	}
}
