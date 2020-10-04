package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"telegram-chat_bot/betypes"
	"telegram-chat_bot/loger"
	actions "telegram-chat_bot/src/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	go func() {
		log.Fatalln(http.ListenAndServe(":"+betypes.GetBotConfig().BotPort, nil))
	}()

	newBot, err := tgbotapi.NewBotAPI(betypes.GetBotConfig().BotToken)
	if err != nil {
		loger.ForLog("Error creating bot.", err)
		panic(err)
	}
	loger.ForLog(fmt.Sprintf("Authorized on account %s.", newBot.Self.FirstName))
	loger.ForLog("Bot have created successfully.")

	getUpdates(newBot)
}

func getUpdates(bot *tgbotapi.BotAPI) {
	setWebhook(bot)
	updates := bot.ListenForWebhook("/")

	for update := range updates {
		go check(&update, bot)
	}
}

func setWebhook(bot *tgbotapi.BotAPI) {
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(betypes.GetBotConfig().WebHook))
	if err != nil {
		loger.ForLog("Error, web hook.", err)
	}
}

func check(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.IsCommand() {
		checkCommand(update.Message, bot)
	}

	if update.CallbackQuery != nil {
		checkCallbackQuery(update.CallbackQuery, bot)
	}
}

func checkCommand(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	loger.ForLog(fmt.Sprintf("Command: \"%s\", form user ID, %v.",
		message.Command(), message.From.ID))
	switch message.Command() {
	case betypes.GetBotCommands().Start.Command:
		actions.StartCommand(&betypes.User{User: tgbotapi.User{
			ID:           message.From.ID,
			FirstName:    message.From.FirstName,
			LastName:     message.From.LastName,
			UserName:     message.From.UserName,
			LanguageCode: message.From.LanguageCode,
			IsBot:        message.From.IsBot,
		}}, bot)
	case betypes.GetBotCommands().Help.Command:
		actions.HelpCommand(int64(message.From.ID), bot)
	case betypes.GetBotCommands().StartChatting.Command:
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
