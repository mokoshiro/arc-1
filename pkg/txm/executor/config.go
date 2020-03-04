package executor

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Mysql struct {
		Host         string
		Port         string
		User         string
		Password     string
		Database     string
		MaxIdleConns int
		MaxOpenConns int
		Timeout      int
	}
	Port        string
	CacheExpire int
}

var executorConf = &config{}

func initConfig(path string) {
	viper.SetConfigName("executor")
	viper.SetConfigType("yaml")
	filepath, _ := filepath.Abs(path)
	viper.AddConfigPath(filepath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(executorConf); err != nil {
		log.Fatal(err)
	}
}
