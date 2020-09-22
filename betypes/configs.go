package betypes

import (
	"encoding/json"
	"io/ioutil"
	"telegram-chat_bot/loger"
)

type config struct {
	WebHook     string `json:"web_hook"`
	BotToken    string `json:"bot_token"`
	BotPort     string `json:"bot_port"`
	RedisConfig struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	}
}

var _config config

func init() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		loger.ForLog("Error, failed to load \"config.json\".", err)
	}

	err = json.Unmarshal(b, &_config)
	if err != nil {
		loger.ForLog("Error, incorrect \"config.json\".", err)
	}
}

func GetBotConfig() config {
	return _config
}
