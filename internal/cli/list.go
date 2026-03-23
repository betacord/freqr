package cli

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all available USB serial-compatibile devices",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
