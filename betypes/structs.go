package betypes

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	// UserNil if the someone field in User struct is not set.
	UserNil = "user: nil"
)

// User struct for saving user data.
type User struct {
	tgbotapi.User
	Age  string `json:"age"`
	City string `json:"city"`
}

type Chat struct {
	Users []User `json:"users"`
	ID    uint64 `json:"id"`
}

type Chats struct {
	Chats      []Chat    `json:"chats"`
	UsersQueue chan User `json:"users_queue"`
}
