// Sending and receiving serial bus-related commands.

package radio

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/schollz/progressbar/v3"
	"go.bug.st/serial"

	"krzywicki.pro/freqr/internal/config"
	"krzywicki.pro/freqr/internal/radio/models"
)

// A function listing all channels saved in the radio device.
//
// Parameters
//   - portName - the port identifier
//   - baudRate - the baud rate
//   - maxRetries - the maximum number of retries while reading full channel list
//
// Returns
//   - slice with channels saved in the device
//   - error - when maximum number of retries occured
func GetAllChannels(portName string, baudRate int, maxRetries int) ([]*models.Channel, error) {
	mode := &serial.Mode{
		BaudRate: baudRate,
	}
	cfg := config.Get()

	port, err := serial.Open(portName, mode)
	if err != nil {
		return nil, err
	}

	_, err = port.Write([]byte("PRG\r"))
	if err != nil {
		return nil, err
	}

	buff := make([]byte, 128)
	n, _ := port.Read(buff)
	if string(buff[:n]) == "PRG OK" {
		buff = make([]byte, 128)
	}

	port.Read(make([]byte, 1024))

	defer func(port serial.Port) {
		_, _ = port.Write([]byte("EPG\n\r"))
	}(port)

	var channels []*models.Channel
	retries := 0
	for range maxRetries {
		channels, err = listChannels(port, cfg.Radio.TotalChannels)
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

// A function saving list of channels in the device
//
// Parameters
//   - portName - the port identifier
//   - baudRate - the baud rate
//   - channels - the slice with channels
//
// Returns
//   - error - when any of the steps failed
func SaveChannels(portName string, baudRate int, channels []*models.Channel) error {
	cfg := config.Get()
	mode := &serial.Mode{
		BaudRate: baudRate,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		return err
	}

	_, err = port.Write([]byte("PRG\r"))
	if err != nil {
		return err
	}

	buff := make([]byte, 128)
	n, _ := port.Read(buff)
	if string(buff[:n]) == "PRG OK" {
		buff = make([]byte, 128)
	}

	port.Read(make([]byte, 1024))

	defer func(port serial.Port) {
		_, _ = port.Write([]byte("EPG\n\r"))
		_ = port.Close()
	}(port)

	bar := progressbar.Default(int64(len(channels)))

	for _, ch := range channels {
		buff := make([]byte, 1024)
		_, err := port.Write([]byte(fmt.Sprintf("%s\r", ch.ToString())))
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(cfg.Radio.WriteSleepTime) * time.Millisecond)
		bar.Add(1)

		n, err := port.Read(buff)
		if n == 0 {
			return errors.New("Empty response from device")
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// A function listing channels.
//
// Parameters
//   - portName - the port identifier
//   - amount - the size of channels memory in the device
func listChannels(port serial.Port, amount int) ([]*models.Channel, error) {
	cfg := config.Get()
	channels := make([]*models.Channel, 0)
	bar := progressbar.Default(int64(amount))

	for i := range amount {
		buff := make([]byte, 128)
		command := fmt.Sprintf("CIN,%d\r", i+1)
		// fmt.Printf("Input: %s\n", command)
		_, err := port.Write([]byte(command))
		if err != nil {
			return nil, err
		}

		time.Sleep(time.Duration(cfg.Radio.ReadSleepTime) * time.Millisecond)
		n, err := port.Read(buff)
		if err != nil {
			return nil, err
		}
		if n == 0 {
			break
		}

		channelData := string(buff[:n])
		channelObj, err := models.NewChannelFromText(channelData)
		if err != nil {
			log.Printf("Error while parsing channel %d\n", i)
			if cfg.Radio.BreakWhenErr == false {
				continue
			}
			return nil, err
		}

		channels = append(channels, channelObj)
		bar.Add(1)
	}

	return channels, nil
}
