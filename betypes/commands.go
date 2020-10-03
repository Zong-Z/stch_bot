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
	Settings struct {
		Command string `json:"command"`
		Text    string `json:"text"`
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
