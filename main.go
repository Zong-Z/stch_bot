package main

import (
	"fmt"
	"log"
	"net/http"
	"telegram-chat_bot/betypes"
	"telegram-chat_bot/loger"
	"telegram-chat_bot/src/actions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	go func() {
		log.Fatalln(http.ListenAndServe(":"+betypes.GetBotConfig().BotPort, nil))
	}()

	newBot, botError := tgbotapi.NewBotAPI(betypes.GetBotConfig().BotToken)
	if botError != nil {
		loger.ForLog("Error creating bot.", botError)
	}
	loger.ForLog("Bot have created successfully.")
	loger.ForLog(fmt.Sprintf("Authorized on account %s.", newBot.Self.FirstName))

	getUpdates(newBot)
}

func checkOnCommands(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case betypes.GetBotCommands().Start.Command:
				actions.StartCommand(update, bot)
			case betypes.GetBotCommands().Help.Command:
				actions.HelpCommand(update, bot)
			case betypes.GetBotCommands().Settings.Command:
				actions.SettingsCommand(update, bot)
			}
			return
		}
		return
	}

	if update.CallbackQuery != nil {
		actions.SettingsCommand(update, bot)
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
