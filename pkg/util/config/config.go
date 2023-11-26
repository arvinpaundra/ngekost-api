package config

import (
	"github.com/arvinpaundra/ngekost-api/pkg/util/log"
	"github.com/spf13/viper"
)

func GetString(key string) string {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Logging().Error(err.Error())
	}

	return viper.GetString(key)
}

func GetInt(key string) int {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Logging().Error(err.Error())
	}

	return viper.GetInt(key)
}
