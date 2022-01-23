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

		e.Logger.Panic(e.Start(address))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&address, "address", "a", ":3000", "Address to listen")
}
