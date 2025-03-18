//Hlavný program
//Obsahuje aj servisný mód
//

package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"

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

			pkg2.Udajezscd30()
			//pkg3.ReadTemperature()
			pkg1.OvladanieRele()
			webserver.Webserverstriga()
			dbstriga.DbStriga()
			pkg4.DInput()
		}
	}
}

// Servisný mód umožňuje zistiť nový hardware
func serviceMode() {
	fmt.Println("Servisný režim aktivovaný.")
	// Tu môžete pridať kód pre servisný režim
	pkg3.NajdiTeplomer()
	pkg1.TestRele()
	time.Sleep(2 * time.Second) // Simulácia práce v servisnom režime
	fmt.Println("Servisný režim ukončený. Návrat do hlavnej slučky.")
}
