package actions

import (
	"fmt"
	"telegram-chat_bot/betypes"
	"telegram-chat_bot/loger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// HelpCommand sends help text to user.
func HelpCommand(chatID int64, bot *tgbotapi.BotAPI) {
	if _, err := bot.Send(tgbotapi.NewMessage(chatID, betypes.GetBotCommands().Help.Text)); err != nil {
		loger.ForLog(fmt.Sprintf("Error %v, sending message. Chat ID, %v", err, chatID))
	}
}
