// Sending and receiving serial bus-related commands.

package radio

import (
	"fmt"
	"time"

	"go.bug.st/serial"
)

func GetAllChannels(portName string, baudRate int) ([]*Channel, error) {
	mode := &serial.Mode{
		BaudRate: baudRate,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		return nil, err
	}

	_, err = port.Write([]byte("PRG\r"))
	if err != nil {
		return nil, err
	}
	// TODO: read first response as buffer reset from previous session
	buff := make([]byte, 100)
	port.Read(buff) // TODO read resp; should be PRG OK and reset buffer

	channels := make([]*Channel, 0)

	for i := range 300 {
		buff := make([]byte, 100)
		command := fmt.Sprintf("CIN,%d\r", i+1)
		// fmt.Printf("Input: %s\n", command)
		_, err = port.Write([]byte(command))
		if err != nil {
			return nil, err
		}
		// TODO: read sleep time from config file
		time.Sleep(100 * time.Millisecond)
		n, err := port.Read(buff)
		if err != nil {
			return nil, err
		}
		if n == 0 {
			break
		}

		channelData := string(buff[:n])
		fmt.Println(channelData)
		channelObj, err := newChannelFromText(channelData)
		if err != nil {
			fmt.Printf("Error while parsing channel %d\n", i) // TODO: as log level
			continue
		}
		// fmt.Println(channelObj.Text())
		channels = append(channels, channelObj)
	}

	_, err = port.Write([]byte("EPG\n\r"))
	if err != nil {
		return nil, err
	}

	return channels, nil
}
