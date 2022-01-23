package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "remote-bmi",
	Long: `ISUCON Practice Remote Bench Marker Interface Server`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
