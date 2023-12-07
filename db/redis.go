package db

import "github.com/go-redis/redis/v8"

func SetupRedis() {
	// Initialize custom config
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
}
