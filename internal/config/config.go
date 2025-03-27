package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/codepnw/auth-redis-postgres/internal/database"
)

type ApiConfig struct {
	DB          *database.Queries
	RedisClient *redis.Client
}
