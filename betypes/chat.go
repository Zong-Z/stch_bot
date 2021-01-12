package betypes

import (
	"strings"
	"telegram-chat_bot/logger"

	"github.com/google/uuid"
)

// NewChat return new *Chat with unique UUID.
func NewChat() *Chat {
	UUID := uuid.New().String()
	if strings.EqualFold(UUID, "") {
		err := "Error, bad UUID."
		logger.ForLog(err)
		panic(err)
	}

	return &Chat{
		Users: make([]User, 0),
		ID:    UUID,
	}
}

// AddUser add User to Chat.
func (c *Chat) AddUser(user User) {
	c.Users = append(c.Users, user)
}

// DeleteUserFromChat removes the user from the chat.
func (c *Chat) DeleteUserFromChat(userID int) {
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].ID == userID {
			c.Users[len(c.Users)-1], c.Users[i] = c.Users[i], c.Users[len(c.Users)-1]
			c.Users = c.Users[:len(c.Users)-1]
			break
		}
	}
}

// GetUserInterlocutors returns User interlocutors if User is in Chat.
func (c *Chat) GetUserInterlocutors(userID int) []User {
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

// IsUserInChat returns true if user is in chat.
func (c *Chat) IsUserInChat(userID int) bool {
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].ID == userID {
			return true
		}
	}
	return false
}
