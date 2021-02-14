package betypes

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// UserNil if the someone field in User struct is not set.
const UserNil = "user: nil"

// User struct for saving user data.
type User struct {
	tgbotapi.User
	Age                   string `json:"age"`
	City                  string `json:"city"`
	Sex                   string `json:"sex"`
	AgeOfTheInterlocutor  string `json:"age_of_the_interlocutor"`
	CityOfTheInterlocutor string `json:"city_of_the_interlocutor"`
	SexOfTheInterlocutor  string `json:"sex_of_the_interlocutor"`
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
	Chats      []Chat `json:"chats"`
	UsersCount int    `json:"users_count"` // UsersCount max users count in one chat.
	Queue
}
