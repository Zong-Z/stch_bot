package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"telegram-chat_bot/betypes"
	"telegram-chat_bot/commands"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/logger"
	"telegram-chat_bot/markups"

	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	go func() {
		log.Fatalln(http.ListenAndServe(":"+betypes.GetConfig().Bot.Port, nil))
	}()

	bot, err := tgbotapi.NewBotAPI(betypes.GetConfig().Bot.Token)
	chats := betypes.NewChats(betypes.GetConfig().Chat.Queue, betypes.GetConfig().Chat.Users)

	logger.ForLog("Authorized on account.")
	if err != nil {
		logger.ForLog(fmt.Sprintf("Error %s, creating bot.", err.Error()))
		panic(err)
	}

	logger.ForLog("Bot have created successfully.")

	var updates tgbotapi.UpdatesChannel
	if !strings.EqualFold(betypes.GetConfig().Bot.Webhook, "") {
		updates, err = setWebhook(bot)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
		}
	}

	if updates == nil {
		updates, err = setPolling(betypes.GetConfig().Bot.Polling.Offset, betypes.GetConfig().Bot.Polling.Limit,
			betypes.GetConfig().Bot.Polling.Timeout, bot)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	}

	for update := range updates {
		logger.ForLog(fmt.Sprintf("Update ID - %d.", update.UpdateID))
		checkUpdate(update, chats, bot)
	}
}

func checkUpdate(update tgbotapi.Update, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery != nil {
		checkCallbackQuery(*update.CallbackQuery, bot)
	} else if update.Message != nil {
		isCommand := update.Message.IsCommand()
		isUserInChat := chats.IsUserInChat(update.Message.From.ID)
		if isCommand {
			checkCommands(*update.Message, chats, bot)
		} else if !isCommand && !isUserInChat {
			msg := tgbotapi.MessageConfig{
				BaseChat:  tgbotapi.BaseChat{ChatID: int64(update.Message.From.ID)},
				Text:      betypes.GetTexts().Commands.Unknown.Text,
				ParseMode: betypes.GetTexts().ParseMode,
			}

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}
		} else if !isCommand && isUserInChat {
			sendMessageToUserInterlocutors(*update.Message, chats, bot)
		}
	}
}

func checkCommands(message tgbotapi.Message, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	switch message.Command() {
	case betypes.GetTexts().Commands.Start.Command:
		commands.Start(*betypes.NewUser(tgbotapi.User{
			ID:           message.From.ID,
			FirstName:    message.From.FirstName,
			LastName:     message.From.LastName,
			UserName:     message.From.UserName,
			LanguageCode: message.From.LanguageCode,
			IsBot:        message.From.IsBot}), bot)
	case betypes.GetTexts().Commands.Help.Command:
		commands.Help(message.From.ID, bot)
	case betypes.GetTexts().Commands.Chatting.Start:
		commands.StartChatting(message.From.ID, chats, bot)
	case betypes.GetTexts().Commands.Chatting.Stop:
		commands.StopChatting(message.From.ID, chats, bot)
	case betypes.GetTexts().Commands.Settings.Command:
		commands.Settings(message.From.ID, bot)
	case betypes.GetTexts().Commands.Me.Command:
		commands.Me(message.From.ID, bot)
	default:
		msg := tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(message.From.ID)},
			Text:      betypes.GetTexts().Commands.Unknown.Text,
			ParseMode: betypes.GetTexts().ParseMode,
		}

		_, err := bot.Send(msg)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	}
}

