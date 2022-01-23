package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "remote-bmi",
	Long: `ISUCON Practice Remote Bench Marker Interface Server`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cfgFile, err := cmd.PersistentFlags().GetString("config")
		if err != nil {
			cfgFile = ""
		}

		err = loadConfig(cfgFile)
		if err != nil {
			log.Panicf("failed to load config: %v", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "config.json", "config file (default is config.json)")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
