package main

import (
	"log"
	"os"

	"github.com/codepnw/auth-redis-postgres/internal/config"
	"github.com/codepnw/auth-redis-postgres/internal/database"
	"github.com/codepnw/auth-redis-postgres/router"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("loading .env failed")
	}

	dbURL := os.Getenv("DB_URL")
	appPort := os.Getenv("APP_PORT")

	conn := database.ConnectDatabase(dbURL)

	apiConfig := &config.ApiConfig{
		DB:          database.New(conn),
		RedisClient: redisClient,
	}

	router.RegisterRoutes(apiConfig)
	router.Start(appPort)
}
