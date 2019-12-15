package cmd

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/room/api"
	"github.com/Bo0km4n/arc/pkg/room/cmd/option"
	"github.com/Bo0km4n/arc/pkg/room/infra/db"
	"github.com/Bo0km4n/arc/pkg/room/logger"
	"github.com/spf13/cobra"
)

// Cmd represents the web command
var coordinatorCmd = &cobra.Command{
	Use: "coordinator",
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.Init()
		db.InitMysql()
		db.InitRedisPool()
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
	coordinatorCmd.Flags().StringVarP(&option.Opt.MysqlHost, "mysql_host", "", "127.0.0.1", "mysql host address")
	coordinatorCmd.Flags().StringVarP(&option.Opt.MysqlPort, "mysql_port", "", "3306", "mysql port")
	coordinatorCmd.Flags().StringVarP(&option.Opt.MysqlPassword, "mysql_password", "", "root", "mysql password")
	coordinatorCmd.Flags().StringVarP(&option.Opt.MysqlUser, "mysql_user", "", "root", "mysql user")
	coordinatorCmd.Flags().StringVarP(&option.Opt.MysqlDatabase, "mysql_database", "", "arc", "mysql database")
	coordinatorCmd.Flags().StringVarP(&option.Opt.RedisHost, "redis_host", "", "127.0.0.1:6379", "redis host address")
	coordinatorCmd.Flags().IntVarP(&option.Opt.RedisMaxIdle, "redis_max_idle", "", 32, "redis max idle connection")
	coordinatorCmd.Flags().IntVarP(&option.Opt.RedisActive, "redis_max_active", "", 64, "redis max active connection")
	coordinatorCmd.Flags().IntVarP(&option.Opt.RedisIdleTimeout, "redis_idle_timeout", "", 240, "redis idle timeout connection")
	coordinatorCmd.Flags().IntVarP(&option.Opt.RedisKeyExpire, "redis_key_expire", "", 86400, "redis key expire")
	coordinatorCmd.Flags().BoolVarP(&option.Opt.Debug, "debug", "", false, "debug flag")
}

func Coordinator(ctx context.Context) {
	logger.L.Info("[Task]: Room Coordinator")
	api.Run(context.Background())
}
