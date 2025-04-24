//Hlavný program
//Obsahuje aj servisný mód

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
	_ "github.com/mattn/go-sqlite3" // Import ovládača SQLite

	"github.com/jojco/go-striga/pkg/pkg1"
	"github.com/jojco/go-striga/pkg/pkg2"
	"github.com/jojco/go-striga/pkg/pkg3"
	//"github.com/jojco/go-striga/pkg/pkg4"
	//webserver "github.com/jojco/go-striga/web"
)

func main() {

	//Inicializácia modulov
	pkg1.InitRele()      //inicializácia rel
	pkg3.InitTeplomery() //inicializácia teplomerov

	//Táto časť má umožniť
	//za behu programu spustiť servisný mód stlačením S
	//za behu programu ukončiť program stlačením Q

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	serviceChan := make(chan bool) // Kanál na signalizáciu servisného režimu

	go func() {
		for {
			char, key, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}

			if char == 's' || char == 'S' {
				serviceChan <- true // Signalizácia servisného režimu
			} else if key == keyboard.KeyCtrlC {
				fmt.Println("Ukončujem...")
				close(serviceChan)
				return
			}
		}
	}()

	for {
		select {
		case <-serviceChan:
			serviceMode()
		default:
			// Tu môžete pridať kód, ktorý sa má vykonávať v hlavnom programe
			fmt.Println("Hlavný program beží...")
			nacitanieUdajov()
			vykurovanieUK()
			pripravaTeplejVodyTUV()
			akumulacia()
			vetranieVZT()
			fotovoltikaFVE()
			zahrada()
		}
	}
}

// ***********************************************************************
// Načítanie údajov zo senzorov	a uloženie do databázy
// ***********************************************************************

func nacitanieUdajov() {
	//  Otvorenie alebo vytvorenie databázy SQLite3
	db, err := sql.Open("sqlite3", "./striga.db")
	if err != nil {
		log.Fatalf("Chyba pri otvorení databázy: %v", err)
	}
	defer db.Close() // zabezpečí zatvorenie db po ukončení funkcie
	// Vytvorenie tabuľky, ak neexistuje
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS striga (
					miesto TEXT,
					hodnota float32,
					cas time.Time
			)
	`)

	if err != nil {
		log.Fatalf("Chyba pri vytváraní tabuľky: %v", err)
	}

	// Volanie funkcie na načítanie údajov z SCD30
	fmt.Println("Spustená funkcia načítania údajov z SCD30")
	co2, vlhkost, teplota := pkg2.Udajezscd30()
	fmt.Printf("CO2: %f, Vlhkosť: %f, Teplota: %f\n", co2, vlhkost, teplota)

	// Volanie funkcie na čítanie hodnôt z teplomerov
	fmt.Println("Spustená funkcia načítania údajov z teplomerov na 1WIRE / I2C")
	location := "t1UK"
	sensorid, temperature, timestamp, err := pkg3.ReadTemperature(location)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Location:%v, SensorID: %v, Teplota: %f, Time: %s\n", location, sensorid, temperature, timestamp)

	fmt.Println("Dáta sú úspešne uložené do databázy.")
	fmt.Println("---") // Oddeľovač pre lepšiu čitateľnosť
	fmt.Println("")    // Oddeľovač pre lepšiu čitateľnosť
}

// ***********************************************************************
// Obsluha modulov riadiaceho systému - hlavná a výkonná činnosť programu
// ***********************************************************************
func vykurovanieUK() {
	fmt.Println("Modul UK")
	time.Sleep(2 * time.Second) // Simulácia práce v module UK
}
func pripravaTeplejVodyTUV() {
	fmt.Println("Modul TUV")
	time.Sleep(2 * time.Second)      // Simulácia práce v module TUV
	pkg1.ZapniRele("rele4", "zapni") // cirkulacne cerpadlo TUV
}
func akumulacia() {
	fmt.Println("Modul AKU")
	time.Sleep(2 * time.Second) // Simulácia práce v module AKU
}
func vetranieVZT() {
	fmt.Println("Modul VZT")
	time.Sleep(2 * time.Second) // Simulácia práce v module VZT
}
func fotovoltikaFVE() {
	fmt.Println("Modul FVE")
	time.Sleep(2 * time.Second) // Simulácia práce v module FVE
}
func zahrada() {
	fmt.Println("Modul Zahrada")
	time.Sleep(2 * time.Second) // Simulácia práce v module záhrada
}

// *******************************************************************
// Servisný mód t.j. napríklad umožňuje zistiť nový hardware a pod.
// *******************************************************************
func serviceMode() {
	fmt.Println("Servisný režim aktivovaný.")
	// Tu môžete pridať kód pre servisný režim
	pkg3.NajdiTeplomer()
	pkg1.TestReleAll()
	pkg1.TestReleIndividual()
	time.Sleep(2 * time.Second) // Simulácia práce v servisnom režime
	fmt.Println("Servisný režim ukončený. Návrat do hlavnej slučky.")

}
