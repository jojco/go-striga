package gpio_control

import (
	"fmt"
	"os/exec"
)

const (
	gpioPin = "17" // GPIO pin number to control
)

// ControlGPIO controls the state of a GPIO pin.
func ControlGPIO(on bool) error {
	var cmd *exec.Cmd
	if on {
		cmd = exec.Command("gpio", "-g", "write", gpioPin, "1")
	} else {
		cmd = exec.Command("gpio", "-g", "write", gpioPin, "0")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing gpio command: %v, output: %s", err, output)
	}

	return nil
}
