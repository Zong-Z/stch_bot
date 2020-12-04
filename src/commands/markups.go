package commands

import (
	"telegram-chat_bot/betypes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	// SettingsPrefixReplyMarkup If callbackquery is a setting it has this prefix.
	SettingsPrefixReplyMarkup = "SETTINGS_"
	agePrefixReplyMarkup      = "AGE_"
	cityPrefixReplyMarkup     = "CITY_"
)

type replyMarkup struct {
	Name                 string
	InlineKeyboardMarkup tgbotapi.InlineKeyboardMarkup
}

type replyMarkups []replyMarkup

var settingsReplyMarkups = replyMarkups{
	replyMarkup{
		Name: SettingsPrefixReplyMarkup + "SETTINGS",
		InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("AGE", SettingsPrefixReplyMarkup+"age"),
				tgbotapi.NewInlineKeyboardButtonData("CITY", SettingsPrefixReplyMarkup+"city"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("CLOSE", SettingsPrefixReplyMarkup+"close"),
			),
		}},
	},
	replyMarkup{
		Name: SettingsPrefixReplyMarkup + "AGE",
		InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("SIXTEEN OR LESS",
					SettingsPrefixReplyMarkup+agePrefixReplyMarkup+"sixteen_or_less"),
				tgbotapi.NewInlineKeyboardButtonData("FROM EIGHTEEN TO TWENTY-ONE",
					SettingsPrefixReplyMarkup+agePrefixReplyMarkup+"from_eighteen_to_twenty-one"),
				tgbotapi.NewInlineKeyboardButtonData("NOT SPECIFIED",
					SettingsPrefixReplyMarkup+agePrefixReplyMarkup+betypes.UserNil),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("<-GO BACK", SettingsPrefixReplyMarkup+"SETTINGS"),
			),
		}},
	},
	replyMarkup{
		Name: SettingsPrefixReplyMarkup + "CITY",
		InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("LVIV",
					SettingsPrefixReplyMarkup+cityPrefixReplyMarkup+"lviv"),
				tgbotapi.NewInlineKeyboardButtonData("TERNOPIL",
					SettingsPrefixReplyMarkup+cityPrefixReplyMarkup+"ternopil"),
				tgbotapi.NewInlineKeyboardButtonData("NOT SPECIFIED",
					SettingsPrefixReplyMarkup+cityPrefixReplyMarkup+betypes.UserNil),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("<-GO BACK", SettingsPrefixReplyMarkup+"SETTINGS"),
			),
		}},
	},
}
