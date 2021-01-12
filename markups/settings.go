package markups

import (
	"fmt"
	"strings"
	"telegram-chat_bot/betypes"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/logger"

	"github.com/go-redis/redis/v8"
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

	SixteenOrLessCallback           = "SIXTEEN_OR_LESS"
	FromSixteenToEighteenCallback   = "FROM_SIXTEEN_TO_EIGHTEEN"
	FromEighteenToTwentyOneCallback = "FROM_EIGHTEEN_TO_TWENTY_ONE"
)

// SettingsReplyMarkups' city constants.
const (
	OwnCityReplyMarkupName = "OWN CITY"
	OwnCityCallback        = OwnCityReplyMarkupName
	OwnCityPrefix          = OwnCityCallback + "_"

	InterlocutorCityReplyMarkupName = "INTERLOCUTOR CITY"
	InterlocutorCityCallback        = InterlocutorCityReplyMarkupName
	InterlocutorCityPrefix          = InterlocutorCityCallback + "_"
)

var settings = ReplyMarkups{
	ReplyMarkup{
		Name: SettingsPrefix + SettingsReplyMarkupName,
		InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("OWN AGE",
					SettingsPrefix+OwnAgeCallback),
				tgbotapi.NewInlineKeyboardButtonData("INTERLOCUTOR AGE",
					SettingsPrefix+InterlocutorAgeCallback),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("OWN CITY",
					SettingsPrefix+OwnCityCallback),
				tgbotapi.NewInlineKeyboardButtonData("INTERLOCUTOR CITY",
					SettingsPrefix+InterlocutorCityCallback),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("CLOSE", SettingsPrefix+CloseCallback),
			),
		}},
	},
	ReplyMarkup{
		Name: SettingsPrefix + OwnAgeReplyMarkupName,
		InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("SIXTEEN OR LESS",
					SettingsPrefix+OwnAgePrefix+SixteenOrLessCallback),
				tgbotapi.NewInlineKeyboardButtonData("FROM SIXTEEN TO EIGHTEEN",
					SettingsPrefix+OwnAgePrefix+FromSixteenToEighteenCallback),
				tgbotapi.NewInlineKeyboardButtonData("FROM EIGHTEEN TO TWENTY ONE",
					SettingsPrefix+OwnAgePrefix+FromEighteenToTwentyOneCallback),
				tgbotapi.NewInlineKeyboardButtonData("DOES NOT MATTER",
					SettingsPrefix+OwnAgePrefix+betypes.UserNil),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("<-BACK", SettingsPrefix+SettingsReplyMarkupName),
			),
		}},
	},
	ReplyMarkup{
		Name: SettingsPrefix + InterlocutorAgeReplyMarkupName,
		InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("SIXTEEN OR LESS",
					SettingsPrefix+InterlocutorAgePrefix+SixteenOrLessCallback),
				tgbotapi.NewInlineKeyboardButtonData("FROM SIXTEEN TO EIGHTEEN",
					SettingsPrefix+InterlocutorAgePrefix+FromSixteenToEighteenCallback),
				tgbotapi.NewInlineKeyboardButtonData("FROM EIGHTEEN TO TWENTY ONE",
					SettingsPrefix+InterlocutorAgePrefix+FromEighteenToTwentyOneCallback),
				tgbotapi.NewInlineKeyboardButtonData("DOES NOT MATTER",
					SettingsPrefix+InterlocutorAgePrefix+betypes.UserNil),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("<-BACK", SettingsPrefix+SettingsReplyMarkupName),
			),
		}},
	},
	ReplyMarkup{
		Name: SettingsPrefix + OwnCityReplyMarkupName,
		InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("LVIV",
					SettingsPrefix+OwnCityPrefix+"LVIV"),
				tgbotapi.NewInlineKeyboardButtonData("TERNOPIL",
					SettingsPrefix+OwnCityPrefix+"TERNOPIL"),
				tgbotapi.NewInlineKeyboardButtonData("KIEV",
					SettingsPrefix+OwnCityPrefix+"KIEV"),
				tgbotapi.NewInlineKeyboardButtonData("DOES NOT MATTER",
					SettingsPrefix+OwnCityPrefix+betypes.UserNil),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("<-BACK", SettingsPrefix+SettingsReplyMarkupName),
			),
		}},
	},
	{
		Name: SettingsPrefix + InterlocutorCityReplyMarkupName,
		InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("LVIV",
					SettingsPrefix+InterlocutorCityPrefix+"LVIV"),
				tgbotapi.NewInlineKeyboardButtonData("TERNOPIL",
					SettingsPrefix+InterlocutorCityPrefix+"TERNOPIL"),
				tgbotapi.NewInlineKeyboardButtonData("KIEV",
					SettingsPrefix+InterlocutorCityPrefix+"KIEV"),
				tgbotapi.NewInlineKeyboardButtonData("DOES NOT MATTER",
					SettingsPrefix+InterlocutorCityPrefix+betypes.UserNil),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("<-BACK", SettingsPrefix+SettingsReplyMarkupName),
			),
		}},
	},
}

