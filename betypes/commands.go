package betypes

import (
	"encoding/json"
	"io/ioutil"
	"telegram-chat_bot/loger"
)

type commands struct {
	Start struct {
		Command string `json:"command"`
		Text    string `json:"text"`
	}
	Help struct {
		Command string `json:"command"`
		Text    string `json:"text"`
	}
	Settings struct {
		Command string `json:"command"`
		Text    string `json:"text"`
	}
}

var _commands commands

func init() {
	b, err := ioutil.ReadFile("commands.json")
	if err != nil {
		loger.LogFile.Fatalln("Error, failed to load \"commands.json\".", err)
	}

	err = json.Unmarshal(b, &_commands)
	if err != nil {
		loger.LogFile.Fatalln("Error, incorrect bot \"commands.json\".", err)
	}
}

func GetStartCommand() string {
	return _commands.Start.Command
}

func GetHelpCommand() string {
	return _commands.Help.Command
}

func GetSettingsCommand() string {
	return _commands.Settings.Command
}

func GetStartText() string {
	return _commands.Start.Text
}

func GetHelpText() string {
	return _commands.Help.Text
}

func GetSettingsText() string {
	return _commands.Settings.Text
}
