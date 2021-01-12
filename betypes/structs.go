package betypes

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	// UserNil if the someone field in User struct is not set.
	UserNil = "user: nil"
)

// User struct for saving user data.
type User struct {
	tgbotapi.User
	Age              string `json:"age"`
	City             string `json:"city"`
	InterlocutorAge  string `json:"interlocutor_age"`
	InterlocutorCity string `json:"interlocutor_city"`
}

// Chat struct for saving users([]User) in chat.
type Chat struct {
	Users []User `json:"users"`
	ID    string `json:"id"`
}

// Queue users queue.
type Queue chan User

// Chats struct for saving all Chat with User queue.
type Chats struct {
	Chats      []Chat `json:"chat"`
	UsersCount int    `json:"users_count"` // UsersCount max users count in one chat.
	Queue
}
