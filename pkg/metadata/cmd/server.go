package cmd

import (
	"github.com/Bo0km4n/arc/pkg/metadata/cmd/option"
	"github.com/Bo0km4n/arc/pkg/metadata/router"
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
	serverCmd.Flags().StringVarP(&option.Opt.RedisHost, "redis_host", "", "127.0.0.1:6379", "redis host address")
	serverCmd.Flags().IntVarP(&option.Opt.RedisMaxIdle, "redis_max_idle", "", 32, "redis max idle connection")
	serverCmd.Flags().IntVarP(&option.Opt.RedisActive, "redis_max_active", "", 64, "redis max active connection")
	serverCmd.Flags().IntVarP(&option.Opt.RedisIdleTimeout, "redis_idle_timeout", "", 240, "redis idle timeout connection")
}

// Server returns API object
func Server() {
	api := router.New()
	api.Run(":8080")
}
