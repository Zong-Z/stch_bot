package betypes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"telegram-chat_bot/logger"
)

// Text struct for saving texts and bot commands.
type Text struct {
	Commands struct {
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
	} `json:"commands"`
	ParseMode string `json:"parse_mode"`
	Chat      struct {
		NotRegistered             string `json:"not_registered"`
		InterlocutorSearchStarted string `json:"interlocutor_search_begun"`
		ChatFound                 string `json:"chat_found"`
		AlreadyInChat             string `json:"already_in_chat"`
		NotInChat                 string `json:"not_in_chat"`
		ChatEnded                 string `json:"chat_ended"`
	} `json:"chat"`
	ReplyKeyboardMarkup struct {
		Opened  string `json:"opened"`
		Changed string `json:"changed"`
		Closed  string `json:"closed"`
	} `json:"reply_keyboard_markup"`
}

var text Text

func init() {
	b, err := ioutil.ReadFile("configs/texts.json")
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
	}

	err = json.Unmarshal(b, &text)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
	}
}

// GetTexts return texts and bot commands.
func GetTexts() Text {
	return text
}
