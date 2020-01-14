package executor

import (
	"github.com/spf13/cobra"
)

type option struct {
	configFilePath string
}

var driverOption = &option{}

func NewExecutorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "executor",
		Short: "e",
		Run: func(cmd *cobra.Command, args []string) {
			initConfig(driverOption.configFilePath)
			db := mysql(
				executorConf.Mysql.Host, executorConf.Mysql.Port, executorConf.Mysql.User,
				executorConf.Mysql.Password, executorConf.Mysql.Database, executorConf.Mysql.MaxIdleConns,
				executorConf.Mysql.MaxOpenConns,
			)
			s := NewExecutorServer(db)
			s.run()
		},
	}

	cmd.Flags().StringVarP(&driverOption.configFilePath, "config", "c", "./driver/config.yaml", "config file path")
	return cmd
}
