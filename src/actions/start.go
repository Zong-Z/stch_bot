package actions

import (
	"fmt"
	"telegram-chat_bot/betypes"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/loger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func StartCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	_, err := database.GetUser(update.Message.From.ID)
	if err != nil && err.Error() == "redis: nil" /*If no user is found*/ {
		if err := database.SaveUser(
			betypes.User{
				User: tgbotapi.User{
					ID:           update.Message.From.ID,
					FirstName:    update.Message.From.FirstName,
					LastName:     update.Message.From.LastName,
					UserName:     update.Message.From.UserName,
					LanguageCode: update.Message.From.LanguageCode,
					IsBot:        update.Message.From.IsBot,
				},
				Age:  betypes.UserNull,
				City: betypes.UserNull,
			}); err != nil {
			loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %v",
				err, update.Message.From.ID))
		}
	}
	if _, err := bot.Send(tgbotapi.NewMessage(
		int64(update.Message.From.ID), betypes.GetBotCommands().Start.Text)); err != nil {
		loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %v",
			err, update.Message.From.ID))
	}
}
