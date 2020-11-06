package betypes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"telegram-chat_bot/logger"
)

// Commands struct for saving bot actions.
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
	Unknown struct {
		Text string `json:"text"`
	} `json:"unknown"`
}

var commands Commands

func init() {
	b, err := ioutil.ReadFile("config/commands.json")
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s. Failed to load \"commands.json\".", err.Error()))
		panic(err)
	}

	err = json.Unmarshal(b, &commands)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s. Incorrect \"commands.json\".", err.Error()))
		panic(err)
	}
}

// GetBotCommands return bot actions.
func GetBotCommands() Commands {
	return commands
}
