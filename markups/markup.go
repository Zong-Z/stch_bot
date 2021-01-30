package markups

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type (
	// ReplyMarkup struct for saving InlineKeyboardMarkup
	ReplyMarkup struct {
		Name                 string                        `json:"name"`
		InlineKeyboardMarkup tgbotapi.InlineKeyboardMarkup `json:"inline_keyboard_markup"`
	}

	// ReplyMarkups reply markups array([]ReplyMarkup)
	ReplyMarkups []ReplyMarkup
)

// ReplyMarkup close callback.
const (
	CloseCallback = "CLOSE"
	DoesNotMatter = "DOES NOT MATTER"
	Back          = "<-BACK"
)

// FindInlineKeyboardMarkup return *tgbotapi.InlineKeyboardMarkup by reply markup name.
//
// Return nil if reply markup do not found.
func (m ReplyMarkups) FindInlineKeyboardMarkup(replyMarkupName string) *tgbotapi.InlineKeyboardMarkup {
	for _, markup := range m {
		if !strings.EqualFold(markup.Name, replyMarkupName) {
			continue
		}

		buttons := make([][]tgbotapi.InlineKeyboardButton, len(markup.InlineKeyboardMarkup.InlineKeyboard))
		for i := 0; i < len(markup.InlineKeyboardMarkup.InlineKeyboard); i++ {
			for j := 0; j < len(markup.InlineKeyboardMarkup.InlineKeyboard[i]); j++ {
				buttons[i] = append(buttons[i], tgbotapi.InlineKeyboardButton{
					Text:                         markup.InlineKeyboardMarkup.InlineKeyboard[i][j].Text,
					URL:                          markup.InlineKeyboardMarkup.InlineKeyboard[i][j].URL,
					CallbackData:                 markup.InlineKeyboardMarkup.InlineKeyboard[i][j].CallbackData,
					SwitchInlineQuery:            markup.InlineKeyboardMarkup.InlineKeyboard[i][j].SwitchInlineQuery,
					SwitchInlineQueryCurrentChat: markup.InlineKeyboardMarkup.InlineKeyboard[i][j].SwitchInlineQueryCurrentChat,
					CallbackGame:                 markup.InlineKeyboardMarkup.InlineKeyboard[i][j].CallbackGame,
					Pay:                          markup.InlineKeyboardMarkup.InlineKeyboard[i][j].Pay,
				})
			}
		}

		return &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: buttons}
	}

	return nil
}
