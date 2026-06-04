package radio

import (
	"fmt"

	"go.bug.st/serial"
)

func GetUSBDevices() ([]string, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}
	if len(ports) == 0 {
		return nil, fmt.Errorf("Ports not found")
	}

	return ports, nil
}
