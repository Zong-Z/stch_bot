package betypes

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// UserNil if the someone field is not set.
const UserNil = "user: nil"

// User struct for saving user data.
type User struct {
	tgbotapi.User
	Age  string `json:"age"`
	City string `json:"city"`
}
