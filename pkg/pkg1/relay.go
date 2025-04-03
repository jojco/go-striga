package pkg1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/stianeikeland/go-rpio"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

// ********************************************************
// Knihovňa na ovládanie RELE cez I2C
//
// ********************************************************

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

// ********************************************************
// Skúška funkčnosti relé na doske - individuálne zopnutie
// podľ stlačeného čísla od 0 do 7
// ********************************************************
func TestReleIndividual() {

	var arr = [8]byte{1, 4, 64, 16, 32, 128, 8, 2} // Pole s 8 prvkami - čísla reprezentujú relé od 1 po 8

	reader := bufio.NewReader(os.Stdin)

	// Vyber číslo dosky
	fmt.Print("Zadaj číslo dosky (0-1): ")
	boardInput, _ := reader.ReadString('\n')
	boardInput = strings.TrimSpace(boardInput)
	boardNum, err := strconv.Atoi(boardInput)
	if err != nil || boardNum < 0 || boardNum > 1 {
		log.Fatal("Neplatné číslo dosky. Zadaj 0 alebo 1.")
	}

	// Nastav adresu I2C na základe výberu dosky.
	switch boardNum {
	case 0:
		i2cAddress = i2cAddress + 0 // Príklad základnej adresy pre dosku 0
	case 1:
		i2cAddress = i2cAddress + 1 // Príklad adresy pre dosku 1
	}
	fmt.Printf("Používam I2C adresu: 0x%02X\n", i2cAddress)

	fmt.Print("Zadaj číslo relé (0-7): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	i, err := strconv.Atoi(input)
	if err != nil || i < 0 || i > 7 {
		log.Fatal("Neplatné číslo relé. Zadaj číslo od 0 do 7.")
	}

	var rele byte = arr[i]
	fmt.Println("Relé", i+1)
	// Zapni relé
	if err := toggleRelay(&device, true, rele); err != nil {
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

// ********************************************************
// Skúška funkčnosti všetkých relé na doskách
// postupne zapína relé od 1 do 8 na dvoch doskách
// ********************************************************
func TestReleAll() {

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

// ********************************************************
// Knihovňa na ovládanie RELE cez RPIO
//
// ********************************************************

const relayPin = 17 // GPIO pin na Raspberry Pi, ktorý ovláda relé (môžeš zmeniť podľa potreby)

// Inicializácia GPIO pinov
func InitRelay() error {
	err := rpio.Open()
	if err != nil {
		return fmt.Errorf("chyba pri otváraní GPIO: %v", err)
	}

	// Nastavenie GPIO pinu pre relé ako výstup
	pin := rpio.Pin(relayPin)
	pin.Output()

	return nil
}

// Zapnutie relé
func TurnRelayOn() {
	pin := rpio.Pin(relayPin)
	pin.High() // Nastaví pin na "vysokú" hodnotu (relé sa zapne)
}

// Vypnutie relé
func TurnRelayOff() {
	pin := rpio.Pin(relayPin)
	pin.Low() // Nastaví pin na "nízku" hodnotu (relé sa vypne)
}

// Zavretie GPIO pri ukončení
func CloseRelay() {
	rpio.Close()
}
