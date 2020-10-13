package betypes

import (
	"fmt"
	"telegram-chat_bot/loger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewChats returns a pointer to new Сhats.
func NewChats(usersCount, queueSize int, bot *tgbotapi.BotAPI) *Chats {
	var chats Chats
	chats.init(usersCount, queueSize, bot)
	return &chats
}

func (chats *Chats) init(usersCountInChat, queueSize int, bot *tgbotapi.BotAPI) {
	chats.Chats = make([]Chat, 0)
	chats.User = make(chan User, queueSize)

	go func() {
		chats.Chats = append(chats.Chats, Chat{})
		for user := range chats.User {
			if len(chats.Chats[len(chats.Chats)-1].Users) < usersCountInChat {
				chats.Chats[len(chats.Chats)-1].Users = append(chats.Chats[len(chats.Chats)-1].Users, user)
			}

			if len(chats.Chats[len(chats.Chats)-1].Users) == usersCountInChat {
				for _, u := range chats.Chats[len(chats.Chats)-1].Users {
					if _, err := bot.Send(tgbotapi.MessageConfig{
						BaseChat: tgbotapi.BaseChat{
							ChatID: int64(u.ID),
						},
						Text:      "*The interlocutor is found.\nWho will be the first?*",
						ParseMode: "MARKDOWN",
					}); err != nil {
						loger.ForLog(fmt.Sprintf("Error %s, sending message, User ID - %d", err.Error(), u.ID))
					}
				}
				chats.Chats = append(chats.Chats, Chat{})
			}
		}
	}()
}

// AddUserToQueue adds a user to the queue.
func (chats *Chats) AddUserToQueue(user *User, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: int64(user.ID),
		},
		Text:      "*You are already added to the queue.*",
		ParseMode: "MARKDOWN",
	}
	if !chats.isTheUserInChat(user.ID) {
		chats.User <- *user
		msg.Text = "*You are successfully added to the queue.\nPlease wait...\nWe are looking for your interlocutor...*"
	}

	if _, err := bot.Send(msg); err != nil {
		loger.ForLog(fmt.Sprintf("Error %s, sending message. User ID - %d.", err.Error(), user.ID))
	}
}

// SendMessageToInterlocutors sends messages to all іnterlocutors of the user
func (chats *Chats) SendMessageToInterlocutors(message string, userID int, bot *tgbotapi.BotAPI) {
	users := chats.getUserInterlocutors(userID)
	if users != nil {
		for i, u := range *users {
			if _, err := bot.Send(tgbotapi.MessageConfig{
				BaseChat: tgbotapi.BaseChat{
					ChatID: int64(u.ID),
				},
				Text:      fmt.Sprintf("*Message forom interlocutor %d:* %s", i+1, message),
				ParseMode: "MARKDOWN",
			}); err != nil {
				loger.ForLog(fmt.Sprintf("Error %s, sending message. User ID - %d.", err.Error(), userID))
			}
		}
	}
}

func (chats *Chats) getUserInterlocutors(userID int) *[]*User {
	users := make([]*User, 0)
	if chats.isTheUserInChat(userID) {
		chat := chats.getChatByUserID(userID)
		for i := 0; i < len(chat.Users); i++ {
			if chat.Users[i].ID != userID {
				users = append(users, &chat.Users[i])
			}
		}
		return &users
	}
	return nil
}

func (chats *Chats) getChatByUserID(userID int) *Chat {
	for i := 0; i < len(chats.Chats); i++ {
		for j := 0; j < len(chats.Chats[i].Users); j++ {
			if chats.Chats[i].Users[j].ID == userID {
				return &chats.Chats[i]
			}
		}
	}
	return nil
}

func (chats *Chats) isTheUserInChat(userID int) bool {
	for _, chat := range chats.Chats {
		for _, user := range chat.Users {
			if user.ID == userID {
				return true
			}
		}
	}
	return false
}
