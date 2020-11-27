package main

import (
	"fmt"
	"log"
	"net/http"
	"telegram-chat_bot/betypes"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/logger"
	"telegram-chat_bot/src/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	newBot, botErr = tgbotapi.NewBotAPI(betypes.GetBotConfig().Bot.Token)
	newChats       = betypes.NewChats(betypes.GetBotConfig().Chat.Queue,
		betypes.GetBotConfig().Chat.Users)
)

func main() {
	go func() {
		log.Fatalln(http.ListenAndServe(":"+betypes.GetBotConfig().Bot.Port, nil))
	}()

	logger.ForLog("Authorized on account.")
	if botErr != nil {
		logger.ForLog(fmt.Sprintf("Error %s, creating bot.", botErr.Error()))
		panic(botErr)
	}
	logger.ForLog("Bot have created successfully.")

	logger.ForLog("Setting up webhook.")
	_, err := newBot.SetWebhook(
		tgbotapi.NewWebhook(betypes.GetBotConfig().Bot.WebHook))
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s. Problem in setting Webhook.", err.Error()))
		panic(err)
	}

	// Check for updates.
	for update := range newBot.ListenForWebhook("/") {
		logger.ForLog(fmt.Sprintf("Update : %v", update))
		checkUpdate(update, newChats, newBot)
	}
}

func checkUpdate(update tgbotapi.Update, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		// Is this message to another user.
		if !update.Message.IsCommand() && chats.UserIsChatting(update.Message.From.ID) {
			interlocutors := chats.GetUserInterlocutors(update.Message.From.ID)
			// If the interlocutors are not found.
			if interlocutors == nil {
				_, err := bot.Send(tgbotapi.MessageConfig{
					BaseChat:  tgbotapi.BaseChat{ChatID: int64(update.Message.From.ID)},
					Text:      "*YOU ARE NOT IN THE CHAT YET...*",
					ParseMode: "MARKDOWN"})
				if err != nil {
					logger.ForLog(fmt.Sprintf("Error %s. Sending a message. User ID - %d.",
						err.Error(), update.Message.From.ID))
				}

				return
			}

			// Sending a message.
			for i, user := range interlocutors {
				var msg tgbotapi.Chattable
				if update.Message.Document != nil {
					msg = tgbotapi.NewDocumentShare(int64(user.ID), update.Message.Document.FileID)
				} else if update.Message.Photo != nil {
					msg = tgbotapi.NewPhotoShare(int64(user.ID), (*update.Message.Photo)[0].FileID)
				} else if update.Message.Video != nil {
					tgbotapi.NewVideoShare(int64(user.ID), update.Message.Video.FileID)
				} else {
					msg = tgbotapi.MessageConfig{
						BaseChat:  tgbotapi.BaseChat{ChatID: int64(user.ID)},
						Text:      fmt.Sprintf("*Interlocutor %d:* %s", i+1, update.Message.Text),
						ParseMode: "MARKDOWN"}
				}

				_, err := bot.Send(msg)
				if err != nil {
					logger.ForLog(fmt.Sprintf("Error %s. Sending a message. User ID - %d.",
						err.Error(), update.Message.From.ID))
				}
			}

			return
		}

		// Is the message a command.
		if update.Message.IsCommand() {
			checkCommands(*update.Message, chats, bot)
			return
		}
	}

	// Is the message a callback query.
	if update.CallbackQuery != nil {
		return
	}
}

func checkCommands(message tgbotapi.Message, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	switch message.Command() {
	case betypes.GetBotCommands().Start.Command:
		user := betypes.User{User: tgbotapi.User{
			ID:           message.From.ID,
			FirstName:    message.From.FirstName,
			LastName:     message.From.LastName,
			UserName:     message.From.UserName,
			LanguageCode: message.From.LanguageCode,
			IsBot:        message.From.IsBot,
		}, Age: betypes.UserNil, City: betypes.UserNil}
		commands.Start(user, bot)
	case betypes.GetBotCommands().Help.Command:
		commands.Help(message.From.ID, bot)
	case betypes.GetBotCommands().Chatting.Start:
		msg := tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(message.From.ID)},
			Text:      "*You are already in chat.*",
			ParseMode: "MARKDOWN"}
		if !chats.UserIsChatting(message.From.ID) {
			u, err := database.DBRedis.GetUser(int64(message.From.ID))
			if err != nil && err.Error() == "redis: nil" /* If no user is found */ {
				logger.ForLog(fmt.Sprintf("User not found. User ID - %d.",
					int64(message.From.ID)))

				_, err := bot.Send(tgbotapi.NewMessage(
					int64(message.From.ID), "You are not registered.\"/start\""))
				if err != nil {
					logger.ForLog(fmt.Sprintf(
						"Error %s. Sending a message. User ID - %d.",
						err.Error(), message.From.ID))
				}
			}

			if err != nil {
				logger.ForLog(fmt.Sprintf("Error, %s.", err.Error()))
				panic(err)
			}

			chats.AddUserToTheQueue(*u)

			msg.Text = "*The search for the interlocutor has begun....*"
		}

		_, err := bot.Send(msg)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s. Sending a message. User ID - %d.",
				err.Error(), message.From.ID))
		}

		chat := chats.GetChatByUserID(message.From.ID)
		if chat != nil {
			if int64(len(chat.Users)) == int64(chats.UsersCount) {
				msg = tgbotapi.MessageConfig{Text: "*CHAT FOUND*", ParseMode: "MARKDOWN"}
				for _, user := range chat.Users {
					msg.ChatID = int64(user.ID)
					_, err = bot.Send(msg)
					if err != nil {
						logger.ForLog(fmt.Sprintf(
							"Error %s. Sending a message. User ID - %d.",
							err.Error(), message.From.ID))
					}
				}
			}
		}
	case betypes.GetBotCommands().Chatting.Stop:
		chat := chats.GetChatByUserID(message.From.ID)
		if chat != nil {
			chats.StopChatting(message.From.ID)

			for _, user := range chat.Users {
				msg := tgbotapi.MessageConfig{
					BaseChat:  tgbotapi.BaseChat{ChatID: int64(user.ID)},
					Text:      "*CHAT ENDED*",
					ParseMode: "MARKDOWN",
				}

				_, err := bot.Send(msg)
				if err != nil {
					logger.ForLog(fmt.Sprintf(
						"Error %s, sending message. User ID - %d.",
						err.Error(), message.From.ID))
				}
			}
		}

		logger.ForLog(fmt.Sprintf("User ID - %d. Chat not found.",
			message.From.ID))
	default:
		_, err := bot.Send(tgbotapi.NewMessage(int64(message.From.ID),
			betypes.GetBotCommands().Unknown.Text))
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s, sending message. User ID - %d.",
				err.Error(), message.From.ID))
		}
	}
}
