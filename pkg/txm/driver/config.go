package driver

import (
	"log"
	"os"
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
	Redis struct {
		Host        string // IP:Port
		MaxIdle     int
		Active      int
		IdleTimeout int
	}
	CacheExpire     int
	Port            string
	GeoHashAccuracy int
}
 
var driverConf = &config{}

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
	if err := viper.Unmarshal(driverConf); err != nil {
		log.Fatal(err)
	}
}
