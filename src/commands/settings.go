package commands

// import (
// 	"fmt"
// 	"strings"
// 	"telegram-chat_bot/betypes"
// 	database "telegram-chat_bot/db"
// 	"telegram-chat_bot/loger"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// )

// const agePrefix = "user age:"

// // Markup stores the markup and keyboard name.
// type Markup struct {
// 	Name   string                        `json:"name"`
// 	Markup tgbotapi.InlineKeyboardMarkup `json:"keyboard"`
// }

// var markups = []Markup{
// 	{
// 		Name: "settings",
// 		Markup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
// 			tgbotapi.NewInlineKeyboardRow(
// 				tgbotapi.NewInlineKeyboardButtonData("Age.", "age"),
// 				/*//TODO:*/ tgbotapi.NewInlineKeyboardButtonData("City.", "city"),
// 			),
// 			tgbotapi.NewInlineKeyboardRow(
// 				tgbotapi.NewInlineKeyboardButtonData("Close.", "close"),
// 			),
// 		}},
// 	},
// 	{
// 		Name: "age",
// 		Markup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
// 			tgbotapi.NewInlineKeyboardRow(
// 				tgbotapi.NewInlineKeyboardButtonData("Sixteen and below.", betypes.UserAgeSixteenAndBelow),
// 				tgbotapi.NewInlineKeyboardButtonData("Eighteen or more.", betypes.UserAgeEighteenOrMore),
// 			),
// 			tgbotapi.NewInlineKeyboardRow(
// 				tgbotapi.NewInlineKeyboardButtonData("<-Go back.", "settings"),
// 			),
// 		}},
// 	},
// }

// // SettingsCommandMarkup sends the user a keyboard with settings.
// func SettingsCommandMarkup(chatID int64, bot *tgbotapi.BotAPI) {
// 	markup := findMarkup("settings")
// 	if markup != nil {
// 		if _, err := bot.Send(tgbotapi.MessageConfig{
// 			BaseChat: tgbotapi.BaseChat{
// 				ChatID:      chatID,
// 				ReplyMarkup: (*markup).Markup,
// 			},
// 			ParseMode: "MARKDOWN",
// 			Text:      fmt.Sprintf("*%s*", strings.ToTitle(markup.Name)),
// 		}); err != nil {
// 			loger.ForLog(fmt.Sprintf("Error %s, sending message. Chat ID - %d", err.Error(), chatID))
// 		}
// 	}
// }

// // AnswerOnCallback
// func AnswerOnCallbackSettings(chatID int64, messageID int,
// 	callbackQueryData, callbackQueryID string, bot *tgbotapi.BotAPI) {
// 	if strings.EqualFold(callbackQueryData, "close") {
// 		deleteMessage(chatID, messageID, bot)

// 		newCallback := tgbotapi.NewCallback(callbackQueryID, "Closed")
// 		if _, err := bot.AnswerCallbackQuery(newCallback); err != nil {
// 			loger.ForLog(fmt.Sprintf("Error %s, answer CallbackQuery. Chat ID - %d",
// 				err.Error(), chatID))
// 		}

// 		return
// 	}

// 	markup := findMarkup(callbackQueryData)
// 	if markup != nil {
// 		if _, err := bot.Send(tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup.Markup)); err != nil {
// 			loger.ForLog(fmt.Sprintf("Error %s, sending message. Chat ID - %d", err.Error(), chatID))
// 		}

// 		return
// 	}

// 	u, err := database.GetUser(chatID)
// 	if err != nil {
// 		loger.ForLog("Error, could not read user from DB. Chat ID - ", chatID)
// 		return
// 	}

// 	u.Age = callbackQueryData
// 	err = database.SaveUser(*u)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func deleteMessage(chatID int64, messageID int, bot *tgbotapi.BotAPI) {
// 	if _, err := bot.Send(tgbotapi.NewDeleteMessage(chatID, messageID)); err != nil {
// 		loger.ForLog(fmt.Sprintf("Error %s, sending message. Chat ID - %d", err.Error(), chatID))
// 	}
// }

// func findMarkup(markupName string) *Markup {
// 	for i := 0; i < len(markups); i++ {
// 		if strings.EqualFold(markups[i].Name, markupName) {
// 			return &markups[i]
// 		}
// 	}

// 	return nil
// }
