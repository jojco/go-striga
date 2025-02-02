package main

import (
	"fmt"
	"log"

	"github.com/d2r2/go-i2c"
)

func main() {
	// Otvorenie I2C zbernice
	i2c, err := i2c.NewI2C(1, 0x27) // 1 je číslo zbernice a 0x20 je adresa zariadenia
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	// Zápis masky do registra
	err = i2c.WriteRegU8(0x00, 0x08)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Všetky relé zapnuté")

	// Zápis masky do registra
	err = i2c.WriteRegU8(0x00, 0x00)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Všetky relé vypnuté")
}
