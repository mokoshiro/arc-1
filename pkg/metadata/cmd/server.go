package cmd

import (
	"context"

	"github.com/Bo0km4n/arc/internal/rpc"
	"github.com/Bo0km4n/arc/pkg/metadata/api"
	"github.com/Bo0km4n/arc/pkg/metadata/cmd/option"
	"github.com/Bo0km4n/arc/pkg/metadata/domain/repository"
	"github.com/Bo0km4n/arc/pkg/metadata/infra/db"
	"github.com/Bo0km4n/arc/pkg/metadata/usecase"
	"github.com/spf13/cobra"
)

// serverCmd represents the web command
var serverCmd = &cobra.Command{
	Use: "server",
	PreRun: func(cmd *cobra.Command, args []string) {
		db.InitDB()
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
	serverCmd.Flags().IntVarP(&option.Opt.RedisKeyExpire, "redis_key_expire", "", 3600, "redis key expire")
}

// Server returns API object
func Server(ctx context.Context) error {
	memberRepo := repository.NewMemberRepository(db.DB_REDIS)
	memberUsecase := usecase.NewMemberUsecase(memberRepo)
	metadataAPI := api.NewMetadataAPI(memberUsecase)

	server := rpc.NewServer(metadataAPI, rpc.WithPort(option.Opt.Port))
	defer func() {
		server.Stop(10)
	}()
	go server.Run()

	// Waiting the signals
	<-ctx.Done()
	return nil
}
