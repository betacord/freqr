// A root CLI command implementation.

package cli

import "github.com/spf13/cobra"

// The root command.
var rootCmd = &cobra.Command{
	Use:   "freqr",
	Short: "A tool for Uniden radio scanner frequency programming",
}

// A function executing command handling.
//
// Returns
//   - error: when unable to parse arguments.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// another flags
	// rootCmd.PersistentFlags().BoolP("list-devices", "l", false, "List all devices")
}
