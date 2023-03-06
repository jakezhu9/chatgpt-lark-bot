package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault("port", 3000)
	v.SetDefault("lark_base_url", "https://open.larksuite.com")
	v.SetDefault("bot_name", "bot")

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./../..")
	v.AddConfigPath("./../")
	v.AddConfigPath("./")
	_ = v.ReadInConfig()

	return v
}
