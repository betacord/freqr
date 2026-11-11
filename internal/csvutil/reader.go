package csvutil

import (
	"encoding/csv"
	"os"

	"krzywicki.pro/freqr/internal/radio/models"
)

func ReadChannelsFromCSV(fp string) ([]*models.Channel, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	channels := make([]*models.Channel, 0)

	for _, rec := range records[1:] {
		ch, err := models.NewChannelFromCSVRecord(rec)
		if err != nil {
			return nil, err
		}

		channels = append(channels, ch)
	}

	return channels, nil
}
