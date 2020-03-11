package executor

import (
	"log"
	"path/filepath"

	"os"

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
	viper.SetConfigType("yaml")
	filepath, _ := filepath.Abs(path)
	f, err := os.Open(filepath)
	if err != nil {
		logger.Fatal(err)
	}
	defer f.Close()
	viper.AddConfigPath(filepath)
	viper.AutomaticEnv()

	if err := viper.ReadConfig(f); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(executorConf); err != nil {
		log.Fatal(err)
	}
}
