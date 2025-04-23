package pkg1

import (
	"bufio"
	"database/sql"
	"encoding/json"
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

type ReleOnBoards struct {
	ReleID      string `json:"releid"`
	ReleCode    string `json:"relecode"`    // kodované číslo relé v hexa
	WhatControl string `json:"whatcontrol"` // čo relé ovláda
	Board       string `json:"board"`       //adresa dosky 0x26, 0x27,...(nastav na doske)
}

type Config struct {
	Relays []ReleOnBoards `json:"relays"`
}

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

	//Načítanie zoznamu teplomerov zo súboru json do databázy
	config, err := loadConfig("config_relays.json")
	if err != nil {
		fmt.Println("Chyba pri načítaní konfigurácie rele:", err)
		return
	}
	// Tlač obsahu súboru config_w1.json
	for _, device := range config.Relays {
		fmt.Println("obsah súboru config_relays.json") //
		fmt.Println("ReleID :", device.ReleID)
		fmt.Println("ReleCode:", device.ReleCode)
		fmt.Println("WhatControl:", device.WhatControl)
		fmt.Println("Board:", device.Board)
		fmt.Println("---") // Oddeľovač pre lepšiu čitateľnosť

	}
	// Otvorenie alebo vytvorenie databázy SQLite3
	db, err := sql.Open("sqlite3", "./config_relays.db")
	if err != nil {
		log.Fatalf("Chyba pri otvorení databázy: %v", err)
	}
	//defer db.Close() // zabezpečí zatvorenie db po ukončení funkcie ale ja chcem aby bola prístupná počas chodu programu
	// Vytvorenie tabuľky, ak neexistuje
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS config_relays (
					releid TEXT,
					relecode TEXT,
					whatcontrol TEXT,
					board BYTE
			)
	`)

	if err != nil {
		log.Fatalf("Chyba pri vytváraní tabuľky: %v", err)
	}

	// Vloženie dát z JSON do databázy
	for _, device := range config.Relays {
		_, err = db.Exec(
			"INSERT INTO config_relays (releid, relecode, whatcontrol, board) VALUES (?, ?, ?, ?)",
			device.ReleID, device.ReleCode, device.WhatControl, device.Board,
		)
		if err != nil {
			log.Printf("Chyba pri vkladaní dát: %v", err)
		}
	}

	fmt.Println("Dáta o relé sú úspešne uložené do databázy.")
	fmt.Println("---") // Oddeľovač pre lepšiu čitateľnosť
	fmt.Println("")    // Oddeľovač pre lepšiu čitateľnosť
}

func loadConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

// **************************************************************
// Zapínanie príslušného relé - hlavný program
// **************************************************************
func ZapniRele(releid string, stav string) error {
	fmt.Println("ReleID :", releid)
	fmt.Println("Stav rele :", stav)
	return nil
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
