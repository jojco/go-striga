// modul na ovládanie relé cez i2c
// i2c_brcmstb            12288  0
// i2c_bcm2835            16384  0
// i2c_dev                16384  0
// sudo apt-get install libbcm2835-dev
// go get -u periph.io/x/periph

package main

import (
	"fmt"
	//"go/format"
	"log"
	"time"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

// Nastaví adresu I2C zariadenia na 0x26 (alebo podľa potreby uprav adresu)
const i2cAddress = 0x26 // Adresa I2C zariadenia nastavená na 0x26 HEXA t.j. 38 DEC
const bus = 0x01

func main() {
	// Inicializácia host systému (Raspberry Pi)
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Pripojenie na I2C bus (na Raspberry Pi je zvyčajne /dev/i2c-1)
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	// Vytvorenie I2C zariadenia na základe adresy
	device := i2c.Dev{Bus: bus, Addr: i2cAddress}

	// Vypíšeme info o pripojení
	fmt.Println("I2C zariadenie pripojené na adrese", i2cAddress)

	for i := 0; i < 50; i++ {

		// Zapni relé
		if err := toggleRelay(&device, true); err != nil { //jojco err má svoje císlo chyby a ak nie je chyba tak je nil
			log.Fatal(err)
		}

		// Počkajte 5 sekúnd
		time.Sleep(5 * time.Second)

		// Vypni relé
		if err := toggleRelay(&device, false); err != nil {
			log.Fatal(err)
		}

		// Počkajte 5 sekúnd
		time.Sleep(5 * time.Second)
	}
}

func toggleRelay(device *i2c.Dev, state bool) error {
	var value byte
	if state {
		value = 0xFF // Predpokladáme, že 0x01 zapne relé
	} else {
		value = 0x00 // Predpokladáme, že 0x00 vypne relé
	}

	// Zapíšeme hodnotu do zariadenia na určený register
	// Tento príklad používá register 0x01
	_, err := device.Write([]byte{0x01, value})
	if err != nil {
		return fmt.Errorf("chyba pri zapise do I2C zariadenia: %v", err)
	}

	fmt.Printf("Relé je teraz na %v\n", state)
	return nil
}
