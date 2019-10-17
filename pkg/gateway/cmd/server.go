package cmd

import (
	"context"
	"log"

	"github.com/Bo0km4n/arc/pkg/gateway/cmd/option"
	"github.com/Bo0km4n/arc/pkg/gateway/router"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
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
	serverCmd.Flags().StringVarP(&option.Opt.TrackerHost, "tracker_host", "", "127.0.0.1:50052", "tracker host address")
}

// Server returns API object
func Server() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	api, err := router.New(context.Background(), logger, option.Opt)
	if err != nil {
		log.Fatal(err)
	}
	api.Run(":8080")
}
