package gpio_control

import (
	"fmt"
	"os/exec"
)

const (
	gpioCommand = "gpio"
	gpioPin     = "17" // GPIO pin number to control
)

// ControlGPIO controls the state of a GPIO pin.
func ControlGPIO(on bool) error {
	// Look for gpio executable in the PATH
	gpioPath, err := exec.LookPath(gpioCommand)
	if err != nil {
		return fmt.Errorf("gpio command not found: %v", err)
	}

	var cmd *exec.Cmd
	if on {
		cmd = exec.Command(gpioPath, "-g", "write", gpioPin, "1")
	} else {
		cmd = exec.Command(gpioPath, "-g", "write", gpioPin, "0")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing gpio command: %v, output: %s", err, output)
	}

	return nil
}
