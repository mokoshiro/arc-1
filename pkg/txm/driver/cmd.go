package driver

import (
	"github.com/spf13/cobra"
)

type option struct {
	configFilePath string
}

var driverOption = &option{}

func NewDriverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "driver",
		Short: "d",
		Run: func(cmd *cobra.Command, args []string) {
			initConfig(driverOption.configFilePath)
			sqldb := mysql(
				driverConf.Mysql.Host,
				driverConf.Mysql.Port,
				driverConf.Mysql.User,
				driverConf.Mysql.Password,
				driverConf.Mysql.Database,
				driverConf.Mysql.MaxIdleConns,
				driverConf.Mysql.MaxOpenConns,
			)
			kvs := redisPool(
				driverConf.Redis.Host,
				driverConf.Redis.MaxIdle,
				driverConf.Redis.Active,
				driverConf.Redis.IdleTimeout,
			)
			s := NewDriverServer(sqldb, kvs)
			s.run()
		},
	}

	cmd.Flags().StringVarP(&driverOption.configFilePath, "config", "c", "./driver/config.yaml", "config file path")
	return cmd
}
