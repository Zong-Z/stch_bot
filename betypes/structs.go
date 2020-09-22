package betypes

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// If the someone field is not set
const UserNull = "User: NULL"

type User struct {
	tgbotapi.User
	Age  string `json:"age"`
	City string `json:"city"`
}
