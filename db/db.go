package database

import (
	"encoding/json"
	"strconv"
	"telegram-chat_bot/betypes"
	"telegram-chat_bot/loger"

	"github.com/go-redis/redis/v8"
)

const userPrefix = "USER_"

var (
	database = redis.NewClient(&redis.Options{
		Addr:     betypes.GetBotConfig().RedisConfig.Addr,
		Password: betypes.GetBotConfig().RedisConfig.Password,
		DB:       betypes.GetBotConfig().RedisConfig.DB,
	})

	ctx = database.Context()
)

func SaveUser(user betypes.User) error {
	loger.ForLog("Saving the user to the DB, ID", user.ID)
	j, err := json.Marshal(user)
	if err != nil {
		loger.ForLog("Error, could not marshal user.", err)
		return err
	}

	err = database.Set(ctx, userPrefix+strconv.FormatInt(int64(user.ID), 10), string(j), 0).Err()
	if err != nil {
		loger.ForLog("Error, could not save user to the DB.", err)
		return err
	}

	loger.ForLog("User have successfully saved to DB, ID", user.ID)
	return err
}

func GetUser(userID int) (*betypes.User, error) {
	loger.ForLog("Getting a user from a DB, ID", userID)
	u := &betypes.User{}
	r, err := database.Get(ctx, userPrefix+strconv.FormatInt(int64(userID), 10)).Result()
	if err == redis.Nil {
		loger.ForLog("User not found.", err)
		return nil, err
	} else if err != nil {
		loger.ForLog("Error, could not read user from DB, ID", userID)
		return nil, err
	}

	err = json.Unmarshal([]byte(r), &u)
	if err != nil {
		loger.ForLog("Error, could not unmarshal user.")
		return nil, err
	}

	loger.ForLog("User successfully received from DB.")
	return u, nil
}
