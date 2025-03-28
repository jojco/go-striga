package pkg2

import (
	"log"
	"time"

	"github.com/pvainio/scd30"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

type Merania struct {
	CO2         float32
	Humidity    float32
	Temperature float32
}

var m Merania

func Udajezscd30() (float32, float32, float32) {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	dev, err := scd30.Open(bus)
	if err != nil {
		log.Fatal(err)
	}

	var interval uint16 = 10

	dev.StartMeasurements(interval)

	//meranie len 1x
	time.Sleep(time.Duration(interval) * time.Second)
	if hasMeasurement, err := dev.HasMeasurement(); err != nil {
		log.Fatalf("error %v", err)
	} else if !hasMeasurement {
		return 0, 0, 0
	}

	m, err := dev.GetMeasurement()
	if err != nil {
		log.Fatalf("error %v", err)
	}

	//log.Printf("Got measure %f ppm %f%% %fC", m.CO2, m.Humidity, m.Temperature)

	return m.CO2, m.Humidity, m.Temperature
}
