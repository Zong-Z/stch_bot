package main

import (
	"fmt"
	"log"
	"net/http"
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

func checkOnCommands(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		if update.Message.IsCommand() {
			loger.ForLog(fmt.Sprintf("Command: \"%s\", form user ID, %v",
				update.Message.Command(), update.Message.From.ID))
			switch update.Message.Command() {
			case betypes.GetBotCommands().Start.Command:
				actions.StartCommand(&betypes.User{
					User: tgbotapi.User{
						ID:           update.Message.From.ID,
						FirstName:    update.Message.From.FirstName,
						LastName:     update.Message.From.LastName,
						UserName:     update.Message.From.UserName,
						LanguageCode: update.Message.From.LanguageCode,
						IsBot:        update.Message.From.IsBot,
					},
				}, bot)
			case betypes.GetBotCommands().Help.Command:
				actions.HelpCommand(int64(update.Message.From.ID), bot)
			case betypes.GetBotCommands().Settings.Command:
				actions.SettingsCommand(int64(update.Message.From.ID), bot)
			}
			return
		}
		return
	}

	if update.CallbackQuery != nil {
		loger.ForLog(fmt.Sprintf("CallbackQuery: \"%v\", form user ID, %v", update.CallbackQuery, update.CallbackQuery.From.ID))
		return
	}
}

func getUpdates(bot *tgbotapi.BotAPI) {
	setWebhook(bot)
	updates := bot.ListenForWebhook("/")

	for update := range updates {
		go checkOnCommands(&update, bot)
	}
}

func setWebhook(bot *tgbotapi.BotAPI) {
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(betypes.GetBotConfig().WebHook))
	if err != nil {
		loger.ForLog("Error, web hook.", err)
	}
}
