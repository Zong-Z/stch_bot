package betypes

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// UserNil if the someone field in User struct is not set.
const UserNil = "user: nil"

// User struct for saving user data.
type User struct {
	tgbotapi.User
	Age  string `json:"age"`
	City string `json:"city"`
}

// Chat struct for saving users([]User) in chat.
type Chat struct {
	Users []User `json:"users"`
	ID    uint   `json:"id"`
}

// UsersQueue users queue.
type UsersQueue struct {
	User   chan User `json:"user"`
	Buffer int       `json:"buffer"`
}

// Chats struct for saving all Chat with User queue.
type Chats struct {
	Chats []Chat     `json:"chats"`
	Queue UsersQueue `json:"queue"`
}
