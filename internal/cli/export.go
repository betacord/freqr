// Commands related to exporting channel list saved in the device.

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"krzywicki.pro/freqr/internal/csvutil"
	"krzywicki.pro/freqr/internal/radio"
)

var exportFilePath string

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
			fmt.Printf("Error while getting saved channels: %s.\n", err)
			return
		}
		err = csvutil.ExportChannelsToCSV(exportFilePath, channels)
		if err != nil {
			fmt.Printf("Error while exporting channels to CSV file: %s.\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(exportChannelsCmd)

	exportChannelsCmd.Flags().StringVarP(&exportFilePath, "filepath", "p", "", "CSV output filepath")
	exportChannelsCmd.MarkFlagRequired("filepath")
}
