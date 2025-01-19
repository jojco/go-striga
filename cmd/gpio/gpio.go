package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const gpioPath = "/sys/class/gpio"

func exportPin(pin int) error {
	pinStr := strconv.Itoa(pin)
	_, err := os.Stat(gpioPath + "/gpio" + pinStr)
	if os.IsNotExist(err) {
		file, err := os.OpenFile(gpioPath+"/export", os.O_WRONLY, os.ModeExclusive)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(pinStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func unexportPin(pin int) error {
	pinStr := strconv.Itoa(pin)
	_, err := os.Stat(gpioPath + "/gpio" + pinStr)
	if err == nil {
		file, err := os.OpenFile(gpioPath+"/unexport", os.O_WRONLY, os.ModeExclusive)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(pinStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func setPinDirection(pin int, direction string) error {
	pinStr := strconv.Itoa(pin)
	file, err := os.OpenFile(gpioPath+"/gpio"+pinStr+"/direction", os.O_WRONLY, os.ModeExclusive)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(direction)
	return err
}

func writePin(pin int, value int) error {
	pinStr := strconv.Itoa(pin)
	file, err := os.OpenFile(gpioPath+"/gpio"+pinStr+"/value", os.O_WRONLY, os.ModeExclusive)
	if err != nil {
		return err
	}
	defer file.Close()

	valueStr := "0"
	if value != 0 {
		valueStr = "1"
	}

	_, err = file.WriteString(valueStr)
	return err
}

func main() {
	// GPIO pin number to control
	pin := 37

	// Export GPIO pin
	err := exportPin(pin)
	if err != nil {
		fmt.Println("Error exporting GPIO pin:", err)
		return
	}
	defer func() {
		// Clean up: Unexport GPIO pin
		err := unexportPin(pin)
		if err != nil {
			fmt.Println("Error unexporting GPIO pin:", err)
		}
	}()

	// Set pin direction to out
	err = setPinDirection(pin, "out")
	if err != nil {
		fmt.Println("Error setting pin direction:", err)
		return
	}

	// Turn on the pin
	err = writePin(pin, 1)
	if err != nil {
		fmt.Println("Error writing to pin:", err)
		return
	}
	fmt.Println("Pin turned on")

	// Sleep for 3 seconds
	time.Sleep(3 * time.Second)

	// Turn off the pin
	err = writePin(pin, 0)
	if err != nil {
		fmt.Println("Error writing to pin:", err)
		return
	}

	fmt.Println("Pin turned off")
}
