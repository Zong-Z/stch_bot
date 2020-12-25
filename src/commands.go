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
	_, err := bot.Send(tgbotapi.NewMessage(int64(chatID), betypes.GetTexts().Commands.Help.Text))
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}
}

// Start display the start text save the user in the database.
func Start(user betypes.User, bot *tgbotapi.BotAPI) {
	// Saving user.
	_, err := database.DB.GetUser(int64(user.ID))
	if err != nil && err.Error() == redis.Nil.Error() {
		err := database.DB.SaveUser(user)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	}

	// Sending to user start text.
	_, err = bot.Send(tgbotapi.NewMessage(int64(user.ID), betypes.GetTexts().Commands.Start.Text))
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}
}

// StartChat add user to chat queue.
func StartChat(message tgbotapi.Message, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{ChatID: int64(message.From.ID)},
		Text:     betypes.GetTexts().Chat.AlreadyInChat, ParseMode: betypes.GetTexts().ParseMode,
	}

	if !chats.UserIsChatting(message.From.ID) {
		u, err := database.DB.GetUser(int64(message.From.ID))
		if err != nil && err.Error() == redis.Nil.Error() {
			logger.ForLog(fmt.Sprintf("User not found. User ID - %d.", int64(message.From.ID)))

			msg.Text = betypes.GetTexts().Chat.NotRegistered
			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}
		}

		if err != nil {
			logger.ForLog(fmt.Sprintf("Error, %s.", err.Error()))
			panic(err)
		}

		chats.AddUserToTheQueue(*u)

		msg.Text = betypes.GetTexts().Chat.InterlocutorSearchStarted
	}

	_, err := bot.Send(msg)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}

	chat := chats.GetChatByUserID(message.From.ID)
	if chat != nil {
		if int64(len(chat.Users)) == int64(chats.UsersCount) {
			msg = tgbotapi.MessageConfig{
				Text:      betypes.GetTexts().Chat.ChatFound,
				ParseMode: betypes.GetTexts().ParseMode,
			}
			for _, user := range chat.Users {
				msg.ChatID = int64(user.ID)
				_, err = bot.Send(msg)
				if err != nil {
					logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
					panic(err)
				}
			}
		}
	}
}

// StopChat ends the chat.
func StopChat(message tgbotapi.Message, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	chat := chats.GetChatByUserID(message.From.ID)
	if chat != nil {
		chats.StopChatting(message.From.ID)

		for _, user := range chat.Users {
			msg := tgbotapi.MessageConfig{
				BaseChat: tgbotapi.BaseChat{ChatID: int64(user.ID)},
				Text:     betypes.GetTexts().Chat.ChatEnded, ParseMode: betypes.GetTexts().ParseMode,
			}

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}
		}
	}

	logger.ForLog(fmt.Sprintf("User ID - %d. Chat did not found.", message.From.ID))
}

// Settings sends to the user settings keyboard.
func Settings(chatID int64, bot *tgbotapi.BotAPI) {
	markup := settingsReplyMarkups.findReplyMarkup(SettingsPrefixReplyMarkup + "SETTINGS")
	if markup == nil {
		return
	}

	_, err := bot.Send(tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{ChatID: chatID, ReplyMarkup: markup.InlineKeyboardMarkup},
		Text: fmt.Sprintf("*%s*",
			strings.Replace(markup.Name, SettingsPrefixReplyMarkup, "", 1)),
		ParseMode: betypes.GetTexts().ParseMode,
	})
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
			Text:     betypes.GetTexts().Chat.NotRegistered, ParseMode: betypes.GetTexts().ParseMode,
		})
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
