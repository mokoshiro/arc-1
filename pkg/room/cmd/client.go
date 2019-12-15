package cmd

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/room/client"
	"github.com/Bo0km4n/arc/pkg/room/logger"
	"github.com/spf13/cobra"
)

var (
	clientID         string
	clientCredential string
	clientHost       string
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
	clientCmd.Flags().StringVarP(&clientID, "id", "", "user1", "client peer id")
	clientCmd.Flags().StringVarP(&clientCredential, "credential", "c", "password", "client credential")
	clientCmd.Flags().StringVarP(&clientHost, "host", "", "localhost:8000", "target host")
}

func Client(ctx context.Context) {
	client.Run(clientID, clientCredential, clientHost)
}
