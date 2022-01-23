package cmd

import (
	"github.com/logica0419/remote-bmi/server/router"
	"github.com/spf13/cobra"
)

var address string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Remote-BMI Server",
	Run: func(cmd *cobra.Command, args []string) {
		e := router.SetupEcho()

		if address != "" {
			cfg.Address = address
		}

		e.Logger.Panic(e.Start(cfg.Address))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&address, "address", "a", "", "Address to listen")
}
