package main

import (
	"fmt"
	"log"
	"time"

	"github.com/stianeves/go-i2c"
)

func main() {
	// Otvorenie I2C zbernice
	bus, err := i2c.New(1, 0x20) // 1 je číslo zbernice a 0x20 je adresa zariadenia
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	// Zapnutie všetkých relé
	err = bus.WriteReg(0x00, 0xFF) // 0xFF je maska pre zapnutie všetkých relé
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Všetky relé zapnuté")

	time.Sleep(1 * time.Second)

	// Vypnutie všetkých relé
	err = bus.WriteReg(0x00, 0x00) // 0x00 je maska pre vypnutie všetkých relé
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Všetky relé vypnuté")
}
