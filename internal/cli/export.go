// Commands related to exporting channel list saved in the device.

package cli

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"krzywicki.pro/freqr/internal/radio"
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

var filePath string

var exportChannelsCmd = &cobra.Command{
	Use:   "export-channels",
	Short: "Export all saved channels to CSV file",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: read port and rate from cfg file
		// TODO: read output (list / table or csv) from cfg
		// TODO: read max retries from config
		channels, err := radio.GetAllChannels(
			"/dev/cu.usbserial-0239ADD6",
			57600,
			1,
		)
		if err != nil {
			fmt.Printf("Error while listing saved channels: %s.\n", err)
			return
		}
		err = exportChannelsToCSV(filePath, channels)
		if err != nil {
			fmt.Printf("Error while exporting channels to CSV file: %s.\n", err)
		}
	},
}

func exportChannelsToCSV(fp string, channels []*models.Channel) error {
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

func init() {
	rootCmd.AddCommand(exportChannelsCmd)

	exportChannelsCmd.Flags().StringVarP(&filePath, "filepath", "p", "", "CSV output filepath")
	exportChannelsCmd.MarkFlagRequired("filepath")
}