func checkCallbackQuery(callbackQuery tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	if markups.IsThereCloseCallback(callbackQuery.Data) {
		_, err := bot.Send(tgbotapi.NewDeleteMessage(int64(callbackQuery.From.ID), callbackQuery.Message.MessageID))
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	} else if markups.SettingsIsThereMarkupRequest(callbackQuery.Data) {
		u, err := database.DB.GetUser(callbackQuery.From.ID)
		if err != nil && err.Error() == redis.Nil.Error() || u == nil {
			msg := tgbotapi.MessageConfig{
				BaseChat:  tgbotapi.BaseChat{ChatID: int64(callbackQuery.From.ID)},
				Text:      betypes.GetTexts().Chat.NotRegistered,
				ParseMode: betypes.GetTexts().ParseMode,
			}

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}
		} else if err != nil && err.Error() != redis.Nil.Error() {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		inlineKeyboardMarkup := markups.GetSettings().FindInlineKeyboardMarkup(callbackQuery.Data)
		if inlineKeyboardMarkup == nil {
			return
		}

		for i := 0; i < len(inlineKeyboardMarkup.InlineKeyboard); i++ {
			for j := 0; j < len(inlineKeyboardMarkup.InlineKeyboard[i]); j++ {
				callbackData := strings.Replace(*inlineKeyboardMarkup.InlineKeyboard[i][j].CallbackData,
					markups.SettingsReplyMarkupPrefix, "", 1)
				if strings.EqualFold(callbackData, markups.OwnAgePrefix+u.Age) ||
					strings.EqualFold(callbackData, markups.AgeOfTheInterlocutorPrefix+u.AgeOfTheInterlocutor) ||
					strings.EqualFold(callbackData, markups.OwnCityPrefix+u.City) ||
					strings.EqualFold(callbackData, markups.CityOfTheInterlocutorPrefix+u.CityOfTheInterlocutor) ||
					strings.EqualFold(callbackData, markups.OwnSexPrefix+u.Sex) ||
					strings.EqualFold(callbackData, markups.SexOfTheInterlocutorPrefix+u.SexOfTheInterlocutor) {
					inlineKeyboardMarkup.InlineKeyboard[i][j].Text = fmt.Sprintf("➡%s⬅",
						inlineKeyboardMarkup.InlineKeyboard[i][j].Text)
				}
			}
		}

		editMessageReplyMarkup := tgbotapi.NewEditMessageReplyMarkup(int64(callbackQuery.From.ID),
			callbackQuery.Message.MessageID, *inlineKeyboardMarkup)

		_, err = bot.Send(editMessageReplyMarkup)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		_, err = bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuery.ID,
			betypes.GetTexts().ReplyKeyboardMarkup.Opened))
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	} else if markups.SettingsIsThereCallbackForChange(callbackQuery.Data) {
		settingsCallbackForChangeUserData(callbackQuery, bot)
	}
}

func sendMessageToUserInterlocutors(message tgbotapi.Message, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	var msg tgbotapi.Chattable
	userInterlocutors := chats.GetUserInterlocutors(message.From.ID)
	if userInterlocutors == nil {
		msg = tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(message.From.ID)},
			Text:      betypes.GetTexts().Chat.NotInChat,
			ParseMode: betypes.GetTexts().ParseMode,
		}

		_, err := bot.Send(msg)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}
	} else {
		for i := 0; i < len(userInterlocutors); i++ {
			if message.Audio != nil {
				msg = tgbotapi.NewAudioShare(
					int64(userInterlocutors[i].ID), message.Audio.FileID)
			} else if message.Document != nil {
				msg = tgbotapi.NewDocumentShare(
					int64(userInterlocutors[i].ID), message.Document.FileID)
			} else if message.Animation != nil {
				msg = tgbotapi.NewAnimationShare(
					int64(userInterlocutors[i].ID), message.Animation.FileID)
			} else if message.Photo != nil {
				msg = tgbotapi.NewPhotoShare(
					int64(userInterlocutors[i].ID), (*message.Photo)[0].FileID)
			} else if message.Sticker != nil {
				msg = tgbotapi.NewStickerShare(
					int64(userInterlocutors[i].ID), message.Sticker.FileID)
			} else if message.Video != nil {
				msg = tgbotapi.NewVideoShare(
					int64(userInterlocutors[i].ID), message.Video.FileID)
			} else if message.VideoNote != nil {
				msg = tgbotapi.NewVideoNoteShare(int64(userInterlocutors[i].ID),
					message.VideoNote.Length, message.VideoNote.FileID)
			} else if message.Voice != nil {
				msg = tgbotapi.NewVoiceShare(int64(userInterlocutors[i].ID), message.Voice.FileID)
			} else {
				if betypes.GetConfig().Chat.Users > 2 { // If there are more than two interlocutors.
					msg = tgbotapi.MessageConfig{
						BaseChat:  tgbotapi.BaseChat{ChatID: int64(userInterlocutors[i].ID)},
						Text:      fmt.Sprintf("*INTERLOCUTOR %d:* %s", i+1, message.Text),
						ParseMode: betypes.GetTexts().ParseMode,
					}
				} else {
					msg = tgbotapi.MessageConfig{
						BaseChat: tgbotapi.BaseChat{ChatID: int64(userInterlocutors[i].ID)},
						Text:     message.Text,
					}
				}
			}

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}
		}
	}
}

