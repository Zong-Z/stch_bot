package betypes

import (
	"fmt"
	"strings"
	"telegram-chat_bot/logger"

	"github.com/google/uuid"
)

// NewChats return new Chats.
//
// In goroutine starts a loop that adds users(User) to the chat([]Chat).
// usersCount saves the maximum number of users in the chat.
func NewChats(buffer, usersCount int) *Chats {
	chats := &Chats{
		Chats:      make([]Chat, 0),
		UsersCount: usersCount,
		UsersQueue: make(chan User, buffer),
	}

	go chats.startAddingUsers()

	return chats
}

// startAddingUsers start adding user to chats.
//
// It is recommended to run in goroutine.
// Gets User from UsersQueue and adds
// new chats(To []Chat) to which User are added depending on age, city, etc...
func (chats *Chats) startAddingUsers() {
	for user := range chats.UsersQueue {
		UUID := uuid.New().String()
		if strings.EqualFold(UUID, "") {
			err := fmt.Sprintf("Error, bad UUID - %s", UUID)
			logger.ForLog(err)
			panic(err)
		}

		if len(chats.Chats) == 0 {
			chats.Chats = append(chats.Chats, Chat{Users: make([]User, 0), ID: UUID})

			// We are adding a user to the last added suitableChats.
			chats.Chats[len(chats.Chats)-1].Users = append(chats.Chats[len(chats.Chats)-1].Users, user)
			continue
		}

		// Search for a suitableChats to which you can add a user.
		suitableChats := chats.findAChatsForTheUser(user)
		if suitableChats != nil {
			// We add the user to the first suitable chat.
			suitableChats[0].Users = append(suitableChats[0].Users, user)
			continue
		}

		chats.Chats = append(chats.Chats, Chat{Users: make([]User, 0), ID: UUID})

		// We are adding a user to the last added suitable chats.
		chats.Chats[len(chats.Chats)-1].Users = append(chats.Chats[len(chats.Chats)-1].Users, user)
	}
}

// findAChatForTheUser search for a chats([]*Chat) to which you can add a user.
//
// Takes into account the age, city and more.
// If the chats is not found, returns nil.
// Returns chats([]*Chat) that are suitable for User.
func (chats *Chats) findAChatsForTheUser(user User) []*Chat {
	suitableChats := make([]*Chat, 0)
	for i := 0; i < len(chats.Chats); i++ {
		// If the chat is crowded, it is not suitable.
		if int64(len(chats.Chats[i].Users)) == int64(chats.UsersCount) {
			continue
		}

		// If there are no users in the chat, then it is suitable.
		if len(chats.Chats[i].Users) == 0 {
			suitableChats = append(suitableChats, &chats.Chats[i])
			continue
		}

		// If there is a user with the correct age and city, the chat is appropriate.
		if chats.Chats[i].Users[0].Age == user.Age && chats.Chats[i].Users[0].City == user.City {
			suitableChats = append(suitableChats, &chats.Chats[i])
		}
	}

	if len(suitableChats) == 0 {
		return nil
	}

	return suitableChats
}

// AddUserToTheQueue adds a user to the queue(UsersQueue).
func (chats *Chats) AddUserToTheQueue(user User) {
	chats.UsersQueue <- user
}

// StopChatting stop chat with interlocutors.
//
// Find user in chat.
// If user found delete chat with user.
func (chats *Chats) StopChatting(userID int) {
	userChat := chats.GetChatByUserID(userID)
	if userChat == nil {
		return
	}

	// We are looking for a chat that has a user, and delete it.
	for i, chat := range chats.Chats {
		if userChat.ID == chat.ID {
			chats.Chats[len(chats.Chats)-1], chats.Chats[i] = chats.Chats[i], chats.Chats[len(chats.Chats)-1]
			chats.Chats = chats.Chats[:len(chats.Chats)-1]

			break
		}
	}
}

// GetChatByUserID return chat by user ID.
//
// Return nil if chat is not found.
func (chats *Chats) GetChatByUserID(userID int) *Chat {
	for _, chat := range chats.Chats {
		for _, user := range chat.Users {
			if user.ID == userID {
				return &chat
			}
		}
	}

	return nil
}

// GetUserInterlocutors returns the user's interlocutors.
//
// Returns a nil when the user is not in the chat.
func (chats *Chats) GetUserInterlocutors(userID int) []User {
	chat := chats.GetChatByUserID(userID)
	if chat != nil { // If the chat was not found.
		interlocutors := make([]User, 0)
		for _, user := range chat.Users {
			if userID != user.ID {
				interlocutors = append(interlocutors, user)
			}
		}

		if len(interlocutors) == 0 {
			return nil
		}

		return interlocutors
	}

	return nil
}

// UserIsChatting if the user is chatting, returns the true.
func (chats *Chats) UserIsChatting(userID int) bool {
	for _, c := range chats.Chats {
		for _, u := range c.Users {
			if u.ID == userID {
				return true
			}
		}
	}

	return false
}
