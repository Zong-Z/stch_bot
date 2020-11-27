package betypes

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Users const's.
const (
	// UserNil if the someone field in User struct is not set.
	UserNil = "user: nil"
	// UserAgeSixteenAndBelow = "user age: sixteen_and_below"
	// UserAgeEighteenOrMore  = "user age: eighteen_or_more"
)

// User struct for saving user data.
type User struct {
	tgbotapi.User
	Age  string `json:"age"`
	City string `json:"city"`
}

// Chat struct for saving users([]User) in chat.
type Chat struct {
	Users []User `json:"users"`
	ID    string `json:"id"`
}

// UsersQueue users queue.
type UsersQueue chan User

// Chats struct for saving all Chat with User queue.
type Chats struct {
	Chats      []Chat `json:"chat"`
	UsersCount int    `json:"users_count"` // UsersCount max users count in one chat.
	UsersQueue
}
