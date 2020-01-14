package driver

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Redis struct {
		Host        string
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
	viper.SetConfigName("driver")
	viper.SetConfigType("yaml")
	filepath, _ := filepath.Abs(path)
	viper.AddConfigPath(filepath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(driverConf); err != nil {
		log.Fatal(err)
	}
}
