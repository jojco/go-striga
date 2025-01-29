// modul na ovládanie relé cez i2c
// i2c_brcmstb            12288  0
// i2c_bcm2835            16384  0
// i2c_dev                16384  0
// sudo apt-get install libbcm2835-dev
// go get -u periph.io/x/periph

package main

import (
	"fmt"
	"log"
	"time"
	
	"periph.io/x/periph/v3/host"
	"periph.io/x/periph/v3/devices/i2c"
)

const (
	i2cAddr = 0x20 // Adresa zariadenia (napr. PCF8574)
)

// Funkcia na nastavenie stavu relé cez I2C
func setRelayState(dev *i2c.Dev, state byte) error {
	_, err := dev.Write([]byte{state})
	return err
}

func main() {
	// Inicializácia periph knižnice
	if _, err := host.Init(); err != nil {
		log.Fatalf("Chyba pri inicializácii periph knižnice: %v", err)
	}

	// Otvorenie I2C busu, obvykle bus 1 na Raspberry Pi
	bus, err := i2c.New("/dev/i2c-1") // Závisí od vášho systému, môže byť aj "/dev/i2c-0"
	if err != nil {
		log.Fatalf("Chyba pri otvorení I2C busu: %v", err)
	}
	defer bus.Close()

	// Pripojenie k I2C zariadeniu (PCF8574 alebo podobné)
	dev := &i2c.Dev{Bus: bus, Addr: i2cAddr}

	// Inicializácia - nastavenie všetkých pinov na výstupy
	if err := setRelayState(dev, 0xFF); err != nil {
		log.Fatalf("Chyba pri nastavení výstupov: %v", err)
	}

	// Ovládanie relé
	for {
		// Zapnutie prvého relé
		if err := setRelayState(dev, 0x01); err != nil {
			log.Printf("Chyba pri zapnutí relé 1: %v", err)
		} else {
			fmt.Println("Relé 1 je zapnuté")
		}
		time.Sleep(1 * time.Second)

		// Vypnutie prvého relé
		if err := setRelayState(dev, 0x00); err != nil {
			log.Printf("Chyba pri vypnutí relé 1: %v", err)
		} else {
			fmt.Println("Relé 1 je vypnuté")
		}
		time.Sleep(1 * time.Second)
	}
}
