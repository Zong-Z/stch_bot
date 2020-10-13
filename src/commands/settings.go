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
				/*//TODO:*/ tgbotapi.NewInlineKeyboardButtonData("Age", "age"),
				/*//TODO:*/ tgbotapi.NewInlineKeyboardButtonData("City", "city"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Close", "close"),
			),
		}},
	},
}

// SettingsCommandMarkup sends the user a keyboard with settings.
func SettingsCommandMarkup(chatID int64, bot *tgbotapi.BotAPI) {
	markup, err := getKeyboard("settings")
	if err != nil {
		loger.ForLog(fmt.Sprintf("Error %v. Chat ID %v.", err, chatID))
		return
	}

	if _, err := bot.Send(tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:      chatID,
			ReplyMarkup: markup.Keyboard,
		},
		ParseMode: "MARKDOWN",
		Text:      fmt.Sprintf("*%s*", strings.ToTitle(markup.Name)),
	}); err != nil {
		loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %v", err, chatID))
	}
}

// DeleteMessage delete message from chat.
func DeleteMessage(chatID int64, messageID int, bot *tgbotapi.BotAPI) {
	if _, err := bot.Send(tgbotapi.NewDeleteMessage(chatID, messageID)); err != nil {
		loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %v.", err, chatID))
	}
}

func getKeyboard(keyboardName string) (*Keyboard, error) {
	for i := 0; i < len(keyboards); i++ {
		if strings.EqualFold(keyboards[i].Name, keyboardName) {
			return &keyboards[i], nil
		}
	}
	return nil, errors.New("keyboard not found")
}
