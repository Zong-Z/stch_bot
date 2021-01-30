package markups

import (
	"strings"
	"telegram-chat_bot/betypes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SettingsReplyMarkups' constants.
const (
	SettingsReplyMarkupName = "SETTINGS"
	SettingsPrefix          = SettingsReplyMarkupName + "_"
)

// SettingsReplyMarkups' age constants.
const (
	OwnAgeReplyMarkupName = "OWN AGE"
	OwnAgeCallback        = OwnAgeReplyMarkupName
	OwnAgePrefix          = OwnAgeCallback + "_"

	InterlocutorAgeReplyMarkupName = "INTERLOCUTOR AGE"
	InterlocutorAgeCallback        = InterlocutorAgeReplyMarkupName
	InterlocutorAgePrefix          = InterlocutorAgeCallback + "_"

	SixteenOrLessCallback           = "SIXTEEN OR LESS"
	FromSixteenToEighteenCallback   = "FROM SIXTEEN TO EIGHTEEN"
	FromEighteenToTwentyOneCallback = "FROM EIGHTEEN TO TWENTY ONE"

	OwnAgeText                  = "OWN AGE"
	InterlocutorAgeText         = "INTERLOCUTOR AGE"
	SixteenOrLessText           = "SIXTEEN OR LESS"
	FromSixteenToEighteenText   = "FROM SIXTEEN TO EIGHTEEN"
	FromEighteenToTwentyOneText = "FROM EIGHTEEN TO TWENTY ONE"
)

// SettingsReplyMarkups' city constants.
const (
	OwnCityReplyMarkupName = "OWN CITY"
	OwnCityCallback        = OwnCityReplyMarkupName
	OwnCityPrefix          = OwnCityCallback + "_"

	InterlocutorCityReplyMarkupName = "INTERLOCUTOR CITY"
	InterlocutorCityCallback        = InterlocutorCityReplyMarkupName
	InterlocutorCityPrefix          = InterlocutorCityCallback + "_"

	OwnCityText          = "OWN CITY"
	InterlocutorCityText = "INTERLOCUTOR CITY"
)

// SettingsReplyMarkups' cities.
const (
	Lviv     = "LVIV"
	Ternopil = "TERNOPIL"
	Kiev     = "KIEV"
)

var (
	ownAgeInlineKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(SixteenOrLessText,
					SettingsPrefix+OwnAgePrefix+SixteenOrLessCallback),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(FromSixteenToEighteenText,
					SettingsPrefix+OwnAgePrefix+FromSixteenToEighteenCallback),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(FromEighteenToTwentyOneText,
					SettingsPrefix+OwnAgePrefix+FromEighteenToTwentyOneCallback),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(DoesNotMatter,
					SettingsPrefix+OwnAgePrefix+betypes.UserNil),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Back, SettingsPrefix+SettingsReplyMarkupName),
			),
		},
	}
	interlocutorAgeInlineKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(SixteenOrLessText,
					SettingsPrefix+InterlocutorAgePrefix+SixteenOrLessCallback),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(FromSixteenToEighteenText,
					SettingsPrefix+InterlocutorAgePrefix+FromSixteenToEighteenCallback),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(FromEighteenToTwentyOneText,
					SettingsPrefix+InterlocutorAgePrefix+FromEighteenToTwentyOneCallback),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(DoesNotMatter,
					SettingsPrefix+InterlocutorAgePrefix+betypes.UserNil),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Back, SettingsPrefix+SettingsReplyMarkupName),
			),
		},
	}
	ownCityInlineKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Lviv,
					SettingsPrefix+OwnCityPrefix+Lviv,
				),
				tgbotapi.NewInlineKeyboardButtonData(Ternopil,
					SettingsPrefix+OwnCityPrefix+Ternopil,
				),
				tgbotapi.NewInlineKeyboardButtonData(Kiev,
					SettingsPrefix+OwnCityPrefix+Kiev,
				),
				tgbotapi.NewInlineKeyboardButtonData(DoesNotMatter,
					SettingsPrefix+OwnCityPrefix+betypes.UserNil,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Back, SettingsPrefix+SettingsReplyMarkupName),
			),
		},
	}
	interlocutorCityInlineKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Lviv,
					SettingsPrefix+InterlocutorCityPrefix+Lviv,
				),
				tgbotapi.NewInlineKeyboardButtonData(Ternopil,
					SettingsPrefix+InterlocutorCityPrefix+Ternopil,
				),
				tgbotapi.NewInlineKeyboardButtonData(Kiev,
					SettingsPrefix+InterlocutorCityPrefix+Kiev,
				),
				tgbotapi.NewInlineKeyboardButtonData(DoesNotMatter,
					SettingsPrefix+InterlocutorCityPrefix+betypes.UserNil),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Back, SettingsPrefix+SettingsReplyMarkupName),
			),
		},
	}
)

var settings = ReplyMarkups{
	ReplyMarkup{
		Name: SettingsPrefix + SettingsReplyMarkupName,
		InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(OwnAgeText,
					SettingsPrefix+OwnAgeCallback),
				tgbotapi.NewInlineKeyboardButtonData(InterlocutorAgeText,
					SettingsPrefix+InterlocutorAgeCallback),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(OwnCityText,
					SettingsPrefix+OwnCityCallback),
				tgbotapi.NewInlineKeyboardButtonData(InterlocutorCityText,
					SettingsPrefix+InterlocutorCityCallback),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(CloseCallback, SettingsPrefix+CloseCallback),
			),
		}},
	},
	ReplyMarkup{
		Name:                 SettingsPrefix + OwnAgeReplyMarkupName,
		InlineKeyboardMarkup: ownAgeInlineKeyboardMarkup,
	},
	ReplyMarkup{
		Name:                 SettingsPrefix + InterlocutorAgeReplyMarkupName,
		InlineKeyboardMarkup: interlocutorAgeInlineKeyboardMarkup,
	},
	ReplyMarkup{
		Name:                 SettingsPrefix + OwnCityReplyMarkupName,
		InlineKeyboardMarkup: ownCityInlineKeyboardMarkup,
	},
	ReplyMarkup{
		Name:                 SettingsPrefix + InterlocutorCityReplyMarkupName,
		InlineKeyboardMarkup: interlocutorCityInlineKeyboardMarkup,
	},
}

// IsThereCloseCallback if the callback contains CloseCallback, returns true.
func IsThereCloseCallback(callbackQueryData string) bool {
	return strings.Contains(callbackQueryData, CloseCallback)
}

// SettingsIsThereMarkupRequest return true if callback contains OwnAgeReplyMarkupName,
// InterlocutorAgeReplyMarkupName, OwnCityReplyMarkupName, InterlocutorCityReplyMarkupName.
func SettingsIsThereMarkupRequest(callbackQueryData string) bool {
	return settings.FindInlineKeyboardMarkup(callbackQueryData) != nil
}

// SettingsIsThereCallbackForChange if there is a callback to change your own age/city or interlocutor, returns true.
func SettingsIsThereCallbackForChange(callbackQueryData string) bool {
	return strings.Contains(callbackQueryData, OwnAgePrefix) ||
		strings.Contains(callbackQueryData, InterlocutorAgePrefix) ||
		strings.Contains(callbackQueryData, OwnCityPrefix) ||
		strings.Contains(callbackQueryData, InterlocutorCityPrefix)
}

// GetSettings return settings(ReplyMarkup).
func GetSettings() ReplyMarkups {
	return settings
}
