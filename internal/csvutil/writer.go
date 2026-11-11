package csvutil

import (
	"encoding/csv"
	"os"

	"krzywicki.pro/freqr/internal/radio/models"
)

var channelsCSvHeader = []string{
	"ID",
	"Name",
	"Freq",
	"Mod",
	"Delay",
	"Lockout",
	"Priority",
}

func ExportChannelsToCSV(fp string, channels []*models.Channel) error {
	file, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	_ = writer.Write(channelsCSvHeader)

	for _, ch := range channels {
		err := writer.Write(ch.ToCSVRow())
		if err != nil {
			return err
		}
	}

	return nil
}
