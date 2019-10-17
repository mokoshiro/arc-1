package cmd

import (
	"github.com/Bo0km4n/arc/pkg/gateway/cmd/option"
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
		// Run Gateway API task
		Server()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&option.Opt.MetadataHost, "metadata_host", "", "127.0.0.1:50051", "metadata host address")
}

// Server returns API object
func Server() {
	api := router.New()
	api.Run(":8080")
}
