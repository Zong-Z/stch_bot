package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"telegram-chat_bot/betypes"
	"telegram-chat_bot/logger"
	commands "telegram-chat_bot/src"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	go func() {
		log.Fatalln(http.ListenAndServe(":"+betypes.GetBotConfig().Bot.Port, nil))
	}()

	bot, err := tgbotapi.NewBotAPI(betypes.GetBotConfig().Bot.Token)
	chats := betypes.NewChats(betypes.GetBotConfig().Chat.Queue, betypes.GetBotConfig().Chat.Users)

	logger.ForLog("Authorized on account.")
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s, creating bot.", err.Error()))
		panic(err)
	}
	logger.ForLog("Bot have created successfully.")

	var updates tgbotapi.UpdatesChannel
	if !strings.EqualFold(betypes.GetBotConfig().Bot.Webhook, "") {
		updates, err = setWebhook(bot)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		}
	}

	if updates == nil {
		updates, err = setPolling(bot)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	}

	for update := range updates {
		logger.ForLog(fmt.Sprintf("Update ID - %d.", update.UpdateID))

		checkUpdate(update, chats, bot)
	}
}

func setPolling(bot *tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}

	return updates, nil
}

func setWebhook(bot *tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	_, err := bot.SetWebhook(
		tgbotapi.NewWebhook(betypes.GetBotConfig().Bot.Webhook))
	if err != nil {
		return nil, err
	}
	updates := bot.ListenForWebhook("/")

	return updates, nil
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
					Text:      betypes.GetTexts().Chat.NotInChat,
					ParseMode: betypes.GetTexts().ParseMode,
				})
				if err != nil {
					logger.ForLog(fmt.Sprintf("Error %s. User ID - %d.", err.Error(), update.Message.From.ID))
					panic(err)
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
					var msgText string
					if betypes.GetBotConfig().Chat.Users > 2 /* If there are more than two interlocutors. */ {
						msgText = fmt.Sprintf("*INTERLOCUTOR %d:* %s", i+1, update.Message.Text)
					} else {
						msgText = update.Message.Text
					}
					msg = tgbotapi.MessageConfig{
						BaseChat: tgbotapi.BaseChat{ChatID: int64(user.ID)},
						Text:     msgText, ParseMode: betypes.GetTexts().ParseMode}
				}

				_, err := bot.Send(msg)
				if err != nil {
					logger.ForLog(fmt.Sprintf("Error %s. User ID - %d.", err.Error(), update.Message.From.ID))
					panic(err)
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
		if strings.Contains(update.CallbackQuery.Data, commands.SettingsPrefixReplyMarkup) {
			commands.AnswerOnCallbackQuerySettings(*update.CallbackQuery, bot)
		}

		return
	}
}

func checkCommands(message tgbotapi.Message, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	switch message.Command() {
	case betypes.GetTexts().Commands.Start.Command:
		user := betypes.User{User: tgbotapi.User{
			ID:           message.From.ID,
			FirstName:    message.From.FirstName,
			LastName:     message.From.LastName,
			UserName:     message.From.UserName,
			LanguageCode: message.From.LanguageCode,
			IsBot:        message.From.IsBot,
		}, Age: betypes.UserNil, City: betypes.UserNil}
		commands.Start(user, bot)
	case betypes.GetTexts().Commands.Help.Command:
		commands.Help(message.From.ID, bot)
	case betypes.GetTexts().Commands.Chatting.Start:
		commands.StartChat(message, chats, bot)
	case betypes.GetTexts().Commands.Chatting.Stop:
		commands.StopChat(message, chats, bot)
	case betypes.GetTexts().Commands.Settings.Command:
		commands.Settings(int64(message.From.ID), bot)
	default:
		msg := tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(message.From.ID)},
			Text:      betypes.GetTexts().Commands.Unknown.Text,
			ParseMode: betypes.GetTexts().ParseMode,
		}

		_, err := bot.Send(msg)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	}
}
