package main

import (
	"fmt"
	"log"
	"time"

	"github.com/d2r2/go-i2c"
)

func main() {
	// Otvorenie I2C zariadenia (bus 1 je pre Raspberry Pi)
	i2c, err := i2c.NewI2C(0x20, 1) // 0x20 je bežná adresa pre I2C relé moduly
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	// Zapnúť všetky relé (bitová maska 0xFF znamená všetky relé ON)
	err = i2c.WriteRegByte(0x00, 0xFF) // Adresa 0x00 je pre ovládanie všetkých relé
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Všetky relé sú zapnuté!")

	// Čakanie 2 sekundy
	time.Sleep(2 * time.Second)

	// Vypnúť všetky relé
	err = i2c.WriteRegByte(0x00, 0x00) // Vypnutie všetkých relé (0x00)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Všetky relé sú vypnuté!")

	// Čakanie pred ukončením programu
	time.Sleep(2 * time.Second)
}
