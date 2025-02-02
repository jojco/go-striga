package main

import (
	"fmt"
	"log"
	"time"

	"github.com/d2r2/go-i2c"
)

func main() {
	// Otvorenie I2C zbernice
	i2c, err := i2c.NewI2C(1, 0x20) // 1 je číslo zbernice a 0x20 je adresa zariadenia
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	// Zapnutie relé č. 4
	err = i2c.WriteReg(0x00, []byte{0x08}) // 0x08 je maska pre zapnutie relé č. 4 (0b00001000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Relé č. 4 zapnuté")

	time.Sleep(1 * time.Second) // Relé zostane zapnuté 1 sekundu

	// Vypnutie relé č. 4
	err = i2c.WriteReg(0x00, []byte{0x00}) // 0x00 je maska pre vypnutie všetkých relé
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Relé č. 4 vypnuté")
}
