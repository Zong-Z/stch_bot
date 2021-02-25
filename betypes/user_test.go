package betypes

import (
	"strings"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func TestNewUser(t *testing.T) {
	u := NewUser(tgbotapi.User{})
	if !strings.EqualFold(u.Age, UserNil) || !strings.EqualFold(u.AgeOfTheInterlocutor, UserNil) ||
		!strings.EqualFold(u.City, UserNil) || !strings.EqualFold(u.CityOfTheInterlocutor, UserNil) ||
		!strings.EqualFold(u.Sex, UserNil) || !strings.EqualFold(u.SexOfTheInterlocutor, UserNil) {
		t.Error(u)
	}
}

func TestUser_IsSuitableAge(t *testing.T) {
	// Suitable cases.
	u1 := User{Age: "0", AgeOfTheInterlocutor: "0"}
	u2 := User{Age: "0", AgeOfTheInterlocutor: "0"}
	if !u1.IsSuitableAge(u2) && !u2.IsSuitableAge(u1) {
		t.Errorf("SUITABLE CASE. User 1 age: %s. User 1 interlocutor age: %s.\n"+
			"User 2 age: %s. User 2 interlocutor age: %s.", u1.Age, u1.AgeOfTheInterlocutor, u2.Age, u2.AgeOfTheInterlocutor)
	}

	u1 = User{Age: "1", AgeOfTheInterlocutor: "0"}
	u2 = User{Age: "0", AgeOfTheInterlocutor: "1"}
	if !u1.IsSuitableAge(u2) && !u2.IsSuitableAge(u1) {
		t.Errorf("SUITABLE CASE. User 1 age: %s. User 1 interlocutor age: %s.\n"+
			"User 2 age: %s. User 2 interlocutor age: %s.", u1.Age, u1.AgeOfTheInterlocutor, u2.Age, u2.AgeOfTheInterlocutor)
	}

	u1 = User{Age: "1", AgeOfTheInterlocutor: "0"}
	u2 = User{Age: "0", AgeOfTheInterlocutor: UserNil}
	if !u1.IsSuitableAge(u2) && !u2.IsSuitableAge(u1) {
		t.Errorf("SUITABLE CASE. User 1 age: %s. User 1 interlocutor age: %s.\n"+
			"User 2 age: %s. User 2 interlocutor age: %s.", u1.Age, u1.AgeOfTheInterlocutor, u2.Age, u2.AgeOfTheInterlocutor)
	}

	u1 = User{Age: UserNil, AgeOfTheInterlocutor: "0"}
	u2 = User{Age: "0", AgeOfTheInterlocutor: UserNil}
	if !u1.IsSuitableAge(u2) && !u2.IsSuitableAge(u1) {
		t.Errorf("SUITABLE CASE. User 1 age: %s. User 1 interlocutor age: %s.\n"+
			"User 2 age: %s. User 2 interlocutor age: %s.", u1.Age, u1.AgeOfTheInterlocutor, u2.Age, u2.AgeOfTheInterlocutor)
	}

	// Unsuitable cases cases.
	u1 = User{Age: "0", AgeOfTheInterlocutor: "0"}
	u2 = User{Age: "0", AgeOfTheInterlocutor: "1"}
	if u1.IsSuitableAge(u2) && u2.IsSuitableAge(u1) {
		t.Errorf("UNSUITABLE CASE. User 1 age: %s. User 1 interlocutor age: %s.\n"+
			"User 2 age: %s. User 2 interlocutor age: %s.", u1.Age, u1.AgeOfTheInterlocutor, u2.Age, u2.AgeOfTheInterlocutor)
	}

	u1 = User{Age: "0", AgeOfTheInterlocutor: "0"}
	u2 = User{Age: "1", AgeOfTheInterlocutor: "0"}
	if u1.IsSuitableAge(u2) && u2.IsSuitableAge(u1) {
		t.Errorf("UNSUITABLE CASE. User 1 age: %s. User 1 interlocutor age: %s.\n"+
			"User 2 age: %s. User 2 interlocutor age: %s.", u1.Age, u1.AgeOfTheInterlocutor, u2.Age, u2.AgeOfTheInterlocutor)
	}

	u1 = User{Age: "0", AgeOfTheInterlocutor: "0"}
	u2 = User{Age: UserNil, AgeOfTheInterlocutor: "0"}
	if u1.IsSuitableAge(u2) && u2.IsSuitableAge(u1) {
		t.Errorf("UNSUITABLE CASE. User 1 age: %s. User 1 interlocutor age: %s.\n"+
			"User 2 age: %s. User 2 interlocutor age: %s.", u1.Age, u1.AgeOfTheInterlocutor, u2.Age, u2.AgeOfTheInterlocutor)
	}
}

func TestUser_IsSuitableCity(t *testing.T) {
	// Suitable cases.
	u1 := User{City: "0", CityOfTheInterlocutor: "0"}
	u2 := User{City: "0", CityOfTheInterlocutor: "0"}
	if !u1.IsSuitableCity(u2) && !u2.IsSuitableCity(u1) {
		t.Errorf("SUITABLE CASE. User 1 city: %s. User 1 interlocutor city: %s.\n"+
			"User 2 city: %s. User 2 interlocutor city: %s.", u1.City, u1.CityOfTheInterlocutor, u2.City, u2.CityOfTheInterlocutor)
	}

	u1 = User{City: "1", CityOfTheInterlocutor: "0"}
	u2 = User{City: "0", CityOfTheInterlocutor: "1"}
	if !u1.IsSuitableCity(u2) && !u2.IsSuitableCity(u1) {
		t.Errorf("SUITABLE CASE. User 1 city: %s. User 1 interlocutor city: %s.\n"+
			"User 2 city: %s. User 2 interlocutor city: %s.", u1.City, u1.CityOfTheInterlocutor, u2.City, u2.CityOfTheInterlocutor)
	}

	u1 = User{City: "1", CityOfTheInterlocutor: "0"}
	u2 = User{City: "0", CityOfTheInterlocutor: UserNil}
	if !u1.IsSuitableCity(u2) && !u2.IsSuitableCity(u1) {
		t.Errorf("SUITABLE CASE. User 1 city: %s. User 1 interlocutor city: %s.\n"+
			"User 2 city: %s. User 2 interlocutor city: %s.", u1.City, u1.CityOfTheInterlocutor, u2.City, u2.CityOfTheInterlocutor)
	}

	u1 = User{City: UserNil, CityOfTheInterlocutor: "0"}
	u2 = User{City: "0", CityOfTheInterlocutor: UserNil}
	if !u1.IsSuitableCity(u2) && !u2.IsSuitableCity(u1) {
		t.Errorf("SUITABLE CASE. User 1 city: %s. User 1 interlocutor city: %s.\n"+
			"User 2 city: %s. User 2 interlocutor city: %s.", u1.City, u1.CityOfTheInterlocutor, u2.City, u2.CityOfTheInterlocutor)
	}

	// Unsuitable cases cases.
	u1 = User{City: "0", CityOfTheInterlocutor: "0"}
	u2 = User{City: "0", CityOfTheInterlocutor: "1"}
	if u1.IsSuitableCity(u2) && u2.IsSuitableCity(u1) {
		t.Errorf("UNSUITABLE CASE. User 1 city: %s. User 1 interlocutor city: %s.\n"+
			"User 2 city: %s. User 2 interlocutor city: %s.", u1.City, u1.CityOfTheInterlocutor, u2.City, u2.CityOfTheInterlocutor)
	}

	u1 = User{City: "0", CityOfTheInterlocutor: "0"}
	u2 = User{City: "1", CityOfTheInterlocutor: "0"}
	if u1.IsSuitableCity(u2) && u2.IsSuitableCity(u1) {
		t.Errorf("UNSUITABLE CASE. User 1 city: %s. User 1 interlocutor city: %s.\n"+
			"User 2 city: %s. User 2 interlocutor city: %s.", u1.City, u1.CityOfTheInterlocutor, u2.City, u2.CityOfTheInterlocutor)
	}

	u1 = User{City: "0", CityOfTheInterlocutor: "0"}
	u2 = User{City: UserNil, CityOfTheInterlocutor: "0"}
	if u1.IsSuitableCity(u2) && u2.IsSuitableCity(u1) {
		t.Errorf("UNSUITABLE CASE. User 1 city: %s. User 1 interlocutor city: %s.\n"+
			"User 2 city: %s. User 2 interlocutor city: %s.", u1.City, u1.CityOfTheInterlocutor, u2.City, u2.CityOfTheInterlocutor)
	}
}

func TestUser_IsSuitableSex(t *testing.T) {
	// Suitable cases.
	u1 := User{Sex: "0", SexOfTheInterlocutor: "0"}
	u2 := User{Sex: "0", SexOfTheInterlocutor: "0"}
	if !u1.IsSuitableSex(u2) && !u2.IsSuitableSex(u1) {
		t.Errorf("SUITABLE CASE. User 1 sex: %s. User 1 interlocutor sex: %s.\n"+
			"User 2 sex: %s. User 2 interlocutor sex: %s.", u1.Sex, u1.SexOfTheInterlocutor, u2.Sex, u2.SexOfTheInterlocutor)
	}

	u1 = User{Sex: "1", SexOfTheInterlocutor: "0"}
	u2 = User{Sex: "0", SexOfTheInterlocutor: "1"}
	if !u1.IsSuitableSex(u2) && !u2.IsSuitableSex(u1) {
		t.Errorf("SUITABLE CASE. User 1 sex: %s. User 1 interlocutor sex: %s.\n"+
			"User 2 sex: %s. User 2 interlocutor sex: %s.", u1.Sex, u1.SexOfTheInterlocutor, u2.Sex, u2.SexOfTheInterlocutor)
	}

	u1 = User{Sex: "1", SexOfTheInterlocutor: "0"}
	u2 = User{Sex: "0", SexOfTheInterlocutor: UserNil}
	if !u1.IsSuitableSex(u2) && !u2.IsSuitableSex(u1) {
		t.Errorf("SUITABLE CASE. User 1 sex: %s. User 1 interlocutor sex: %s.\n"+
			"User 2 sex: %s. User 2 interlocutor sex: %s.", u1.Sex, u1.SexOfTheInterlocutor, u2.Sex, u2.SexOfTheInterlocutor)
	}

	u1 = User{Sex: UserNil, SexOfTheInterlocutor: "0"}
	u2 = User{Sex: "0", SexOfTheInterlocutor: UserNil}
	if !u1.IsSuitableSex(u2) && !u2.IsSuitableSex(u1) {
		t.Errorf("SUITABLE CASE. User 1 sex: %s. User 1 interlocutor sex: %s.\n"+
			"User 2 sex: %s. User 2 interlocutor sex: %s.", u1.Sex, u1.SexOfTheInterlocutor, u2.Sex, u2.SexOfTheInterlocutor)
	}

	// Unsuitable cases cases.
	u1 = User{Sex: "0", SexOfTheInterlocutor: "0"}
	u2 = User{Sex: "0", SexOfTheInterlocutor: "1"}
	if u1.IsSuitableSex(u2) && u2.IsSuitableSex(u1) {
		t.Errorf("UNSUITABLE CASE. User 1 sex: %s. User 1 interlocutor sex: %s.\n"+
			"User 2 sex: %s. User 2 interlocutor sex: %s.", u1.Sex, u1.SexOfTheInterlocutor, u2.Sex, u2.SexOfTheInterlocutor)
	}

	u1 = User{Sex: "0", SexOfTheInterlocutor: "0"}
	u2 = User{Sex: "1", SexOfTheInterlocutor: "0"}
	if u1.IsSuitableSex(u2) && u2.IsSuitableSex(u1) {
		t.Errorf("UNSUITABLE CASE. User 1 sex: %s. User 1 interlocutor sex: %s.\n"+
			"User 2 sex: %s. User 2 interlocutor sex: %s.", u1.Sex, u1.SexOfTheInterlocutor, u2.Sex, u2.SexOfTheInterlocutor)
	}

	u1 = User{Sex: "0", SexOfTheInterlocutor: "0"}
	u2 = User{Sex: UserNil, SexOfTheInterlocutor: "0"}
	if u1.IsSuitableSex(u2) && u2.IsSuitableSex(u1) {
		t.Errorf("UNSUITABLE CASE. User 1 sex: %s. User 1 interlocutor sex: %s.\n"+
			"User 2 sex: %s. User 2 interlocutor sex: %s.", u1.Sex, u1.SexOfTheInterlocutor, u2.Sex, u2.SexOfTheInterlocutor)
	}
}