// AnswerCallbackQuerySettings responds to a callback that has a settings(SettingsPrefix) prefix.
//
// If the callback query is for close, the reply markup is closing.
// Example callback query data: "SettingsPrefix+CloseCallback".
// If a callback query for a new reply markup, sends it.
// Example callback query data: "SettingsPrefix+OwnAgeCallback".
// If there is a callback query to change the age or city, etc... it is changing.
// Example callback query data: "SettingsPrefix+OwnAgePrefix+FromEighteenToTwentyOneCallback".
func AnswerCallbackQuerySettings(callbackQueryID, callbackQueryData string,
	callbackQueryMessageID, userID int, bot *tgbotapi.BotAPI) {
	if func() bool {
		if strings.EqualFold(strings.Replace(callbackQueryData, SettingsPrefix, "", 1), CloseCallback) {
			msg := tgbotapi.NewDeleteMessage(int64(userID), callbackQueryMessageID)

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}

			return true
		}
		return false
	}() {
		callback := tgbotapi.NewCallback(callbackQueryID, betypes.GetTexts().ReplyKeyboardMarkup.Closed)

		_, err := bot.AnswerCallbackQuery(callback)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		return
	}

	if func() bool {
		markup := settings.FindReplyMarkup(callbackQueryData)
		if markup != nil {
			msg := tgbotapi.NewEditMessageReplyMarkup(int64(userID), callbackQueryMessageID, markup.InlineKeyboardMarkup)

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}

			return true
		}

		return false
	}() {
		msg := tgbotapi.NewCallback(callbackQueryID, betypes.GetTexts().ReplyKeyboardMarkup.Opened)

		_, err := bot.AnswerCallbackQuery(msg)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		return
	}

	if func() bool {
		user, err := database.DB.GetUser(userID)
		if (err != nil) && (err.Error() == redis.Nil.Error()) {
			msg := tgbotapi.MessageConfig{
				BaseChat: tgbotapi.BaseChat{ChatID: int64(userID)},
				Text:     betypes.GetTexts().Chat.NotRegistered, ParseMode: betypes.GetTexts().ParseMode,
			}

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}

			return false
		}

		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		callbackQueryData = strings.Replace(callbackQueryData, SettingsPrefix, "", 1)
		if strings.Contains(callbackQueryData, OwnAgePrefix) {
			callbackQueryData = strings.Replace(callbackQueryData, OwnAgePrefix, "", 1)
			if user.Age == callbackQueryData {
				return true
			}

			user.Age = callbackQueryData
		} else if strings.Contains(callbackQueryData, InterlocutorAgePrefix) {
			callbackQueryData = strings.Replace(callbackQueryData, InterlocutorAgePrefix, "", 1)
			if user.InterlocutorAge == callbackQueryData {
				return true
			}

			user.InterlocutorAge = callbackQueryData
		} else if strings.Contains(callbackQueryData, OwnCityPrefix) {
			callbackQueryData = strings.Replace(callbackQueryData, OwnCityPrefix, "", 1)
			if user.City == callbackQueryData {
				return true
			}

			user.City = callbackQueryData
		} else if strings.Contains(callbackQueryData, InterlocutorCityPrefix) {
			callbackQueryData = strings.Replace(callbackQueryData, InterlocutorCityPrefix, "", 1)
			if user.InterlocutorCity == callbackQueryData {
				return true
			}

			user.InterlocutorCity = callbackQueryData
		} else {
			return false
		}

		err = database.DB.SaveUser(*user)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		return true
	}() {
		_, err := bot.AnswerCallbackQuery(
			tgbotapi.NewCallback(callbackQueryID, betypes.GetTexts().ReplyKeyboardMarkup.Changed))
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		return
	}
}

// GetSettings return settings(ReplyMarkup).
func GetSettings() ReplyMarkups {
	return settings
}
