package betypes

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewUser return new *User.
func NewUser(user tgbotapi.User) *User {
	return &User{
		User: user,
		Age:  UserNil, City: UserNil,
		InterlocutorAge: UserNil, InterlocutorCity: UserNil,
	}
}

// IsSuitableAge equals the user's age from the structure with the age of the transferred user.
func (u User) IsSuitableAge(user User) bool {
	return strings.EqualFold(user.InterlocutorAge, u.Age) && strings.EqualFold(u.InterlocutorAge, user.Age) ||
		strings.EqualFold(user.InterlocutorAge, UserNil) && strings.EqualFold(u.InterlocutorAge, user.Age) ||
		strings.EqualFold(user.InterlocutorAge, u.Age) && strings.EqualFold(u.InterlocutorAge, UserNil)
}

// IsSuitableCity equals the user's city from the structure with the city of the transferred user.
func (u User) IsSuitableCity(user User) bool {
	return strings.EqualFold(user.InterlocutorCity, u.City) && strings.EqualFold(u.InterlocutorCity, user.City) ||
		strings.EqualFold(user.InterlocutorCity, UserNil) && strings.EqualFold(u.InterlocutorCity, user.City) ||
		strings.EqualFold(user.InterlocutorCity, u.City) && strings.EqualFold(u.InterlocutorCity, UserNil)
}
