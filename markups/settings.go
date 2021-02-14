package markups

import (
	"strings"
	"telegram-chat_bot/betypes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SettingsReplyMarkups' constants.
const (
	SettingsReplyMarkupName   = "SETTINGS"
	SettingsReplyMarkupPrefix = SettingsReplyMarkupName + "_"
)

// SettingsReplyMarkups' age constants.
const (
	OwnAgeReplyMarkupName     = "OWN AGE"
	OwnAgeReplyMarkupCallback = OwnAgeReplyMarkupName
	OwnAgePrefix              = OwnAgeReplyMarkupCallback + "_"

	AgeOfTheInterlocutorReplyMarkupName     = "AGE OF THE INTERLOCUTOR"
	AgeOfTheInterlocutorReplyMarkupCallback = AgeOfTheInterlocutorReplyMarkupName
	AgeOfTheInterlocutorPrefix              = AgeOfTheInterlocutorReplyMarkupCallback + "_"

	SixteenOrLessCallback           = "SIXTEEN OR LESS"
	FromSixteenToEighteenCallback   = "FROM SIXTEEN TO EIGHTEEN"
	FromEighteenToTwentyOneCallback = "FROM EIGHTEEN TO TWENTY ONE"

	OwnAgeText                  = "📅OWN AGE"
	AgeOfTheInterlocutorText    = "🙍‍♂🙍‍♀📅AGE OF THE INTERLOCUTOR"
	SixteenOrLessText           = "SIXTEEN OR LESS"
	FromSixteenToEighteenText   = "FROM SIXTEEN TO EIGHTEEN"
	FromEighteenToTwentyOneText = "FROM EIGHTEEN TO TWENTY ONE"
)

// SettingsReplyMarkups' city constants.
const (
	OwnCityReplyMarkupName     = "OWN CITY"
	OwnCityReplyMarkupCallback = OwnCityReplyMarkupName
	OwnCityPrefix              = OwnCityReplyMarkupCallback + "_"

	CityOfTheInterlocutorReplyMarkupName     = "CITY OF THE INTERLOCUTOR"
	CityOfTheInterlocutorReplyMarkupCallback = CityOfTheInterlocutorReplyMarkupName
	CityOfTheInterlocutorPrefix              = CityOfTheInterlocutorReplyMarkupCallback + "_"

	OwnCityText               = "🌍OWN CITY"
	CityOfTheInterlocutorText = "🙍‍♂🙍‍♀🌍CITY OF THE INTERLOCUTOR"
)

// SettingsReplyMarkups' cities.
const (
	Lviv     = "LVIV"
	Ternopil = "TERNOPIL"
	Kiev     = "KIEV"
)

// SexReplyMarkups' sex constants.
const (
	OwnSexReplyMarkupName     = "OWN SEX"
	OwnSexReplyMarkupCallback = OwnSexReplyMarkupName
	OwnSexPrefix              = OwnSexReplyMarkupCallback + "_"

	SexOfTheInterlocutorReplyMarkupName     = "SEX OF THE INTERLOCUTOR"
	SexOfTheInterlocutorReplyMarkupCallback = SexOfTheInterlocutorReplyMarkupName
	SexOfTheInterlocutorPrefix              = SexOfTheInterlocutorReplyMarkupCallback + "_"

	OwnSexText               = "🙍‍♂🙍‍♀OWN SEX"
	SexOfTheInterlocutorText = "🙍‍♂🙍‍♀SEX OF THE INTERLOCUTOR"
	MaleText                 = "🙍‍♂MALE"
	FemaleText               = "🙍‍♀FEMALE"
	MaleCallback             = "MALE"
	FemaleCallback           = "FEMALE"
)

var (
	ownAgeInlineKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(SixteenOrLessText,
					SettingsReplyMarkupPrefix+OwnAgePrefix+SixteenOrLessCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(FromSixteenToEighteenText,
					SettingsReplyMarkupPrefix+OwnAgePrefix+FromSixteenToEighteenCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(FromEighteenToTwentyOneText,
					SettingsReplyMarkupPrefix+OwnAgePrefix+FromEighteenToTwentyOneCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(DoesNotMatter,
					SettingsReplyMarkupPrefix+OwnAgePrefix+betypes.UserNil,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Back, SettingsReplyMarkupPrefix+SettingsReplyMarkupName),
			),
		},
	}
	ageOfTheInterlocutorInlineKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(SixteenOrLessText,
					SettingsReplyMarkupPrefix+AgeOfTheInterlocutorPrefix+SixteenOrLessCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(FromSixteenToEighteenText,
					SettingsReplyMarkupPrefix+AgeOfTheInterlocutorPrefix+FromSixteenToEighteenCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(FromEighteenToTwentyOneText,
					SettingsReplyMarkupPrefix+AgeOfTheInterlocutorPrefix+FromEighteenToTwentyOneCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(DoesNotMatter,
					SettingsReplyMarkupPrefix+AgeOfTheInterlocutorPrefix+betypes.UserNil,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Back, SettingsReplyMarkupPrefix+SettingsReplyMarkupName),
			),
		},
	}
	ownCityInlineKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Lviv,
					SettingsReplyMarkupPrefix+OwnCityPrefix+Lviv,
				),
				tgbotapi.NewInlineKeyboardButtonData(Ternopil,
					SettingsReplyMarkupPrefix+OwnCityPrefix+Ternopil,
				),
				tgbotapi.NewInlineKeyboardButtonData(Kiev,
					SettingsReplyMarkupPrefix+OwnCityPrefix+Kiev,
				),
				tgbotapi.NewInlineKeyboardButtonData(DoesNotMatter,
					SettingsReplyMarkupPrefix+OwnCityPrefix+betypes.UserNil,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Back, SettingsReplyMarkupPrefix+SettingsReplyMarkupName),
			),
		},
	}
	cityOfTheInterlocutorInlineKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Lviv,
					SettingsReplyMarkupPrefix+CityOfTheInterlocutorPrefix+Lviv,
				),
				tgbotapi.NewInlineKeyboardButtonData(Ternopil,
					SettingsReplyMarkupPrefix+CityOfTheInterlocutorPrefix+Ternopil,
				),
				tgbotapi.NewInlineKeyboardButtonData(Kiev,
					SettingsReplyMarkupPrefix+CityOfTheInterlocutorPrefix+Kiev,
				),
				tgbotapi.NewInlineKeyboardButtonData(DoesNotMatter,
					SettingsReplyMarkupPrefix+CityOfTheInterlocutorPrefix+betypes.UserNil),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Back, SettingsReplyMarkupPrefix+SettingsReplyMarkupName),
			),
		},
	}
	ownSexInlineKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(MaleText,
					SettingsReplyMarkupPrefix+OwnSexPrefix+MaleCallback,
				),
				tgbotapi.NewInlineKeyboardButtonData(FemaleText,
					SettingsReplyMarkupPrefix+OwnSexPrefix+FemaleCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(DoesNotMatter,
					SettingsReplyMarkupPrefix+OwnSexPrefix+betypes.UserNil,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Back, SettingsReplyMarkupPrefix+SettingsReplyMarkupName),
			),
		},
	}
	sexOfTheInterlocutorInlineKeyboardMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(MaleText,
					SettingsReplyMarkupPrefix+SexOfTheInterlocutorPrefix+MaleCallback,
				),
				tgbotapi.NewInlineKeyboardButtonData(FemaleText,
					SettingsReplyMarkupPrefix+SexOfTheInterlocutorPrefix+FemaleCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(DoesNotMatter,
					SettingsReplyMarkupPrefix+SexOfTheInterlocutorPrefix+betypes.UserNil,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(Back, SettingsReplyMarkupPrefix+SettingsReplyMarkupName),
			),
		},
	}
)

