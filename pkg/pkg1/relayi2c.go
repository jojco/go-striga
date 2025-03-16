// modul na ovládanie relé cez i2c
// i2c_brcmstb            12288  0
// i2c_bcm2835            16384  0
// i2c_dev                16384  0
// sudo apt-get install libbcm2835-dev
// go get -u periph.io/x/periph
// skúškou cez i2cset -y 1 0x26(0x27)adrdosky  0x01register1  0x01 relé1
// 0x01 Relé 1
// 0x04 Relé 2
// 0x40 Relé 3
// 0x10 Relé 4
// 0x20 Relé 5
// 0x80 Relé 6
// 0x08 Relé 7
// 0x02 Relé 8
// 0x00 všetky relé vypnuté
// 0xFF všetky relé zapnuté
//

package pkg1

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
var i2cAddress uint16 = 0x26 // Adresa I2C zariadenia nastavená na 0x26 HEXA t.j. 38 DEC
var device i2c.Dev

//const bus = 0x01

func InitRele() {
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

	// Vypíšeme info o pripojení
	fmt.Println("I2C zariadenie pripojené na adrese", i2cAddress)

	// Vytvorenie I2C zariadenia na základe adresy
	device := i2c.Dev{Bus: bus, Addr: i2cAddress}

	//vynulovanie registra na doske
	device.Write([]byte{0x01, 0})
	device.Write([]byte{0x02, 0})
	device.Write([]byte{0x03, 0})
	fmt.Println("registre zariadenia sú vymazané", device)

}

func OvladanieRele() {

	for j := 0; j < 2; j++ {

		var arr = [8]byte{1, 4, 64, 16, 32, 128, 8, 2} // Pole s 8 prvkami

		for i := 0; i < 8; i++ {

			var rele byte = arr[i]
			fmt.Println("Relé", i+1)
			// Zapni relé
			if err := toggleRelay(&device, true, rele); err != nil { //jojco: err má vždy návratovú chybu rôznu od 0=nil; ak nie je chyba, tak je nil
				log.Fatal(err)
			}

			// Počkajte 2 sekúnd
			time.Sleep(1 * time.Second)

			// Vypni relé
			if err := toggleRelay(&device, false, 0); err != nil {
				log.Fatal(err)
			}

			// Počkajte 2 sekúnd
			time.Sleep(1 * time.Second)

		}
		i2cAddress = i2cAddress + 1
	}
}

func toggleRelay(device *i2c.Dev, state bool, ktorerele byte) error {
	var value byte

	if state {
		value = ktorerele // ktorerele je číslo podľa zistenia...čísla v poli   arr
	} else {
		value = ktorerele // Predpokladáme, že 0x00 vypne relé
	}

	// Zapíšeme hodnotu do zariadenia na určený register
	// Tento príklad používá register 0x01
	_, err := device.Write([]byte{0x01, value})
	if err != nil {
		return fmt.Errorf("chyba pri zapise do I2C zariadenia: %v", err)
	}

	fmt.Printf("je teraz  %v\n", state)
	return nil
}

func TestRele() {

	for j := 0; j < 2; j++ {

		var arr = [8]byte{1, 4, 64, 16, 32, 128, 8, 2} // Pole s 8 prvkami

		for i := 0; i < 8; i++ {

			var rele byte = arr[i]
			fmt.Println("Relé", i+1)
			// Zapni relé
			if err := toggleRelay(&device, true, rele); err != nil { //jojco: err má vždy návratovú chybu rôznu od 0=nil; ak nie je chyba, tak je nil
				log.Fatal(err)
			}

			// Počkajte 2 sekúnd
			time.Sleep(1 * time.Second)

			// Vypni relé
			if err := toggleRelay(&device, false, 0); err != nil {
				log.Fatal(err)
			}

			// Počkajte 2 sekúnd
			time.Sleep(1 * time.Second)

		}
		i2cAddress = i2cAddress + 1
	}
}
