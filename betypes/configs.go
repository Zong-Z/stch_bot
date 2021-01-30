package betypes

import (
	"fmt"
	"io/ioutil"
	"telegram-chat_bot/logger"

	"github.com/pelletier/go-toml"
)

// Config struct for saving bot settings.
type Config struct {
	Bot struct {
		Webhook string `toml:"webhook"`
		Token   string `toml:"token"`
		Port    string `toml:"port"`
		Polling struct {
			Offset  int `toml:"offset"`
			Limit   int `toml:"limit"`
			Timeout int `toml:"timeout"`
		} `toml:"polling"`
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
	b, err := ioutil.ReadFile("configs/configs.toml")
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
	}
}

// GetConfig return bot config.
func GetConfig() Config {
	return config
}
