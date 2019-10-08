package cmd

import (
	"github.com/Bo0km4n/arc/pkg/gateway/router"
	"github.com/spf13/cobra"
)

// serverCmd represents the web command
var serverCmd = &cobra.Command{
	Use: "server",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Nothing
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Handling authenticate middleware.
		Server()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func Server() {
	api := router.New()
	api.Run(":8080")
}
