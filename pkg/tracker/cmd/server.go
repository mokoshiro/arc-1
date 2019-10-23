package cmd

import (
	"context"

	"github.com/Bo0km4n/arc/internal/rpc"
	"github.com/Bo0km4n/arc/pkg/tracker/api"
	"github.com/Bo0km4n/arc/pkg/tracker/cmd/option"
	"github.com/Bo0km4n/arc/pkg/tracker/domain/repository"
	"github.com/Bo0km4n/arc/pkg/tracker/infra/db"
	"github.com/Bo0km4n/arc/pkg/tracker/logger"
	"github.com/Bo0km4n/arc/pkg/tracker/usecase"
	"github.com/spf13/cobra"
)

// serverCmd represents the web command
var serverCmd = &cobra.Command{
	Use: "server",
	PreRun: func(cmd *cobra.Command, args []string) {
		db.InitDB()
		logger.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Run Gateway API task
		Server(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&option.Opt.Port, "port", "p", 50051, "listen port")
	serverCmd.Flags().BoolVarP(&option.Opt.UseRedis, "use_redis", "", true, "use redis")
	serverCmd.Flags().StringVarP(&option.Opt.RedisHost, "redis_host", "", "127.0.0.1:6379", "redis host address")
	serverCmd.Flags().IntVarP(&option.Opt.RedisMaxIdle, "redis_max_idle", "", 32, "redis max idle connection")
	serverCmd.Flags().IntVarP(&option.Opt.RedisActive, "redis_max_active", "", 64, "redis max active connection")
	serverCmd.Flags().IntVarP(&option.Opt.RedisIdleTimeout, "redis_idle_timeout", "", 240, "redis idle timeout connection")
	serverCmd.Flags().IntVarP(&option.Opt.RedisKeyExpire, "redis_key_expire", "", 86400, "redis key expire")
	serverCmd.Flags().IntVarP(&option.Opt.GeoResolution, "geo_resolution", "", 9, "Geo hash resolution")
}

// Server returns API object
func Server(ctx context.Context) error {
	memberRepo := repository.NewMemberRepository(db.DB_REDIS)
	memberUsecase := usecase.NewMemberUsecase(memberRepo)
	trackerAPI := api.NewtrackerAPI(memberUsecase)

	server := rpc.NewServer(trackerAPI, rpc.WithPort(option.Opt.Port))
	defer func() {
		server.Stop(10)
		logger.Destruction()
	}()
	go server.Run()

	// Waiting the signals
	<-ctx.Done()
	return nil
}
