package betypes

import (
	"encoding/json"
	"io/ioutil"
	"telegram-chat_bot/loger"
)

// Commands struct for saving bot commands.
type Commands struct {
	Start struct {
		Command string `json:"command"`
		Text    string `json:"text"`
	} `json:"start"`
	Help struct {
		Command string `json:"command"`
		Text    string `json:"text"`
	} `json:"help"`
	Chatting struct {
		Start string `json:"start"`
		Stop  string `json:"stop"`
	} `json:"chatting"`
	Settings struct {
		Command string `json:"command"`
	} `json:"settings"`
}

var commands Commands

func init() {
	b, err := ioutil.ReadFile("config/commands.json")
	if err != nil {
		loger.ForLog("Error, failed to load \"commands.json\".", err)
	}

	err = json.Unmarshal(b, &commands)
	if err != nil {
		loger.ForLog("Error, incorrect bot \"commands.json\".", err)
	}
}

// GetBotCommands return bot commands.
func GetBotCommands() Commands {
	return commands
}
