package betypes

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// UserNil If there is a field in the User structure that is not set, you can use this constant.
const UserNil = "user: nil"

// User is a structure for storing user information.
type User struct {
	tgbotapi.User
	Age                   string `json:"age"`
	City                  string `json:"city"`
	Sex                   string `json:"sex"`
	AgeOfTheInterlocutor  string `json:"age_of_the_interlocutor"`
	CityOfTheInterlocutor string `json:"city_of_the_interlocutor"`
	SexOfTheInterlocutor  string `json:"sex_of_the_interlocutor"`
}

// Chat is a structure that stores the users who are in the chat and the chat identifier.
type Chat struct {
	Users []User `json:"users"`
	ID    string `json:"id"`
}

// Queue user queue.
type Queue chan User

// Chats has a queue and an array of chats for a certain number of users.
type Chats struct {
	Chats      []Chat `json:"chats"`
	UsersCount int    `json:"users_count"` // Maximum number of users in one chat.
	Queue
}
