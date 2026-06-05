// Sending and receiving serial bus-related commands.

package radio

import (
	"errors"
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
	"go.bug.st/serial"

	"krzywicki.pro/freqr/internal/radio/models"
)

func GetAllChannels(portName string, baudRate int, maxRetries int) ([]*models.Channel, error) {
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

	port.Read(make([]byte, 1000))

	defer func(port serial.Port) {
		_, _ = port.Write([]byte("EPG\n\r"))
	}(port)

	// TODO: read channels number from config
	var channels []*models.Channel
	retries := 0
	for range maxRetries {
		channels, err = listChannels(port, 300)
		if err == nil {
			break
		}
		retries++
	}

	if retries > maxRetries {
		return nil, errors.New("maximum number of retries exceeded.")
	}

	return channels, nil
}

func listChannels(port serial.Port, amount int) ([]*models.Channel, error) {
	channels := make([]*models.Channel, 0)
	bar := progressbar.Default(int64(amount))

	for i := range amount {
		buff := make([]byte, 100)
		command := fmt.Sprintf("CIN,%d\r", i+1)
		// fmt.Printf("Input: %s\n", command)
		_, err := port.Write([]byte(command))
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
		// fmt.Println(channelData)
		channelObj, err := models.NewChannelFromText(channelData)
		if err != nil {
			// fmt.Printf("Error while parsing channel %d\n", i) // TODO: as log level
			// continue                                          // TODO: break or auto reset from start as option
			return nil, err
		}

		channels = append(channels, channelObj)
		bar.Add(1)
	}

	return channels, nil
}
