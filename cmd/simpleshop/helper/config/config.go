package config

import config "github.com/spf13/viper"

//GetConfig :
func GetConfig() (*config.Viper, error) {
	config.SetConfigType("json")
	config.AddConfigPath("./config")
	config.SetConfigName("config.json")

	conferr := config.ReadInConfig()
	return config.GetViper(), conferr
}
