package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/host/v3"
)

func main() {
	// Inicializácia hosta
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Vyhľadanie I2C zariadenia na zbernici 1
	// I2C zbernica 1 je obvykle predvolená na Raspberry Pi 4
	bus, err := i2c.New("/dev/i2c-1")
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	// Adresa I2C zariadenia (relé / GPIO expander)
	deviceAddr := uint16(0x20)

	// Otvorenie zariadenia
	dev := i2c.Dev{Bus: bus, Addr: deviceAddr}

	// Nastavenie GPIO pinu, ktorý ovláda relé (ak používate GPIO expander)
	// Tento krok závisí od toho, aké zariadenie máte pripojené na I2C
	// Povedzme, že ovládate pin 0, nastavíme ho ako výstup
	// (Na tomto mieste by ste prispôsobili zápis podľa toho, čo relé vyžaduje)
	err = dev.Write([]byte{0x01, 0xFF}) // Príklad: nastavte výstupy na "všetky piny vysoko"
	if err != nil {
		log.Fatal(err)
	}

	// Ovládanie relé
	// Predpokladajme, že relé je ovládané prepnutím hodnoty na príslušnom I2C byte
	// V tomto prípade posielame hodnotu na zapnutie relé.
	fmt.Println("Zapínam relé...")
	err = dev.Write([]byte{0x00, 0x01}) // Tento zápis zapne relé (nastavte podľa požiadavky)
	if err != nil {
		log.Fatal(err)
	}

	// Čakanie na nejaký čas, aby ste mohli pozorovať účinok
	time.Sleep(5 * time.Second)

	// Vypnutie relé
	fmt.Println("Vypínam relé...")
	err = dev.Write([]byte{0x00, 0x00}) // Tento zápis vypne relé
	if err != nil {
		log.Fatal(err)
	}

	// Program končí
	fmt.Println("Program ukončený.")
}
