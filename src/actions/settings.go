package actions

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram-chat_bot/loger"
)

type callback struct {
	Text string
	Data string
}

type keyboard struct {
	Name     string
	Callback []callback
	Keyboard tgbotapi.InlineKeyboardMarkup
}

type keyboards []keyboard

var _keyboards = keyboards{
	keyboard{
		Name: "Age",
		Callback: []callback{
			{
				Text: "Less than 16",
				Data: "LessThanSixteen",
			},
			{
				Text: "16-18",
				Data: "SixteenEighteen",
			},
			{
				Text: "More than 21",
				Data: "MoreThanTwentyOne",
			},
			{
				Text: "Not specifying",
				Data: "NotSpecifying",
			},
			{
				Text: "<-Go back",
				Data: "GoBack",
			},
		},
		Keyboard: tgbotapi.InlineKeyboardMarkup{},
	},
	keyboard{
		Name: "City",
		Callback: []callback{
			{
				Text: "<-Go back",
				Data: "GoBack",
			},
		},
		Keyboard: tgbotapi.InlineKeyboardMarkup{},
	},
	keyboard{
		Name: "Settings",
		Callback: []callback{
			{
				Text: "Age",
				Data: "Age",
			},
			{
				Text: "City",
				Data: "City",
			},
			{
				Text: "Close",
				Data: "Close",
			},
		},
		Keyboard: tgbotapi.InlineKeyboardMarkup{},
	},
}

func init() {
	for i := 0; i < len(_keyboards); i++ {
		for j := 0; j < len(_keyboards[i].Callback); j++ {
			_keyboards[i].Keyboard.InlineKeyboard = append(_keyboards[i].Keyboard.InlineKeyboard,
				tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
					_keyboards[i].Callback[j].Text, _keyboards[i].Callback[j].Data)))
		}
	}
}

func SettingsCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery != nil {
		for _, markup := range _keyboards {
			if update.CallbackQuery.Data == markup.Name {
				if _, err := bot.Send(tgbotapi.EditMessageReplyMarkupConfig{
					BaseEdit: struct {
						ChatID          int64
						ChannelUsername string
						MessageID       int
						InlineMessageID string
						ReplyMarkup     *tgbotapi.InlineKeyboardMarkup
					}{
						ChatID:      update.CallbackQuery.Message.Chat.ID,
						MessageID:   update.CallbackQuery.Message.MessageID,
						ReplyMarkup: &markup.Keyboard,
					}}); err != nil {
					loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %v",
						err, update.CallbackQuery.From.ID))
				}
				return
			}
		}
		return
	}

	for _, markup := range _keyboards {
		if markup.Name == "Settings" {
			if _, err := bot.Send(tgbotapi.MessageConfig{
				BaseChat: tgbotapi.BaseChat{
					ChatID:      int64(update.Message.From.ID),
					ReplyMarkup: markup.Keyboard,
				},
				Text:      fmt.Sprintf("*%s*", markup.Name),
				ParseMode: "MARKDOWN",
			}); err != nil {
				loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %v",
					err, update.Message.From.ID))
			}
			return
		}
	}
}
