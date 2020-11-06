package betypes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"telegram-chat_bot/logger"
)

// Config struct for saving bot settings.
type Config struct {
	WebHook     string `json:"web_hook"`
	BotToken    string `json:"bot_token"`
	BotPort     string `json:"bot_port"`
	RedisConfig struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	} `json:"redis_config"`
	ChatsConfig struct {
		QueueSize  uint `json:"queue_size"`
		UsersCount uint `json:"users_count"`
	} `json:"chats_config"`
}

var config Config

func init() {
	b, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s. Failed to load \"config.json\".", err.Error()))
		panic(err)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s. Incorrect \"config.json\".", err.Error()))
		panic(err)
	}
}

// GetBotConfig return bot config.
func GetBotConfig() Config {
	return config
}
