package commands

import (
	"fmt"
	"telegram-chat_bot/betypes"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Help sends help text to user.
func Help(chatID int, bot *tgbotapi.BotAPI) {
	_, err := bot.Send(tgbotapi.NewMessage(int64(chatID), betypes.GetBotCommands().Help.Text))
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s. Sending a message. User ID - %d.",
			err.Error(), chatID))
	}
}

// Start display the start text save the user in the database.
func Start(user betypes.User, bot *tgbotapi.BotAPI) {
	// Saving user.
	_, err := database.DBRedis.GetUser(int64(user.ID))
	if err != nil && err.Error() == "redis: nil" /* If no user is found */ {
		err := database.DBRedis.SaveUser(user)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s. Sending a message. User ID - %d.",
				err.Error(), user.ID))
		}
	}

	// Sending to user start text.
	_, err = bot.Send(tgbotapi.NewMessage(int64(user.ID), betypes.GetBotCommands().Start.Text))
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s. Sending a message. User ID - %d.",
			err.Error(), user.ID))
	}
}
