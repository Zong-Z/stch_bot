package markups

import (
	"fmt"
	"strings"
	"telegram-chat_bot/betypes"
	database "telegram-chat_bot/db"
	"telegram-chat_bot/logger"

	"github.com/go-redis/redis/v8"
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
	AgeOfTheInterlocutorText    = "🙍‍♂🙍‍♀📅🌍AGE OF THE INTERLOCUTOR"
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
	SexOfTheInterlocutorText = "🙍‍♂🙍‍♀🌍SEX OF THE INTERLOCUTOR"
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

var settings = Markups{
	Markup{
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
	Markup{
		Name:                 SettingsReplyMarkupPrefix + OwnAgeReplyMarkupName,
		InlineKeyboardMarkup: ownAgeInlineKeyboardMarkup,
	},
	Markup{
		Name:                 SettingsReplyMarkupPrefix + AgeOfTheInterlocutorReplyMarkupName,
		InlineKeyboardMarkup: ageOfTheInterlocutorInlineKeyboardMarkup,
	},
	Markup{
		Name:                 SettingsReplyMarkupPrefix + OwnCityReplyMarkupName,
		InlineKeyboardMarkup: ownCityInlineKeyboardMarkup,
	},
	Markup{
		Name:                 SettingsReplyMarkupPrefix + CityOfTheInterlocutorReplyMarkupName,
		InlineKeyboardMarkup: cityOfTheInterlocutorInlineKeyboardMarkup,
	},
	Markup{
		Name:                 SettingsReplyMarkupPrefix + OwnSexReplyMarkupName,
		InlineKeyboardMarkup: ownSexInlineKeyboardMarkup,
	},
	Markup{
		Name:                 SettingsReplyMarkupPrefix + SexOfTheInterlocutorReplyMarkupName,
		InlineKeyboardMarkup: sexOfTheInterlocutorInlineKeyboardMarkup,
	},
}

// Settings structure, which stores the functions of the settings.
type Settings struct{}

// IsMarkupRequest return true if callback contains OwnAgeReplyMarkupName,
// AgeOfTheInterlocutorReplyMarkupName, OwnCityReplyMarkupName, CityOfTheInterlocutorReplyMarkupName, etc...
func (Settings) IsMarkupRequest(callbackQueryData string) bool {
	return settings.FindInlineKeyboard(callbackQueryData) != nil
}

// IsCallbackForChangeUserData if there is a callback to change your own age/city/sex or interlocutor,
// returns true.
func (Settings) IsCallbackForChangeUserData(callbackQueryData string) bool {
	return strings.Contains(callbackQueryData, OwnAgePrefix) ||
		strings.Contains(callbackQueryData, AgeOfTheInterlocutorPrefix) ||
		strings.Contains(callbackQueryData, OwnCityPrefix) ||
		strings.Contains(callbackQueryData, CityOfTheInterlocutorPrefix) ||
		strings.Contains(callbackQueryData, OwnSexPrefix) ||
		strings.Contains(callbackQueryData, SexOfTheInterlocutorPrefix)
}

// SendMarkupByCallbackQuery sends necessary keypad by callback.
//
// If the callback is wrong, it does nothing.
func (Settings) SendMarkupByCallbackQuery(callbackQuery tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	settings := Settings{}
	user, err := database.DB.GetUser(callbackQuery.From.ID)
	if err != nil && err.Error() == redis.Nil.Error() || user == nil {
		msg := tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(callbackQuery.From.ID)},
			Text:      betypes.GetTexts().Chat.NotRegistered,
			ParseMode: betypes.GetTexts().ParseMode,
		}

		_, err := bot.Send(msg)
		if err != nil {
			logger.ForError(err.Error())
		}

		logger.ForInfo(fmt.Sprintf("User %d, not registered.", callbackQuery.From.ID))
		return
	} else if err != nil && err.Error() != redis.Nil.Error() {
		logger.ForError(err.Error())
	}

	inlineKeyboard := settings.GetSettings().FindInlineKeyboard(callbackQuery.Data)
	if inlineKeyboard == nil {
		logger.ForWarning("Inline keyboard not found. Unknown callback.")
		return
	}

	for i, buttons := range inlineKeyboard.InlineKeyboard {
		for j, button := range buttons {
			callbackData := strings.Replace(*button.CallbackData, SettingsReplyMarkupPrefix, "", 1)
			isSelectedByUser := strings.EqualFold(callbackData, OwnAgePrefix+user.Age) ||
				strings.EqualFold(callbackData, AgeOfTheInterlocutorPrefix+user.AgeOfTheInterlocutor) ||
				strings.EqualFold(callbackData, OwnCityPrefix+user.City) ||
				strings.EqualFold(callbackData, CityOfTheInterlocutorPrefix+user.CityOfTheInterlocutor) ||
				strings.EqualFold(callbackData, OwnSexPrefix+user.Sex) ||
				strings.EqualFold(callbackData, SexOfTheInterlocutorPrefix+user.SexOfTheInterlocutor)
			if isSelectedByUser {
				inlineKeyboard.InlineKeyboard[i][j].Text = fmt.Sprintf("➡%s⬅", button.Text)
			}
		}
	}

	_, err = bot.Send(tgbotapi.NewEditMessageReplyMarkup(int64(user.ID), callbackQuery.Message.MessageID, *inlineKeyboard))
	if err != nil {
		logger.ForError(err.Error())
	}

	_, err = bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuery.ID, betypes.GetTexts().ReplyKeyboardMarkup.Opened))
	if err != nil {
		logger.ForError(err.Error())
	}
}

