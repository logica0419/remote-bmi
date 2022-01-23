package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var address string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Remote-BMI Server",
	Run: func(cmd *cobra.Command, args []string) {
		if address != "" {
			c.Address = address
		}

		r, err := setupRouter(c)
		if err != nil {
			log.Panicf("failed to setup router: %v", err)
		}

		r.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&address, "address", "a", "", "Address to listen")
}
