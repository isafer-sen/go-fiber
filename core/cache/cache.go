package cache

import (
	"app/core/database"
	"context"
	"time"
)

func Set(key string, value any, exp time.Duration) error {
	return database.RDB.Set(context.Background(), key, value, exp).Err()
}

func Get(key string) (result string, err error) {
	return database.RDB.Get(context.Background(), key).Result()
}

func Del(key string) error {
	return database.RDB.Del(context.Background(), key).Err()
}

func HasCache(key string) bool {
	return database.RDB.Exists(context.Background(), key).Val() > 0
}
