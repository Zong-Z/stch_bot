package database

import (
	"encoding/json"
	"fmt"
	"strconv"
	"telegram-chat_bot/betypes"
	"telegram-chat_bot/logger"

	"github.com/go-redis/redis/v8"
)

const userPrefix = "USER_"

func init() {
	DB = &RedisDB{
		client: redis.NewClient(&redis.Options{
			Addr:     betypes.GetConfig().DB.Redis.Addr,
			Password: betypes.GetConfig().DB.Redis.Password,
			DB:       betypes.GetConfig().DB.Redis.Db,
		}),
	}
	DB.Ctx = DB.client.Context()
}

// SaveUser saves the user to the database.
func (db *RedisDB) SaveUser(user betypes.User) error {
	logger.ForInfo(fmt.Sprintf("Saving the user to the DB. User ID - %d", user.ID))
	j, err := json.Marshal(user)
	if err != nil {
		logger.ForWarning(fmt.Sprintf("Error %s. Could not marshal user", err.Error()))
		return err
	}

	err = db.client.Set(db.Ctx, userPrefix+strconv.FormatInt(int64(user.ID), 10), string(j), 0).Err()
	if err != nil {
		logger.ForWarning(fmt.Sprintf("Error %s. Could not save user to the DB", err.Error()))
		return err
	}

	logger.ForInfo(fmt.Sprintf("User have successfully saved to DB. ID - %d", user.ID))
	return err
}

// GetUser returns the user from the database.
func (db *RedisDB) GetUser(userID int) (*betypes.User, error) {
	logger.ForInfo(fmt.Sprintf("Getting a user from a DB. ID - %d", userID))
	r, err := db.client.Get(db.Ctx, userPrefix+strconv.FormatInt(int64(userID), 10)).Result()
	if err == redis.Nil {
		logger.ForWarning(fmt.Sprintf("User not found, %v", err))
		return nil, err
	} else if err != nil {
		logger.ForWarning(fmt.Sprintf("Could not read user from DB. ID - %d", userID))
		return nil, err
	}

	u := &betypes.User{}
	err = json.Unmarshal([]byte(r), &u)
	if err != nil {
		logger.ForWarning("Could not unmarshal user")
		return nil, err
	}

	logger.ForInfo("User successfully received from DB")
	return u, nil
}
