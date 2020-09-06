package actions

import (
	"telegram-chat_bot/betypes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HelpCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	bot.Send(tgbotapi.NewMessage(int64(update.Message.From.ID), betypes.GetHelpText()))
}
