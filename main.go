package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"telegram-chat_bot/betypes"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/loger"
	actions "telegram-chat_bot/src/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	go func() {
		log.Fatalln(http.ListenAndServe(":"+betypes.GetBotConfig().BotPort, nil))
	}()

	bot, err := tgbotapi.NewBotAPI(betypes.GetBotConfig().BotToken)
	if err != nil {
		loger.ForLog(fmt.Sprintf("Error %v, creating bot.", err))
		panic(err)
	}
	loger.ForLog(fmt.Sprintf("Authorized on account %s.", bot.Self.FirstName))
	loger.ForLog("Bot have created successfully.")

	getUpdates(bot, betypes.NewChats(2 /*Chat for two users*/, betypes.GetBotConfig().ChatsConfig.QueueSize, bot))
}

func getUpdates(bot *tgbotapi.BotAPI, chats *betypes.Chats) {
	setWebhook(bot)
	updates := bot.ListenForWebhook("/")

	for update := range updates {
		go checkUpdate(&update, chats, bot)
	}
}

func setWebhook(bot *tgbotapi.BotAPI) {
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(betypes.GetBotConfig().WebHook))
	if err != nil {
		loger.ForLog("Error, web hook.", err)
	}
}

func checkUpdate(update *tgbotapi.Update, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.IsCommand() {
		checkCommand(update.Message, chats, bot)
		return
	}

	if update.CallbackQuery != nil {
		checkCallbackQuery(update.CallbackQuery, bot)
		return
	}

	if update.Message != nil && !update.Message.IsCommand() {
		if chats.IsTheUserInChat(update.Message.From.ID) {
			chats.SendMessageToInterlocutors(update.Message.Text, update.Message.From.ID, bot)
			return
		}
		if _, err := bot.Send(tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: int64(update.Message.From.ID),
			},
			Text:      "*Unknown message* \"/help\".",
			ParseMode: "MARKDOWN",
		}); err != nil {
			loger.ForLog(fmt.Sprintf("Error %s, sending message. User ID - %d.", err.Error(), update.Message.From.ID))
		}
		return
	}
}

func checkCommand(message *tgbotapi.Message, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	loger.ForLog(fmt.Sprintf("Command: \"%s\", form user ID, %v.",
		message.Command(), message.From.ID))
	switch message.Command() {
	case betypes.GetBotCommands().Start.Command:
		actions.StartCommand(&betypes.User{
			User: tgbotapi.User{
				ID:           message.From.ID,
				FirstName:    message.From.FirstName,
				LastName:     message.From.LastName,
				UserName:     message.From.UserName,
				LanguageCode: message.From.LanguageCode,
				IsBot:        message.From.IsBot,
			},
			Age:  betypes.UserNil,
			City: betypes.UserNil,
		}, bot)
	case betypes.GetBotCommands().Help.Command:
		actions.HelpCommand(int64(message.From.ID), bot)
	case betypes.GetBotCommands().Chatting.Start:
		u, err := database.GetUser(int64(message.From.ID))
		if err != nil && err.Error() == "redis: nil" /*If no user is found*/ {
			loger.ForLog(fmt.Sprintf("User not found, user ID, %d.", int64(message.From.ID)))
			if _, err := bot.Send(tgbotapi.MessageConfig{
				BaseChat: tgbotapi.BaseChat{
					ChatID: int64(message.From.ID),
				},
				Text:      "*You are not registered.*\"/start\"",
				ParseMode: "MARKDOWN",
			}); err != nil {
				loger.ForLog(fmt.Sprintf("Error %s, sending message. User ID - %d.", err.Error(), message.From.ID))
			}
			return
		}

		if err != nil {
			loger.ForLog(fmt.Sprintf("Error, %s.", err.Error()))
			panic(err)
		}

		chats.StartChatting(u, bot)
	case betypes.GetBotCommands().Chatting.Stop:
		chats.StopChatting(message.From.ID, bot)
	case betypes.GetBotCommands().Settings.Command:
		actions.SettingsCommandMarkup(int64(message.From.ID), bot)
	}
}

func checkCallbackQuery(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	loger.ForLog(fmt.Sprintf("CallbackQuery: \"%v\", form user ID, %v.", *callbackQuery, callbackQuery.From.ID))
	if strings.EqualFold(callbackQuery.Data, "close") {
		actions.DeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, bot)
	}
}
