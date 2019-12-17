package cmd

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/room/client"
	"github.com/Bo0km4n/arc/pkg/room/logger"
	"github.com/spf13/cobra"
)

var (
	in string
)

var clientCmd = &cobra.Command{
	Use: "client",
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Run Gateway API task
		Client(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringVarP(&in, "in", "i", "", "json input")
}

func Client(ctx context.Context) {
	client.Run(in)
}
