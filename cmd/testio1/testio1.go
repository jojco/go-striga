package main

// testovanie dosky s tromi relé, ktoré sú priamo pripojené na zbernicu
// CH1 je pin 37, CH2 je pin 38, CH3 je pin 40
// https://www.waveshare.com/wiki/RPi_Relay_Board

import (
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"
)

func main() {
	host.Init()
	t := time.NewTicker(500 * time.Millisecond)

	for l := gpio.Low; ; l = !l { // l = !l každu polsekundu neguje t.j. prepína relé
		rpi.P1_37.Out(l)
		<-t.C
	}
}
