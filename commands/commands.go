package commands

import (
	"fmt"
	"strings"
	"telegram-chat_bot/betypes"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/logger"
	"telegram-chat_bot/markups"

	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Start sends to user start text and save him to database.
func Start(user betypes.User, bot *tgbotapi.BotAPI) {
	_, err := database.DB.GetUser(user.ID)
	if err != nil && err.Error() == redis.Nil.Error() {
		err := database.DB.SaveUser(user)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	} else if err != nil && err.Error() != redis.Nil.Error() {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}

	msg := tgbotapi.MessageConfig{
		BaseChat:  tgbotapi.BaseChat{ChatID: int64(user.ID)},
		Text:      betypes.GetTexts().Commands.Start.Text,
		ParseMode: betypes.GetTexts().ParseMode,
	}

	_, err = bot.Send(msg)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}
}

// Help sends help text to the user.
func Help(chatID int, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.MessageConfig{
		BaseChat:  tgbotapi.BaseChat{ChatID: int64(chatID)},
		Text:      betypes.GetTexts().Commands.Help.Text,
		ParseMode: betypes.GetTexts().ParseMode,
	}

	_, err := bot.Send(msg)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}
}

// StartChatting starts a chat with a random user or users(See ).
func StartChatting(userID int, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.MessageConfig{
		BaseChat:  tgbotapi.BaseChat{ChatID: int64(userID)},
		Text:      betypes.GetTexts().Chat.AlreadyInChat,
		ParseMode: betypes.GetTexts().ParseMode,
	}

	if !chats.IsUserInChat(userID) {
		user, err := database.DB.GetUser(userID)
		if err != nil && err.Error() == redis.Nil.Error() || user == nil {
			msg.Text = betypes.GetTexts().Chat.NotRegistered
			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}

			return
		} else if err != nil {
			logger.ForLog(fmt.Sprintf("Error, %s.", err.Error()))
			panic(err)
		}

		chats.AddUserToQueue(*user)

		msg.Text = betypes.GetTexts().Chat.InterlocutorSearchStarted
	}

	_, err := bot.Send(msg)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}

	interlocutors := chats.GetUserInterlocutors(userID)
	if interlocutors != nil && len(interlocutors)+1 >= 2 {
		msg := tgbotapi.MessageConfig{
			Text:      betypes.GetTexts().Chat.ChatFound,
			ParseMode: betypes.GetTexts().ParseMode,
		}

		for i := 0; i < len(interlocutors); i++ {
			msg.ChatID = int64(interlocutors[i].ID)

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}
		}

		msg.ChatID = int64(userID)

		_, err := bot.Send(msg)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	}
}

// StopChatting if there are more than one people in the chat, removes the user from the chat.
// If less than two interlocutors delete the chat.
func StopChatting(userID int, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.MessageConfig{
		ParseMode: betypes.GetTexts().ParseMode,
		Text:      betypes.GetTexts().Chat.ChatEnded,
	}

	userInterlocutors := chats.GetUserInterlocutors(userID)
	if userInterlocutors == nil {
		if chats.IsUserInChat(userID) {
			chats.DeleteChatWithUser(userID)

			msg.ChatID = int64(userID)
			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}
		}

		return
	}

	if len(userInterlocutors) > 2 {
		for i := 0; i < len(userInterlocutors); i++ {
			msg.ParseMode = "MARKDOWN"
			msg.ChatID = int64(userInterlocutors[i].ID)
			msg.Text = fmt.Sprintf("*INTERLOCUTOR %d LEFT THE CHAT*", func() int {
				var userNumber int
				interlocutors := chats.GetUserInterlocutors(userInterlocutors[i].ID)
				for j := 0; j < len(interlocutors); j++ {
					if userID == interlocutors[j].ID {
						userNumber = j + 1
						break
					}
				}
				return userNumber
			}())

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}
		}
		chats.DeleteUserFromChat(userID)
	} else {
		for i := 0; i < len(userInterlocutors); i++ {
			msg.ChatID = int64(userInterlocutors[i].ID)
			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}
		}
		chats.DeleteChatWithUser(userID)
	}

	msg.ChatID = int64(userID)

	_, err := bot.Send(msg)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}
}

// Settings sends to user settings reply markup.
func Settings(userID int, bot *tgbotapi.BotAPI) {
	settingsInlineKeyboardMarkup := markups.GetSettings().FindInlineKeyboardMarkup(
		markups.SettingsReplyMarkupPrefix + markups.SettingsReplyMarkupName)
	if settingsInlineKeyboardMarkup == nil {
		return
	}

	msg := tgbotapi.MessageConfig{
		BaseChat:  tgbotapi.BaseChat{ChatID: int64(userID), ReplyMarkup: settingsInlineKeyboardMarkup},
		Text:      "⚙" + strings.Replace(markups.SettingsReplyMarkupName, markups.SettingsReplyMarkupPrefix, "", 1),
		ParseMode: betypes.GetTexts().ParseMode,
	}

	_, err := bot.Send(msg)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}
}

// Me sends information about itself to the user.
func Me(userID int, bot *tgbotapi.BotAPI) {
	u, err := database.DB.GetUser(userID)
	if err != nil && err.Error() == redis.Nil.Error() || u == nil {
		msg := tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(userID)},
			Text:      betypes.GetTexts().Chat.NotRegistered,
			ParseMode: betypes.GetTexts().ParseMode,
		}

		_, err := bot.Send(msg)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		return
	} else if err != nil && err.Error() != redis.Nil.Error() {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}

	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{ChatID: int64(userID)},
		Text: fmt.Sprintf(
			"📝*Information about you.*\n"+
				"📅*Your age:* %s.\n"+
				"🙍‍♂🙍‍♀📅*Interlocutor age:* %s.\n"+
				"🌍*Your city:* %s.\n"+
				"🙍‍♂🙍‍♀🌍*Interlocutor city:* %s.\n"+
				"🙍‍♂🙍‍♀*Your sex:* %s.\n"+
				"🙍‍♂🙍‍♀*Sex of the interlocutor:* %s.",
			u.Age, u.AgeOfTheInterlocutor, u.City, u.CityOfTheInterlocutor, u.Sex, u.SexOfTheInterlocutor,
		),
		ParseMode: "MARKDOWN",
	}

	_, err = bot.Send(msg)
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		panic(err)
	}
}
