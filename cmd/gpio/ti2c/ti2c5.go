package main

import (
	"fmt"
	"log"

	"github.com/d2r2/go-i2c"
)

func main() {
	// Otvorenie I2C zbernice
	i2c, err := i2c.New(1, 0x20) // 1 je číslo zbernice a 0x20 je adresa zariadenia
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	// Zapnutie relé 1
	err = i2c.WriteReg(0x00, []byte{0x01}) // 0x00 je register a 0x01 je hodnota pre zapnutie relé 1
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Relé 1 zapnuté")

	// Vypnutie relé 1
	err = i2c.WriteReg(0x00, []byte{0x00}) // 0x00 je register a 0x00 je hodnota pre vypnutie relé 1
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Relé 1 vypnuté")

	// Zapnutie relé 2
	err = i2c.WriteReg(0x01, []byte{0x01}) // 0x01 je register a 0x01 je hodnota pre zapnutie relé 2
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Relé 2 zapnuté")

	// Vypnutie relé 2
	err = i2c.WriteReg(0x01, []byte{0x00}) // 0x01 je register a 0x00 je hodnota pre vypnutie relé 2
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Relé 2 vypnuté")

	// ... a tak ďalej pre ostatné relé
}
