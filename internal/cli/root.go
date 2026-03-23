package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "freqr",
	Short: "A tool for Uniden radio scanner frequency programming",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// another flags
	// rootCmd.PersistentFlags().BoolP("list-devices", "l", false, "List all devices")
}
