package chat

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram-chat_bot/betypes"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/loger"
)

var (
	chats betypes.Chats
	id    uint64
)

func init() {
	chats.Chats = make([]betypes.Chat, 0)
	go func() {
		for {
			for user := range chats.UsersQueue {
				if len(chats.Chats[id].Users) == 2 /*TODO: chat only on two users*/ {
					id++
					chats.Chats = make([]betypes.Chat, 0)
				}
				chats.GetChatByID(id).Users = append(chats.GetChatByID(id).Users, user)
			}
		}
	}()
}

func AddUserToQueue(userID int64, bot *tgbotapi.BotAPI) {
	u, err := database.GetUser(userID)
	if err == redis.Nil /*If no user is found*/ {
		bot.Send(tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: userID,
			},
			Text:      fmt.Sprintf("*%s*", "*You are not registered.*"),
			ParseMode: "MARKDOWN",
		})
		return
	}

	if err != nil {
		loger.ForLog(fmt.Sprintf("Error, %v.", err))
	}

	bot.Send(tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: userID,
		},
		Text:      fmt.Sprintf("*@%s you are succssesfuly added to the queue.*", u.UserName),
		ParseMode: "MARKDOWN",
	})

	chats.UsersQueue <- *u
}
