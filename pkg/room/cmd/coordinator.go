package cmd

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/room/cmd/option"
	"github.com/Bo0km4n/arc/pkg/room/logger"
	"github.com/spf13/cobra"
)

// Cmd represents the web command
var coordinatorCmd = &cobra.Command{
	Use: "coordinator",
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Run Gateway API task
		Coordinator(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(coordinatorCmd)
	coordinatorCmd.Flags().BoolVarP(&option.Opt.Ws, "ws", "", true, "use websockets")
	coordinatorCmd.Flags().IntVarP(&option.Opt.Port, "port", "p", 50051, "listen port")
	coordinatorCmd.Flags().StringVarP(&option.Opt.RedisHost, "redis_host", "", "127.0.0.1:6379", "redis host address")
	coordinatorCmd.Flags().IntVarP(&option.Opt.RedisMaxIdle, "redis_max_idle", "", 32, "redis max idle connection")
	coordinatorCmd.Flags().IntVarP(&option.Opt.RedisActive, "redis_max_active", "", 64, "redis max active connection")
	coordinatorCmd.Flags().IntVarP(&option.Opt.RedisIdleTimeout, "redis_idle_timeout", "", 240, "redis idle timeout connection")
	coordinatorCmd.Flags().IntVarP(&option.Opt.RedisKeyExpire, "redis_key_expire", "", 86400, "redis key expire")
}

func Coordinator(ctx context.Context) {
	// brokerAPI := api.NewHTTPTrackerAPI()

	// if !option.Opt.Ws {
	// 	brokerAPI.Run()
	// } else {
	// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 		socket.Serve(w, r)
	// 	})
	// 	if err := http.ListenAndServe(":8000", nil); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	logger.L.Info("[Task]: Room Coordinator")
}
