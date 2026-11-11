// Commands related to exporting channel list saved in the device.

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"krzywicki.pro/freqr/internal/config"
	"krzywicki.pro/freqr/internal/csvutil"
	"krzywicki.pro/freqr/internal/radio"
)

var exportFilePath string

var exportChannelsCmd = &cobra.Command{
	Use:   "export-channels",
	Short: "Export all saved channels to CSV file",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()

		channels, err := radio.GetAllChannels(
			cfg.Radio.PortName,
			cfg.Radio.BaudRate,
			cfg.Radio.MaxRetries,
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
