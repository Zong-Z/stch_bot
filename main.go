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
	newBot, botErr = tgbotapi.NewBotAPI(betypes.GetBotConfig().BotToken)
	newChats       = betypes.NewChats(betypes.GetBotConfig().ChatsConfig.QueueSize,
		betypes.GetBotConfig().ChatsConfig.UsersCount)
)

func main() {
	go func() {
		log.Fatalln(http.ListenAndServe(":"+betypes.GetBotConfig().BotPort, nil))
	}()

	logger.ForLog("Authorized on account.")
	if botErr != nil {
		logger.ForLog(fmt.Sprintf("Error %s, creating bot.", botErr.Error()))
		panic(botErr)
	}
	logger.ForLog("Bot have created successfully.")

	logger.ForLog("Setting up webhook.")
	_, err := newBot.SetWebhook(tgbotapi.NewWebhook(betypes.GetBotConfig().WebHook))
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s. Problem in setting Webhook.", err.Error()))
		panic(err)
	}

	// Check for updates.
	for update := range newBot.ListenForWebhook("/") {
		if update.Message != nil {
			// Is this message to another user.
			if !update.Message.IsCommand() && newChats.UserIsChatting(update.Message.From.ID) {
				interlocutors := newChats.GetUserInterlocutors(update.Message.From.ID)
				// If the interlocutors are not found.
				if interlocutors == nil {
					continue
				}

				// Sending a message.
				for i, user := range interlocutors {
					_, err := newBot.Send(tgbotapi.MessageConfig{
						BaseChat:  tgbotapi.BaseChat{ChatID: int64(user.ID)},
						Text:      fmt.Sprintf("*Interlocutor %d:* %s", i+1, update.Message.Text),
						ParseMode: "MARKDOWN"})
					if err != nil {
						logger.ForLog(fmt.Sprintf("Error %s. Sending a message. User ID - %d.",
							err.Error(), update.Message.From.ID))
					}
				}

				continue
			}

			// Is the message a command.
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case betypes.GetBotCommands().Start.Command:
					user := betypes.User{User: tgbotapi.User{
						ID:           update.Message.From.ID,
						FirstName:    update.Message.From.FirstName,
						LastName:     update.Message.From.LastName,
						UserName:     update.Message.From.UserName,
						LanguageCode: update.Message.From.LanguageCode,
						IsBot:        update.Message.From.IsBot,
					}, Age: betypes.UserNil, City: betypes.UserNil}
					commands.Start(user, newBot)
					continue
				case betypes.GetBotCommands().Help.Command:
					commands.Help(update.Message.From.ID, newBot)
					continue
				case betypes.GetBotCommands().Chatting.Start:
					msg := tgbotapi.MessageConfig{
						BaseChat:  tgbotapi.BaseChat{ChatID: int64(update.Message.From.ID)},
						Text:      "*You are already in chat.*",
						ParseMode: "MARKDOWN"}
					if !newChats.UserIsChatting(update.Message.From.ID) {
						u, err := database.DBRedis.GetUser(int64(update.Message.From.ID))
						if err != nil && err.Error() == "redis: nil" /* If no user is found */ {
							logger.ForLog(fmt.Sprintf("User not found. User ID - %d.",
								int64(update.Message.From.ID)))

							_, err := newBot.Send(tgbotapi.NewMessage(
								int64(update.Message.From.ID), "You are not registered.\"/start\""))
							if err != nil {
								logger.ForLog(fmt.Sprintf(
									"Error %s. Sending a message. User ID - %d.",
									err.Error(), update.Message.From.ID))
							}
							continue
						}

						if err != nil {
							logger.ForLog(fmt.Sprintf("Error, %s.", err.Error()))
							panic(err)
						}

						newChats.AddUserToTheQueue(*u)

						msg.Text = "*The search for the interlocutor has begun....*"
					}

					_, err = newBot.Send(msg)
					if err != nil {
						logger.ForLog(fmt.Sprintf("Error %s. Sending a message. User ID - %d.",
							err.Error(), update.Message.From.ID))
					}

					chat := newChats.GetChatByUserID(update.Message.From.ID)
					if chat != nil {

						if int64(len(chat.Users)) == int64(newChats.UsersCount) {
							msg = tgbotapi.MessageConfig{
								BaseChat: tgbotapi.BaseChat{
									ChatID: int64(update.Message.From.ID),
								},
								Text:      "*CHAT FOUND*",
								ParseMode: "MARKDOWN",
							}
							_, err = newBot.Send(msg)
							if err != nil {
								logger.ForLog(fmt.Sprintf(
									"Error %s. Sending a message. User ID - %d.",
									err.Error(), update.Message.From.ID))
							}
						}
					}

					continue
				case betypes.GetBotCommands().Chatting.Stop:
					chat := newChats.GetChatByUserID(update.Message.From.ID)
					if chat != nil {
						newChats.StopChatting(update.Message.From.ID)

						for _, user := range chat.Users {
							msg := tgbotapi.MessageConfig{
								BaseChat:  tgbotapi.BaseChat{ChatID: int64(user.ID)},
								Text:      "*CHAT ENDED*",
								ParseMode: "MARKDOWN",
							}

							_, err := newBot.Send(msg)
							if err != nil {
								logger.ForLog(fmt.Sprintf(
									"Error %s, sending message. User ID - %d.",
									err.Error(), update.Message.From.ID))
							}
						}
					}

					logger.ForLog(fmt.Sprintf("User ID - %d. Chat not found.",
						update.Message.From.ID))

					continue
				}
				continue
			}
		}

		// Is the message a callback query.
		if update.CallbackQuery != nil {
			continue
		}
	}
}
