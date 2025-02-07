package main

import (
	"fmt"
	"log"
	"time"

	"github.com/d2r2/go-i2c"
)

func main() {
	// Otvorenie I2C zbernice
	i2c, err := i2c.NewI2C(0x26, 1) //  0x26 je adresa relay8 karty a 1 je číslo zbernice a
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	// Zápis masky do registra
	err = i2c.WriteRegU8(0x01, 0xFF) //rele 2
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Všetky relé zapnuté")

	// Počkajte 5 sekúnd
	time.Sleep(5 * time.Second)

	// Zápis masky do registra
	err = i2c.WriteRegU8(0x01, 0x00)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Všetky relé vypnuté")
}
