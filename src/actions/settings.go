package actions

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var (
	mainSettingsMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Change City", "Change City"),
			tgbotapi.NewInlineKeyboardButtonData("Your Age", "Your Age"),
			tgbotapi.NewInlineKeyboardButtonData("Age Of The Interlocutor", "Age Of The Interlocutor")),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Close Settings", "Close Settings")),
	)
)

func SettingsCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	bot.Send(tgbotapi.NewMessage(int64(update.Message.From.ID), "Settings"))
}
