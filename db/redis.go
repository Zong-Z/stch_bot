package database

import (
	"encoding/json"
	"strconv"
	"telegram-chat_bot/betypes"
	"telegram-chat_bot/logger"

	"github.com/go-redis/redis/v8"
)

// DBRedis redis database.
var DBRedis *RedisDB

func init() {
	DBRedis = &RedisDB{
		client: redis.NewClient(&redis.Options{
			Addr:     betypes.GetBotConfig().RedisConfig.Addr,
			Password: betypes.GetBotConfig().RedisConfig.Password,
			DB:       betypes.GetBotConfig().RedisConfig.DB,
		}),
	}
	DBRedis.Ctx = DBRedis.client.Context()
}

// SaveUser save the user in the database.
func (db *RedisDB) SaveUser(user betypes.User) error {
	logger.ForLog("Saving the user to the DBRedis, ID", user.ID)
	j, err := json.Marshal(user)
	if err != nil {
		logger.ForLog("Error, could not marshal user.", err)
		return err
	}

	err = db.client.Set(db.Ctx, userPrefix+strconv.FormatInt(int64(user.ID), 10), string(j), 0).Err()
	if err != nil {
		logger.ForLog("Error, could not save user to the DBRedis.", err)
		return err
	}

	logger.ForLog("User have successfully saved to DBRedis, ID", user.ID)
	return err
}

// GetUser return user form database.
func (db *RedisDB) GetUser(userID int64) (*betypes.User, error) {
	logger.ForLog("Getting a user from a DBRedis, ID", userID)
	u := &betypes.User{}
	r, err := db.client.Get(db.Ctx, userPrefix+strconv.FormatInt(userID, 10)).Result()
	if err == redis.Nil {
		logger.ForLog("User not found.", err)
		return nil, err
	} else if err != nil {
		logger.ForLog("Error, could not read user from DBRedis, ID", userID)
		return nil, err
	}

	err = json.Unmarshal([]byte(r), &u)
	if err != nil {
		logger.ForLog("Error, could not unmarshal user.")
		return nil, err
	}

	logger.ForLog("User successfully received from DBRedis.")
	return u, nil
}
