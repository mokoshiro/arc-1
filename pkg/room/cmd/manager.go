package cmd

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/room/cmd/option"
	"github.com/Bo0km4n/arc/pkg/room/infra/db"
	"github.com/Bo0km4n/arc/pkg/room/logger"
	"github.com/spf13/cobra"
)

// Cmd represents the web command
var managerCmd = &cobra.Command{
	Use: "manager",
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.Init()
		db.InitMysql()
		db.InitRedisPool()
	},
	Run: func(cmd *cobra.Command, args []string) {
		Coordinator(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(managerCmd)
	managerCmd.Flags().StringVarP(&option.Opt.GlobalAddress, "global", "g", "0.0.0.0", "global address")
	managerCmd.Flags().BoolVarP(&option.Opt.Ws, "ws", "", true, "use websockets")
	managerCmd.Flags().IntVarP(&option.Opt.Port, "port", "p", 50051, "listen port")
	managerCmd.Flags().StringVarP(&option.Opt.MysqlHost, "mysql_host", "", "127.0.0.1", "mysql host address")
	managerCmd.Flags().StringVarP(&option.Opt.MysqlPort, "mysql_port", "", "3306", "mysql port")
	managerCmd.Flags().StringVarP(&option.Opt.MysqlPassword, "mysql_password", "", "root", "mysql password")
	managerCmd.Flags().StringVarP(&option.Opt.MysqlUser, "mysql_user", "", "root", "mysql user")
	managerCmd.Flags().StringVarP(&option.Opt.MysqlDatabase, "mysql_database", "", "arc", "mysql database")
	managerCmd.Flags().StringVarP(&option.Opt.RedisHost, "redis_host", "", "127.0.0.1:6379", "redis host address")
	managerCmd.Flags().IntVarP(&option.Opt.RedisMaxIdle, "redis_max_idle", "", 32, "redis max idle connection")
	managerCmd.Flags().IntVarP(&option.Opt.RedisActive, "redis_max_active", "", 64, "redis max active connection")
	managerCmd.Flags().IntVarP(&option.Opt.RedisIdleTimeout, "redis_idle_timeout", "", 240, "redis idle timeout connection")
	managerCmd.Flags().IntVarP(&option.Opt.RedisKeyExpire, "redis_key_expire", "", 86400, "redis key expire")
	managerCmd.Flags().BoolVarP(&option.Opt.Debug, "debug", "", false, "debug flag")
}

func Manager(ctx context.Context) {
	logger.L.Info("[Task]: Room Manager")
}
