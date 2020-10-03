package actions

import (
	"errors"
	"fmt"
	"strings"
	"telegram-chat_bot/loger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Keyboard stores the keyboard and keyboard name.
type Keyboard struct {
	Name     string                        `json:"name"`
	Keyboard tgbotapi.InlineKeyboardMarkup `json:"keyboard"`
}

var keyboards = []Keyboard{
	{
		Name: "settings",
		Keyboard: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
			//TODO:	tgbotapi.NewInlineKeyboardButtonData("Age", "age"),
			//TODO:	tgbotapi.NewInlineKeyboardButtonData("City", "city"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Close", "close"),
			),
		}},
	},
}

// SettingsCommand sends the user a keyboard with settings.
func SettingsCommand(chatID int64, bot *tgbotapi.BotAPI) {
	keyboard, err := getKeyboard("settings")
	if err != nil {
		loger.ForLog(fmt.Sprintf("Error %v. Chat ID", err))
		return
	}

	if _, err := bot.Send(tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:      chatID,
			ReplyMarkup: (*keyboard).Keyboard,
		},
		ParseMode: "MARKDOWN",
		Text:      fmt.Sprintf("*%s*", strings.ToTitle(keyboard.Name)),
	}); err != nil {
		loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %v", err, chatID))
	}
}

// CloseReplyMarkupCommand delete markup from chat.
func CloseReplyMarkupCommand(chatID int64, messageID int, bot *tgbotapi.BotAPI) {
	if _, err := bot.Send(tgbotapi.EditMessageReplyMarkupConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:    chatID,
			MessageID: messageID,
			ReplyMarkup: &tgbotapi.InlineKeyboardMarkup{
				InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0),
			},
		},
	}); err != nil {
		loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %v", err, chatID))
	}
}

func getKeyboard(keyboardName string) (*Keyboard, error) {
	for _, i := range keyboards {
		if i.Name == keyboardName {
			return &i, nil
		}
	}
	return nil, errors.New("keyboard not found")
}
