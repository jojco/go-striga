// Balík na obsluhu digitálnych vstupov
//*************************************
//Používané PINy
//
//
//
//

package pkg4

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func DInput() {
	// Otvorí pamäť pre prístup k GPIO
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close() // Zatvorí pamäť po ukončení programu

	// Definícia vstupného pinu (napr. GPIO 17)
	inputPin := rpio.Pin(17)
	inputPin.Input() // Nastaví pin ako vstupný

	// Obsluha prerušenia pre elegantné ukončenie programu
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	// Nekonečná slučka pre čítanie vstupu
	for {
		select {
		case <-signalChannel:
			fmt.Println("\nUkončujem program...")
			return
		default:
			// Čítanie hodnoty vstupu
			inputValue := inputPin.Read()
			fmt.Printf("Hodnota vstupu (GPIO 17): %v\n", inputValue)
			time.Sleep(500 * time.Millisecond) // Pauza 500 ms
		}
	}
}
