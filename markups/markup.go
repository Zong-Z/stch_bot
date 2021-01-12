package markups

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type (
	// ReplyMarkup struct for saving InlineKeyboardMarkup
	ReplyMarkup struct {
		Name                 string                        `json:"name"`
		InlineKeyboardMarkup tgbotapi.InlineKeyboardMarkup `json:"inline_keyboard_markup"`
	}

	// ReplyMarkups reply markups array([][]ReplyMarkup)
	ReplyMarkups []ReplyMarkup
)

// ReplyMarkup close callback.
const (
	CloseCallback = "CLOSE"
)

// FindReplyMarkup return ReplyMarkup by reply markup name.
//
// Return "nil" if reply markup do not found.
func (m ReplyMarkups) FindReplyMarkup(replyMarkupName string) *ReplyMarkup {
	for i := 0; i < len(m); i++ {
		if m[i].Name == replyMarkupName {
			return &m[i]
		}
	}
	return nil
}
