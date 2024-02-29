package main

import (
	"log"
	"time"

	gpio_control "github.com/jojco/go-striga/gpio" // Update with your actual package path
)

func main() {
	// Continuously control GPIO pin
	for {
		// Control GPIO pin
		err := gpio_control.ControlGPIO(true) // Turn on GPIO pin
		if err != nil {
			log.Fatal("Error controlling GPIO pin:", err)
		}
		time.Sleep(1 * time.Second)           // Wait for 1 second
		err = gpio_control.ControlGPIO(false) // Turn off GPIO pin
		if err != nil {
			log.Fatal("Error controlling GPIO pin:", err)
		}

		time.Sleep(2 * time.Second) // Control GPIO pin every 2 seconds
	}
}
