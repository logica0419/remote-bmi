package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string
	version    string
)

var rootCmd = &cobra.Command{
	Use:  "remote-bmi",
	Long: "Remote-BMI - Bench Marker Web UI for ISUCON Practice",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := loadConfig(configFile)
		if err != nil {
			log.Panicf("failed to load config: %v", err)
		}

		if version != "" {
			c.Version = version
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.json", "config file")
	rootCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "running ISUCON exercise version")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
