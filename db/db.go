package database

import (
	"context"
	"telegram-chat_bot/betypes"

	"github.com/go-redis/redis/v8"
)

// IDatabase the interface that each database structure must implement.
// Designed for easy and quick replacement of one database to another.
type IDatabase interface {
	SaveUser(user betypes.User) error
	GetUser(userID int) (*betypes.User, error)
}

// RedisDB saves all redis settings.
// Implements IDatabase.
type RedisDB struct {
	IDatabase
	client *redis.Client
	Ctx    context.Context
}

// DB your database.
var DB *RedisDB
