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
		loger.LogFile.Fatalln("Error creating bot.", botError)
	}
	loger.LogFile.Println("Bot have created successfully.")
	loger.LogFile.Println(fmt.Sprintf("Authorized on account %s.", newBot.Self.FirstName))

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
		loger.LogFile.Fatalln("Error, web hook.", err)
	}
}
