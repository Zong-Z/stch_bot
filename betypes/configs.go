package betypes

import (
	"encoding/json"
	"io/ioutil"
	"telegram-chat_bot/loger"
)

var Config config

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

func init() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		loger.LogFile.Fatalln("Error, failed to load \"config.json\".", err)
	}

	err = json.Unmarshal(b, &Config)
	if err != nil {
		loger.LogFile.Fatalln("Error, incorrect \"config.json\".", err)
	}
}