func settingsCallbackForChangeUserData(callbackQuery tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	if func() bool {
		u, err := database.DB.GetUser(callbackQuery.From.ID)
		if err != nil && err.Error() == redis.Nil.Error() {
			msg := tgbotapi.MessageConfig{
				BaseChat:  tgbotapi.BaseChat{ChatID: int64(callbackQuery.From.ID)},
				Text:      betypes.GetTexts().Chat.NotRegistered,
				ParseMode: betypes.GetTexts().ParseMode,
			}

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
				panic(err)
			}

			return false
		}

		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		callbackQueryData := strings.Replace(callbackQuery.Data, markups.SettingsReplyMarkupPrefix, "", 1)
		if strings.Contains(callbackQueryData, markups.OwnAgePrefix) {
			callbackQueryData = strings.Replace(callbackQueryData, markups.OwnAgePrefix, "", 1)
			if u.Age == callbackQueryData {
				return true
			}

			u.Age = callbackQueryData
		} else if strings.Contains(callbackQueryData, markups.AgeOfTheInterlocutorPrefix) {
			callbackQueryData = strings.Replace(callbackQueryData, markups.AgeOfTheInterlocutorPrefix, "", 1)
			if u.AgeOfTheInterlocutor == callbackQueryData {
				return true
			}

			u.AgeOfTheInterlocutor = callbackQueryData
		} else if strings.Contains(callbackQueryData, markups.OwnCityPrefix) {
			callbackQueryData = strings.Replace(callbackQueryData, markups.OwnCityPrefix, "", 1)
			if u.City == callbackQueryData {
				return true
			}

			u.City = callbackQueryData
		} else if strings.Contains(callbackQueryData, markups.CityOfTheInterlocutorPrefix) {
			callbackQueryData = strings.Replace(callbackQueryData, markups.CityOfTheInterlocutorPrefix, "", 1)
			if u.CityOfTheInterlocutor == callbackQueryData {
				return true
			}

			u.CityOfTheInterlocutor = callbackQueryData
		} else if strings.Contains(callbackQueryData, markups.OwnSexPrefix) {
			callbackQueryData = strings.Replace(callbackQueryData, markups.OwnSexPrefix, "", 1)
			if u.Sex == callbackQueryData {
				return true
			}

			u.Sex = callbackQueryData
		} else if strings.Contains(callbackQueryData, markups.SexOfTheInterlocutorPrefix) {
			callbackQueryData = strings.Replace(callbackQueryData, markups.SexOfTheInterlocutorPrefix, "", 1)
			if u.SexOfTheInterlocutor == callbackQueryData {
				return true
			}

			u.SexOfTheInterlocutor = callbackQueryData
		} else {
			return false
		}

		err = database.DB.SaveUser(*u)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		return true
	}() {
		_, err := bot.AnswerCallbackQuery(
			tgbotapi.NewCallback(callbackQuery.ID, betypes.GetTexts().ReplyKeyboardMarkup.Changed))
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		settingsInlineKeyboard := markups.GetSettings().FindInlineKeyboardMarkup(
			markups.SettingsReplyMarkupPrefix + markups.SettingsReplyMarkupName)
		if settingsInlineKeyboard == nil {
			return
		}

		editMessageReplyMarkup := tgbotapi.NewEditMessageReplyMarkup(int64(callbackQuery.From.ID),
			callbackQuery.Message.MessageID, *settingsInlineKeyboard)
		_, err = bot.Send(editMessageReplyMarkup)
		if err != nil {
			logger.ForLog(fmt.Sprintf("Error %s.", err.Error()))
			panic(err)
		}

		return
	}
}

func setPolling(offset, limit, timeout int, bot *tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.UpdateConfig{
		Offset:  offset,
		Limit:   limit,
		Timeout: timeout,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}

	return updates, nil
}

func setWebhook(bot *tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(betypes.GetConfig().Bot.Webhook))
	if err != nil {
		return nil, err
	}

	updates := bot.ListenForWebhook("/")

	return updates, nil
}
