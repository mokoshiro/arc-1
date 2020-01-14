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
			kvsPool := kvs(
				driverConf.Redis.Host,
				driverConf.Redis.MaxIdle,
				driverConf.Redis.Active,
				driverConf.Redis.IdleTimeout,
			)
			s := NewDriverServer(kvsPool)
			s.run()
		},
	}

	cmd.Flags().StringVarP(&driverOption.configFilePath, "config", "c", "./driver/config.yaml", "config file path")
	return cmd
}
