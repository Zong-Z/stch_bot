package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"telegram-chat_bot/betypes"
	"telegram-chat_bot/commands"
	"telegram-chat_bot/logger"
	"telegram-chat_bot/markups"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	go func() {
		log.Fatalln(http.ListenAndServe(":"+betypes.GetConfig().Bot.Port, nil))
	}()

	bot, err := tgbotapi.NewBotAPI(betypes.GetConfig().Bot.Token)
	chats := betypes.NewChats(betypes.GetConfig().Chat.Queue, betypes.GetConfig().Chat.Users)

	logger.ForInfo("Authorized on account")
	if err != nil {
		logger.ForError(err.Error())
	}

	logger.ForInfo("Bot have created successfully")

	var updates tgbotapi.UpdatesChannel
	if !strings.EqualFold(betypes.GetConfig().Bot.Webhook, "") {
		updates, err = setWebhook(bot)
		if err != nil {
			logger.ForWarning(err.Error())
		}
	}

	if updates == nil {
		updates, err = setPolling(betypes.GetConfig().Bot.Polling.Offset, betypes.GetConfig().Bot.Polling.Limit,
			betypes.GetConfig().Bot.Polling.Timeout, bot)
		if err != nil {
			logger.ForError(err.Error())
		}
	}

	for update := range updates {
		logger.ForInfo(fmt.Sprintf("Update ID - %d", update.UpdateID))
		go checkUpdate(update, chats, bot)
	}
}

func checkUpdate(update tgbotapi.Update, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
	switch {
	case update.CallbackQuery != nil:
		checkCallbackQuery(*update.CallbackQuery, bot)
		break
	case update.Message != nil:
		isCommand := update.Message.IsCommand()
		isUserInChat := chats.IsUserInChat(update.Message.From.ID)
		switch {
		case isCommand:
			checkCommand(*update.Message, chats, bot)
			break
		case !isCommand && isUserInChat: // Message to another user.
			var msg tgbotapi.Chattable
			interlocutors := chats.GetInterlocutorsByUserID(update.Message.From.ID)
			if interlocutors == nil {
				msg = tgbotapi.MessageConfig{
					BaseChat:  tgbotapi.BaseChat{ChatID: int64(update.Message.From.ID)},
					Text:      betypes.GetTexts().Chat.NotInChat,
					ParseMode: betypes.GetTexts().ParseMode,
				}

				_, err := bot.Send(msg)
				if err != nil {
					logger.ForError(err.Error())
				}

				return
			}

			for i, interlocutor := range interlocutors {
				switch {
				case update.Message.Audio != nil:
					msg = tgbotapi.NewAudioShare(int64(interlocutor.ID), update.Message.Audio.FileID)
					break
				case update.Message.Document != nil:
					msg = tgbotapi.NewDocumentShare(int64(interlocutor.ID), update.Message.Document.FileID)
					break
				case update.Message.Animation != nil:
					msg = tgbotapi.NewAnimationShare(int64(interlocutor.ID), update.Message.Animation.FileID)
					break
				case update.Message.Photo != nil:
					msg = tgbotapi.NewPhotoShare(int64(interlocutor.ID), (*update.Message.Photo)[0].FileID)
					break
				case update.Message.Sticker != nil:
					msg = tgbotapi.NewStickerShare(int64(interlocutor.ID), update.Message.Sticker.FileID)
					break
				case update.Message.Video != nil:
					msg = tgbotapi.NewVideoShare(int64(interlocutor.ID), update.Message.Video.FileID)
					break
				case update.Message.VideoNote != nil:
					msg = tgbotapi.NewVideoNoteShare(int64(interlocutor.ID), update.Message.VideoNote.Length,
						update.Message.VideoNote.FileID)
					break
				case update.Message.Voice != nil:
					msg = tgbotapi.NewVoiceShare(int64(interlocutor.ID), update.Message.Voice.FileID)
					break
				default:
					if betypes.GetConfig().Chat.Users > 2 { // If there are more than two users in the chat.
						msg = tgbotapi.MessageConfig{
							BaseChat:  tgbotapi.BaseChat{ChatID: int64(interlocutor.ID)},
							Text:      fmt.Sprintf("*INTERLOCUTOR %d:* %s", i+1, update.Message.Text),
							ParseMode: "MARKDOWN",
						}

						break
					}

					msg = tgbotapi.MessageConfig{
						BaseChat: tgbotapi.BaseChat{ChatID: int64(interlocutor.ID)},
						Text:     update.Message.Text,
					}
				}

				_, err := bot.Send(msg)
				if err != nil {
					logger.ForError(err.Error())
				}
			}

			break
		default:
			msg := tgbotapi.MessageConfig{
				BaseChat:  tgbotapi.BaseChat{ChatID: int64(update.Message.From.ID)},
				Text:      betypes.GetTexts().Commands.Unknown.Text,
				ParseMode: betypes.GetTexts().ParseMode,
			}

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForError(err.Error())
			}
		}
	}
}

func checkCommand(message tgbotapi.Message, chats *betypes.Chats, bot *tgbotapi.BotAPI) {
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
			logger.ForError(err.Error())
		}
	}
}

func checkCallbackQuery(callbackQuery tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	settings := markups.Settings{}
	switch {
	case markups.IsThereCloseCallback(callbackQuery.Data):
		_, err := bot.Send(tgbotapi.NewDeleteMessage(int64(callbackQuery.From.ID), callbackQuery.Message.MessageID))
		if err != nil {
			logger.ForError(err.Error())
		}
	case settings.IsMarkupRequest(callbackQuery.Data):
		settings.SendMarkupByCallbackQuery(callbackQuery, bot)
	case settings.IsCallbackForChangeUserData(callbackQuery.Data):
		settings.ChangeUserDataByCallbackQuery(callbackQuery, bot)
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