// ChangeUserDataByCallbackQuery changes the user settings by the callback and saves them to the database.
// If the callback is wrong, it does nothing.
func (Settings) ChangeUserDataByCallbackQuery(callbackQuery tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	haveUserSettingsBeenChanged := func() bool {
		user, err := database.DB.GetUser(callbackQuery.From.ID)
		if err != nil && err.Error() == redis.Nil.Error() || user == nil {
			msg := tgbotapi.MessageConfig{
				BaseChat:  tgbotapi.BaseChat{ChatID: int64(callbackQuery.From.ID)},
				Text:      betypes.GetTexts().Chat.NotRegistered,
				ParseMode: betypes.GetTexts().ParseMode,
			}

			_, err := bot.Send(msg)
			if err != nil {
				logger.ForError(err.Error())
			}

			logger.ForInfo(fmt.Sprintf("User %d, not registered.", callbackQuery.From.ID))
			return false
		} else if err != nil && err.Error() != redis.Nil.Error() {
			logger.ForError(err.Error())
		}

		callbackQueryData := strings.Replace(callbackQuery.Data, SettingsReplyMarkupPrefix, "", 1)
		switch {
		case strings.Contains(callbackQueryData, OwnAgePrefix):
			callbackQueryData = strings.Replace(callbackQueryData, OwnAgePrefix, "", 1)
			if user.Age == callbackQueryData {
				return true
			}

			user.Age = callbackQueryData
		case strings.Contains(callbackQueryData, AgeOfTheInterlocutorPrefix):
			callbackQueryData = strings.Replace(callbackQueryData, AgeOfTheInterlocutorPrefix, "", 1)
			if user.AgeOfTheInterlocutor == callbackQueryData {
				return true
			}

			user.AgeOfTheInterlocutor = callbackQueryData
		case strings.Contains(callbackQueryData, OwnCityPrefix):
			callbackQueryData = strings.Replace(callbackQueryData, OwnCityPrefix, "", 1)
			if user.City == callbackQueryData {
				return true
			}

			user.City = callbackQueryData
		case strings.Contains(callbackQueryData, CityOfTheInterlocutorPrefix):
			callbackQueryData = strings.Replace(callbackQueryData, CityOfTheInterlocutorPrefix, "", 1)
			if user.CityOfTheInterlocutor == callbackQueryData {
				return true
			}

			user.CityOfTheInterlocutor = callbackQueryData
		case strings.Contains(callbackQueryData, OwnSexPrefix):
			callbackQueryData = strings.Replace(callbackQueryData, OwnSexPrefix, "", 1)
			if user.Sex == callbackQueryData {
				return true
			}

			user.Sex = callbackQueryData
		case strings.Contains(callbackQueryData, SexOfTheInterlocutorPrefix):
			callbackQueryData = strings.Replace(callbackQueryData, SexOfTheInterlocutorPrefix, "", 1)
			if user.SexOfTheInterlocutor == callbackQueryData {
				return true
			}

			user.SexOfTheInterlocutor = callbackQueryData
		default:
			return false
		}

		err = database.DB.SaveUser(*user)
		if err != nil {
			logger.ForError(err.Error())
		}

		return true
	}()

	if haveUserSettingsBeenChanged {
		_, err := bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuery.ID,
			betypes.GetTexts().ReplyKeyboardMarkup.Changed))
		if err != nil {
			logger.ForError(err.Error())
		}

		settingsInlineKeyboard := settings.FindInlineKeyboard(SettingsReplyMarkupPrefix + SettingsReplyMarkupName)
		if settingsInlineKeyboard == nil {
			logger.ForWarning("Inline keyboard not found. Unknown callback.")
			return
		}

		_, err = bot.Send(tgbotapi.NewEditMessageReplyMarkup(int64(callbackQuery.From.ID),
			callbackQuery.Message.MessageID, *settingsInlineKeyboard))
		if err != nil {
			logger.ForError(err.Error())
		}
	}
}

// GetSettings return settings(Markup).
func (Settings) GetSettings() Markups {
	return settings
}
