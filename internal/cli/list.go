// Commands related to listing connected USB devices and saved channels.

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"krzywicki.pro/freqr/internal/radio"
	"krzywicki.pro/freqr/internal/uiutil"
)

// A `list` cobra command handling listing USB devices.
var listUSBCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "Lists all available USB serial-compatibile devices",
	Run: func(cmd *cobra.Command, args []string) {
		ports, err := radio.GetUSBDevices()
		if err != nil {
			fmt.Println("USB devices not found!")
		}

		for i, port := range ports {
			fmt.Printf("[%d] %s\n", i, port)
		}
	},
}

var listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "List all saved channels",
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

		tbl := uiutil.PrepareChannelsTable(channels)
		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(listUSBCmd)
	rootCmd.AddCommand(listChannelsCmd)
}
