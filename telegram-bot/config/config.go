package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Bot BotConfig `mapstructure:"bot"`
}

type BotConfig struct {
	Token  string `mapstructure:"token"`
	ChatId string `mapstructure:"chat_id"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("bot")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
