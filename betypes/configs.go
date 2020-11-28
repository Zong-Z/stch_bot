package betypes

import (
	"fmt"
	"telegram-chat_bot/logger"

	"github.com/BurntSushi/toml"
)

// Config struct for saving bot settings.
type Config struct {
	Bot struct {
		Webhook string `toml:"webhook"`
		Token   string `toml:"token"`
		Port    string `toml:"port"`
	} `toml:"bot"`
	DB struct {
		Redis struct {
			Addr     string `toml:"addr"`
			Password string `toml:"password"`
			Db       int    `toml:"db"`
		} `toml:"redis"`
	} `toml:"database"`
	Chat struct {
		Queue int `toml:"queue"`
		Users int `toml:"users"`
	} `toml:"chat"`
}

var config Config

func init() {
	_, err := toml.DecodeFile("configs/configs.toml", &config)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s. Failed to load \"configs.toml\".", err.Error()))
		panic(err)
	}
}

// GetBotConfig return bot config.
func GetBotConfig() Config {
	return config
}
