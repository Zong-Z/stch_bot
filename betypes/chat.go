package betypes

import (
	"strings"
	"telegram-chat_bot/logger"

	"github.com/google/uuid"
)

// NewChat returns the *Chat with the UUID.
func NewChat() *Chat {
	UUID := uuid.New().String()
	if strings.EqualFold(UUID, "") {
		err := "Error, bad UUID."
		logger.ForInfo(err)
		panic(err)
	}

	return &Chat{
		Users: make([]User, 0),
		ID:    UUID,
	}
}

// AddUser adds a user to the chat.
func (c *Chat) AddUser(user User) {
	c.Users = append(c.Users, user)
}

// DeleteUserFromChat removes the user from the chat room he is in.
//
// If the user is not in the chat, does nothing.
func (c *Chat) DeleteUserFromChat(userID int) {
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].ID == userID {
			c.Users[len(c.Users)-1], c.Users[i] = c.Users[i], c.Users[len(c.Users)-1]
			c.Users = c.Users[:len(c.Users)-1]
			break
		}
	}
}

// GetInterlocutorsByUserID returns all interlocutors by user ID.
//
// If the user is not in the chat, or there are no interlocutors yet, returns nil.
func (c *Chat) GetInterlocutorsByUserID(userID int) []User {
	if !c.IsUserInChat(userID) {
		return nil
	}

	users := make([]User, 0)
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].ID != userID {
			users = append(users, c.Users[i])
		}
	}

	if len(users) != 0 {
		return users
	}

	return nil
}

// IsUserInChat returns true if the user is in the chat.
func (c *Chat) IsUserInChat(userID int) bool {
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].ID == userID {
			return true
		}
	}
	return false
}
