package config

import (
	"log"

	"github.com/spf13/viper"
)

var cfg *viper.Viper

func Init(env string) {
	cfg = viper.New()

	readConfig("default")
	readConfig(env)
}

func readConfig(env string) {
	config := viper.New()
	config.AddConfigPath("./config")
	config.SetConfigName(env)
	config.SetConfigType("json")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(env + " configuration file not found")
		} else {
			log.Fatal("error on parsing " + env + " configuration file")
		}
	}
	cfg.MergeConfigMap(config.AllSettings())
}

func GetConfig() *viper.Viper {
	return cfg
}
