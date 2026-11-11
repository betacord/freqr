// Commands related to importing channel list to the device.

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"krzywicki.pro/freqr/internal/csvutil"
	"krzywicki.pro/freqr/internal/radio"
)

var importFilePath string

var importChannelsCmd = &cobra.Command{
	Use:   "import-channels",
	Short: "Import channels from provided CSV file",
	Run: func(cmd *cobra.Command, args []string) {
		channels, err := csvutil.ReadChannelsFromCSV(importFilePath)
		if err != nil {
			fmt.Printf("Error while reading channels from CSV: %s.\n", err)
			return
		}

		if radio.SaveChannels(
			"/dev/cu.usbserial-0239ADD6",
			57600,
			channels,
		) != nil {
			fmt.Printf("Error while saving channels in the device: %s.\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(importChannelsCmd)

	importChannelsCmd.Flags().StringVarP(&importFilePath, "filepath", "p", "", "CSV input filepath")
	importChannelsCmd.MarkFlagRequired("filepath")
}
