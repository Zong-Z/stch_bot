package markups

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type (
	// Markup struct for saving InlineKeyboardMarkup
	Markup struct {
		Name                 string                        `json:"name"`
		InlineKeyboardMarkup tgbotapi.InlineKeyboardMarkup `json:"inline_keyboard_markup"`
	}

	// Markups reply markups array([]Markup)
	Markups []Markup
)

// Markup close callback.
const (
	CloseCallback = "CLOSE"
	DoesNotMatter = "DOES NOT MATTER"
	Back          = "<-BACK"
)

// IsThereCloseCallback if the callback contains CloseCallback, returns true.
func IsThereCloseCallback(callbackQueryData string) bool {
	return strings.Contains(callbackQueryData, CloseCallback)
}

// FindInlineKeyboard return *tgbotapi.InlineKeyboardMarkup by reply markup name.
//
// Return nil if reply markup do not found.
func (m Markups) FindInlineKeyboard(replyMarkupName string) *tgbotapi.InlineKeyboardMarkup {
	for _, markup := range m {
		if !strings.EqualFold(markup.Name, replyMarkupName) {
			continue
		}

		newButtons := make([][]tgbotapi.InlineKeyboardButton, len(markup.InlineKeyboardMarkup.InlineKeyboard))
		for i, buttons := range markup.InlineKeyboardMarkup.InlineKeyboard {
			for _, button := range buttons {
				newButtons[i] = append(newButtons[i], tgbotapi.InlineKeyboardButton{Text: button.Text, URL: button.URL,
					CallbackData: button.CallbackData, SwitchInlineQuery: button.SwitchInlineQuery,
					SwitchInlineQueryCurrentChat: button.SwitchInlineQueryCurrentChat, CallbackGame: button.CallbackGame,
					Pay: button.Pay})
			}
		}

		return &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: newButtons}
	}

	return nil
}
