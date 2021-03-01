package betypes

import (
	"io/ioutil"
	"telegram-chat_bot/logger"

	"github.com/pelletier/go-toml"
)

const configsFile = "configs/configs.toml"

// Config structure to save the settings of the bot, database, etc.
type Config struct {
	Bot struct {
		Webhook   string `toml:"webhook"`
		Token     string `toml:"token"`
		Port      string `toml:"port"`
		ChannelID string `toml:"channel_id"`
		Polling   struct {
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
	b, err := ioutil.ReadFile(configsFile)
	if err != nil {
		logger.ForWarning(err.Error())
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		logger.ForWarning(err.Error())
	}
}

// GetConfig returns the configuration.
func GetConfig() Config {
	return config
}
