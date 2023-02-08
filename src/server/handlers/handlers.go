package handlers

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
)

// TODO: Add comment documentation (type Handler)
type Handler struct {
	DB     *gorm.DB
	client *redis.Client
}

// TODO: Add comment documentation (func NewHandler)
func NewHandler(db *gorm.DB, redis *redis.Client) Handler {
	return Handler{db, redis}
}

func Init() *redis.Client {
	// Initialize Redis
	dsn := "localhost:6379"
	client := redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}

	return client
}
