package betypes

import (
	"fmt"
	"telegram-chat_bot/loger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewChats returns a pointer to new Chats.
func NewChats(usersCount, buffer int, bot *tgbotapi.BotAPI) *Chats {
	chats := &Chats{
		Chats: make([]Chat, 0),
		Queue: UsersQueue{
			User:   make(chan User, buffer),
			Buffer: buffer,
		},
	}

	go func(chats *Chats, usersCount int, bot *tgbotapi.BotAPI) {
		chats.Chats = append(chats.Chats, Chat{})

		for user := range chats.Queue.User {
			currentChat := &chats.Chats[len(chats.Chats)-1]

			if len(currentChat.Users) < usersCount {
				currentChat.Users = append(currentChat.Users, user)
			}

			if len(currentChat.Users) == usersCount {
				for _, user := range currentChat.Users {
					_, err := bot.Send(tgbotapi.NewMessage(int64(user.ID),
						"The interlocutor is found.\nWho will be the first?"))
					if err != nil {
						loger.ForLog(fmt.Sprintf("Error %s, sending message, User ID - %d", err.Error(), user.ID))
					}
				}

				chats.Chats = append(chats.Chats, Chat{})
			}
		}
	}(chats, usersCount, bot)

	return chats
}

// StartChatting adds a user to the chatting queue.
func (chats *Chats) StartChatting(user *User, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: int64(user.ID),
		},
		Text:      "*You are already added to the queue.*",
		ParseMode: "MARKDOWN",
	}
	if !chats.IsTheUserInChat(user.ID) {
		chats.Queue.User <- *user
		msg.Text = "*You are successfully added to the queue.\nPlease wait...\nWe are looking for your interlocutor...*"
	}

	if _, err := bot.Send(msg); err != nil {
		loger.ForLog(fmt.Sprintf("Error %s, sending message. User ID - %d.", err.Error(), user.ID))
	}
}

// StopChatting stops chatting.
func (chats *Chats) StopChatting(userID int, bot *tgbotapi.BotAPI) {
	if chats.IsTheUserInChat(userID) {
		chat := chats.getChatByUserID(userID)
		for _, user := range chat.Users {
			_, err := bot.Send(tgbotapi.MessageConfig{
				BaseChat: tgbotapi.BaseChat{
					ChatID: int64(user.ID),
				},
				Text:      "*Chat ended.*",
				ParseMode: "MARKDOWN"},
			)

			if err != nil {
				loger.ForLog(fmt.Sprintf("Error %s, sending message. User ID - %d.", err.Error(), userID))
			}
		}

		for i, c := range chats.Chats {
			if chat.ID == c.ID {
				chats.Chats[len(chats.Chats)-1], chats.Chats[i] = chats.Chats[i], chats.Chats[len(chats.Chats)-1] // Deleting chat
				chats.Chats = chats.Chats[:len(chats.Chats)-1]
				break
			}
		}
	}
}

// SendMessageToInterlocutors sends messages to all interlocutors of the user.
func (chats *Chats) SendMessageToInterlocutors(message string, userID int, bot *tgbotapi.BotAPI) {
	users := chats.getUserInterlocutors(userID)

	if users != nil {
		for i, u := range *users {
			_, err := bot.Send(tgbotapi.MessageConfig{
				BaseChat: tgbotapi.BaseChat{
					ChatID: int64(u.ID),
				},
				Text:      fmt.Sprintf("*Message from interlocutor %d:* %s", i+1, message),
				ParseMode: "MARKDOWN"},
			)

			if err != nil {
				loger.ForLog(fmt.Sprintf("Error %s, sending message. User ID - %d.", err.Error(), userID))
			}
		}
	}
}

func (chats *Chats) getUserInterlocutors(userID int) *[]*User {
	users := make([]*User, 0)
	if chats.IsTheUserInChat(userID) {
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

// IsTheUserInChat return true if user is in chat.
func (chats *Chats) IsTheUserInChat(userID int) bool {
	for _, chat := range chats.Chats {
		for _, user := range chat.Users {
			if user.ID == userID {
				return true
			}
		}
	}

	return false
}
