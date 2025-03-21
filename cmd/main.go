//Hlavný program
//Obsahuje aj servisný mód
//

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
	_ "github.com/mattn/go-sqlite3" // Import ovládača SQLite

	dbstriga "github.com/jojco/go-striga/db"
	"github.com/jojco/go-striga/pkg/pkg1"
	"github.com/jojco/go-striga/pkg/pkg2"
	"github.com/jojco/go-striga/pkg/pkg3"
	"github.com/jojco/go-striga/pkg/pkg4"
	webserver "github.com/jojco/go-striga/web"
)

func main() {

	//Inicializácia modulov
	pkg1.InitRele()          //načítanie súboru config_rele.json a uloženie do databázy
	pkg3.VytvorDBTeplomery() //načítanie súboru config_w1.json a uloženie údajov do databázy

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
// Načítanie údajov zo senzorov	a uloženie do databáz
// ***********************************************************************
func nacitanieUdajov() {
	pkg2.Udajezscd30()
	//pkg3.ReadTemperature()
	pkg1.OvladanieRele()
	webserver.Webserverstriga()
	dbstriga.DbStriga()
	pkg4.DInput()

	// Volanie funkcie ReadTemperature na location a uloženie do databázy
	location := "t1UK"
	temperatureData, err := pkg3.ReadTemperature(location)
	if err != nil {
		log.Fatal(err)
	}

	// Otvorenie alebo vytvorenie databázy SQLite3
	db, err := sql.Open("sqlite3", "./config_w1.db")
	if err != nil {
		log.Fatalf("Chyba pri otvorení databázy: %v", err)
	}
	defer db.Close()

	// Vytvorenie tabuľky, ak neexistuje
	_, err = db.Exec(`
	 	CREATE TABLE IF NOT EXISTS w1_devices (
			 sensorid TEXT PRIMARY KEY,
			 path TEXT,
			 location TEXT
	 	)
	`)
	if err != nil {
		log.Fatalf("Chyba pri vytváraní tabuľky: %v", err)
	}
	// Vloženie dát do databázy
	_, err = db.Exec(`
				INSERT INTO temperatures (sensor_id, location, temperature, timestamp)
				VALUES (?, ?, ?, ?);
		`, temperatureData.SensorID, temperatureData.Location, temperatureData.Temperature, temperatureData.Timestamp)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully")
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
	time.Sleep(2 * time.Second) // Simulácia práce v module TUV
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
	pkg1.TestRele()
	time.Sleep(2 * time.Second) // Simulácia práce v servisnom režime
	fmt.Println("Servisný režim ukončený. Návrat do hlavnej slučky.")
}
