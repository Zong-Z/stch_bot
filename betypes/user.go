package betypes

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewUser return new *User.
func NewUser(user tgbotapi.User) *User {
	return &User{
		User: user,
		Age:  UserNil, City: UserNil, Sex: UserNil,
		AgeOfTheInterlocutor: UserNil, CityOfTheInterlocutor: UserNil, SexOfTheInterlocutor: UserNil,
	}
}

// IsSuitableAge equals the user's age from the structure with the age of the transferred user.
func (u User) IsSuitableAge(user User) bool {
	return strings.EqualFold(user.AgeOfTheInterlocutor, u.Age) && strings.EqualFold(u.AgeOfTheInterlocutor, user.Age) ||
		strings.EqualFold(user.AgeOfTheInterlocutor, UserNil) && strings.EqualFold(u.AgeOfTheInterlocutor, user.Age) ||
		strings.EqualFold(user.AgeOfTheInterlocutor, u.Age) && strings.EqualFold(u.AgeOfTheInterlocutor, UserNil)
}

// IsSuitableCity equals the user's city from the structure with the city of the transferred user.
func (u User) IsSuitableCity(user User) bool {
	return strings.EqualFold(user.CityOfTheInterlocutor, u.City) && strings.EqualFold(u.CityOfTheInterlocutor, user.City) ||
		strings.EqualFold(user.CityOfTheInterlocutor, UserNil) && strings.EqualFold(u.CityOfTheInterlocutor, user.City) ||
		strings.EqualFold(user.CityOfTheInterlocutor, u.City) && strings.EqualFold(u.CityOfTheInterlocutor, UserNil)
}

// IsSuitableSex equals the user's sex from the structure with the sex of the transferred user.
func (u User) IsSuitableSex(user User) bool {
	return strings.EqualFold(user.SexOfTheInterlocutor, u.Sex) && strings.EqualFold(u.SexOfTheInterlocutor, user.Sex) ||
		strings.EqualFold(user.SexOfTheInterlocutor, UserNil) && strings.EqualFold(u.SexOfTheInterlocutor, user.Sex) ||
		strings.EqualFold(user.SexOfTheInterlocutor, u.Sex) && strings.EqualFold(u.SexOfTheInterlocutor, UserNil)
}
