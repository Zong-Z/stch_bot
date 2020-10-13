package actions

import (
	"fmt"
	"telegram-chat_bot/betypes"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/loger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartCommand display the start text save the user in the database.
func StartCommand(user *betypes.User, bot *tgbotapi.BotAPI) {
	if _, err := database.GetUser(int64(user.ID)); err != nil && err.Error() == "redis: nil" /*If no user is found*/ {
		if err := database.SaveUser(*user); err != nil {
			loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %d", err, user.ID))
		}
	}

	if _, err := bot.Send(tgbotapi.NewMessage(int64(user.ID), betypes.GetBotCommands().Start.Text)); err != nil {
		loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %d", err, user.ID))
	}
}
