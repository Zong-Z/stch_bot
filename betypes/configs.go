package betypes

import (
	"encoding/json"
	"io/ioutil"
	"telegram-chat_bot/loger"
)

// Сonfig struct for saving bot settings.
type Сonfig struct {
	WebHook     string `json:"web_hook"`
	BotToken    string `json:"bot_token"`
	BotPort     string `json:"bot_port"`
	RedisConfig struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	} `json:"redis_config"`
}

var config Сonfig

func init() {
	b, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		loger.ForLog("Error, failed to load \"config.json\".", err)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		loger.ForLog("Error, incorrect \"config.json\".", err)
	}
}

// GetBotConfig return bot config.
func GetBotConfig() Сonfig {
	return config
}
