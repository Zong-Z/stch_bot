package database

import (
	"context"
	"telegram-chat_bot/betypes"

	"github.com/go-redis/redis/v8"
)

type database interface {
	SaveUser(user betypes.User) error
	GetUser(userID int) (*betypes.User, error)
}

// RedisDB saves redis database settings.
type RedisDB struct {
	database
	client *redis.Client
	Ctx    context.Context
}

// DB your database.
var DB *RedisDB
