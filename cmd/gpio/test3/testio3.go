// relay/relay.go
package relay

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
)

const relayPin = 17 // GPIO pin na Raspberry Pi, ktorý ovláda relé (môžeš zmeniť podľa potreby)

// Inicializácia GPIO pinov
func InitRelay() error {
	err := rpio.Open()
	if err != nil {
		return fmt.Errorf("Chyba pri otváraní GPIO: %v", err)
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
