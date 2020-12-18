package commands

import (
	"fmt"
	"strings"
	"telegram-chat_bot/betypes"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/logger"

	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Help sends help text to user.
func Help(chatID int, bot *tgbotapi.BotAPI) {
	_, err := bot.Send(tgbotapi.NewMessage(int64(chatID), betypes.GetBotCommands().Help.Text))
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}
}

// Start display the start text save the user in the database.
func Start(user betypes.User, bot *tgbotapi.BotAPI) {
	// Saving user.
	_, err := database.DB.GetUser(int64(user.ID))
	if err != nil && err.Error() == redis.Nil.Error() /* If no user is found */ {
		err := database.DB.SaveUser(user)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	}

	// Sending to user start text.
	_, err = bot.Send(tgbotapi.NewMessage(int64(user.ID), betypes.GetBotCommands().Start.Text))
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}
}

// Settings sends to the user settings keyboard.
func Settings(chatID int64, bot *tgbotapi.BotAPI) {
	markup := settingsReplyMarkups.findReplyMarkup(SettingsPrefixReplyMarkup + "SETTINGS")
	if markup == nil {
		return
	}

	_, err := bot.Send(tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{ChatID: chatID, ReplyMarkup: markup.InlineKeyboardMarkup},
		Text:     fmt.Sprintf("*%s*", markup.Name), ParseMode: "MARKDOWN"})
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}
}

// AnswerOnCallbackQuerySettings callback query answers concerning setting.
//
// If the callback query is for close, the reply markup is closing.
// SettingsPrefixReplyMarkup+"close".
//
// If a callback query for a new reply markup, sends it.
// SettingsPrefixReplyMarkup+"age", etc...
//
// If there is a callback query to change the age or city, etc... it is changed.
// SettingsPrefixReplyMarkup+agePrefixReplyMarkup+"user: nil", etc...
func AnswerOnCallbackQuerySettings(callbackQuery tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	// The anonymous function returns true if
	// callbackQuery.Data is for close the reply markup.
	if func() bool {
		if strings.EqualFold(strings.Replace(callbackQuery.Data, SettingsPrefixReplyMarkup, "", 1), "close") {
			_, err := bot.Send(tgbotapi.NewDeleteMessage(int64(callbackQuery.From.ID), callbackQuery.Message.MessageID))
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s. Chat ID - %d.", err.Error(), callbackQuery.From.ID))
				panic(err)
			}

			return true
		}

		return false
	}() {
		_, err := bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuery.ID, "CLOSED"))
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s. Chat ID - %d.", err.Error(), callbackQuery.From.ID))
			panic(err)
		}

		return
	}

	// If callback query is for new reply markup.
	if replaceReplyMarkup(callbackQuery, bot) {
		_, err := bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuery.ID, "OPENED"))
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s. Chat ID - %d.", err.Error(), callbackQuery.From.ID))
			panic(err)
		}

		return
	}

	if changeUserAgeOrCity(callbackQuery, bot) {
		_, err := bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuery.ID, "CHANGED"))
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s. Chat ID - %d.", err.Error(), callbackQuery.From.ID))
			panic(err)
		}

		return
	}
}

// replaceReplyMarkup return true if reply markup was replaced.
func replaceReplyMarkup(callbackQuery tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) bool {
	replyMarkup := settingsReplyMarkups.findReplyMarkup(strings.ToTitle(callbackQuery.Data))
	if replyMarkup != nil {
		_, err := bot.Send(tgbotapi.NewEditMessageReplyMarkup(
			int64(callbackQuery.From.ID), callbackQuery.Message.MessageID, replyMarkup.InlineKeyboardMarkup))
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s. Chat ID - %d.", err.Error(), callbackQuery.From.ID))
			panic(err)
		}

		return true
	}

	return false
}

// changeUserAgeOrCity if the callback request concerns a change of age or city,
// it will replace the old user data.
func changeUserAgeOrCity(callbackQuery tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) bool {
	u, err := database.DB.GetUser(int64(callbackQuery.From.ID))
	if err != nil && err.Error() == redis.Nil.Error() {
		_, err := bot.Send(tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{ChatID: int64(callbackQuery.From.ID)},
			Text:     "*YOU ARE NOT REGISTERED* \"/start\"", ParseMode: "MARKDOWN"})
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

	callbackQueryData := strings.Replace(callbackQuery.Data, SettingsPrefixReplyMarkup, "", 1)
	if strings.Contains(callbackQueryData, agePrefixReplyMarkup) {
		callbackQueryData = strings.Replace(callbackQueryData, agePrefixReplyMarkup, "", 1)
		if u.Age == callbackQueryData {
			return true
		}

		u.Age = callbackQueryData
	} else if strings.Contains(callbackQueryData, cityPrefixReplyMarkup) {
		callbackQueryData = strings.Replace(callbackQueryData, cityPrefixReplyMarkup, "", 1)
		if u.City == callbackQueryData {
			return true
		}

		u.City = callbackQueryData
	}

	err = database.DB.SaveUser(*u)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}

	return true
}