var settings = ReplyMarkups{
	ReplyMarkup{
		Name: SettingsReplyMarkupPrefix + SettingsReplyMarkupName,
		InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(OwnSexText,
					SettingsReplyMarkupPrefix+OwnSexReplyMarkupCallback,
				),
				tgbotapi.NewInlineKeyboardButtonData(SexOfTheInterlocutorText,
					SettingsReplyMarkupPrefix+SexOfTheInterlocutorReplyMarkupCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(OwnAgeText,
					SettingsReplyMarkupPrefix+OwnAgeReplyMarkupCallback,
				),
				tgbotapi.NewInlineKeyboardButtonData(AgeOfTheInterlocutorText,
					SettingsReplyMarkupPrefix+AgeOfTheInterlocutorReplyMarkupCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(OwnCityText,
					SettingsReplyMarkupPrefix+OwnCityReplyMarkupCallback,
				),
				tgbotapi.NewInlineKeyboardButtonData(CityOfTheInterlocutorText,
					SettingsReplyMarkupPrefix+CityOfTheInterlocutorReplyMarkupCallback,
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(CloseCallback, SettingsReplyMarkupPrefix+CloseCallback),
			),
		}},
	},
	ReplyMarkup{
		Name:                 SettingsReplyMarkupPrefix + OwnAgeReplyMarkupName,
		InlineKeyboardMarkup: ownAgeInlineKeyboardMarkup,
	},
	ReplyMarkup{
		Name:                 SettingsReplyMarkupPrefix + AgeOfTheInterlocutorReplyMarkupName,
		InlineKeyboardMarkup: ageOfTheInterlocutorInlineKeyboardMarkup,
	},
	ReplyMarkup{
		Name:                 SettingsReplyMarkupPrefix + OwnCityReplyMarkupName,
		InlineKeyboardMarkup: ownCityInlineKeyboardMarkup,
	},
	ReplyMarkup{
		Name:                 SettingsReplyMarkupPrefix + CityOfTheInterlocutorReplyMarkupName,
		InlineKeyboardMarkup: cityOfTheInterlocutorInlineKeyboardMarkup,
	},
	ReplyMarkup{
		Name:                 SettingsReplyMarkupPrefix + OwnSexReplyMarkupName,
		InlineKeyboardMarkup: ownSexInlineKeyboardMarkup,
	},
	ReplyMarkup{
		Name:                 SettingsReplyMarkupPrefix + SexOfTheInterlocutorReplyMarkupName,
		InlineKeyboardMarkup: sexOfTheInterlocutorInlineKeyboardMarkup,
	},
}

// SettingsIsThereMarkupRequest return true if callback contains OwnAgeReplyMarkupName,
// AgeOfTheInterlocutorReplyMarkupName, OwnCityReplyMarkupName, CityOfTheInterlocutorReplyMarkupName, etc...
func SettingsIsThereMarkupRequest(callbackQueryData string) bool {
	return settings.FindInlineKeyboardMarkup(callbackQueryData) != nil
}

// SettingsIsThereCallbackForChange if there is a callback to change your own age/city/sex or interlocutor, returns
// true.
func SettingsIsThereCallbackForChange(callbackQueryData string) bool {
	return strings.Contains(callbackQueryData, OwnAgePrefix) ||
		strings.Contains(callbackQueryData, AgeOfTheInterlocutorPrefix) ||
		strings.Contains(callbackQueryData, OwnCityPrefix) ||
		strings.Contains(callbackQueryData, CityOfTheInterlocutorPrefix) ||
		strings.Contains(callbackQueryData, OwnSexPrefix) ||
		strings.Contains(callbackQueryData, SexOfTheInterlocutorPrefix)
}

// GetSettings return settings(ReplyMarkup).
func GetSettings() ReplyMarkups {
	return settings
}
