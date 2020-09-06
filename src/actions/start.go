package actions

import (
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
			}); err != nil {
			loger.LogFile.Println("Error, failed to save user to database.")
		}
	}
	bot.Send(tgbotapi.NewMessage(int64(update.Message.From.ID), betypes.GetStartText()))
}
