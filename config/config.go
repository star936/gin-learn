package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

var Config *viper.Viper

func SetUp() {
	Config = viper.New()
	Config.SetConfigType("yaml")
	Config.SetConfigName("settings")
	Config.AddConfigPath("./config")

	if err := Config.ReadInConfig(); err != nil {
		log.Fatalf("Read configuration file failed: %s", err)
	}
}

func GetString(k ...string) string {
	key := strings.Join(k, ".")
	return Config.GetString(key)
}