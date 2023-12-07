package cache

import (
	"app/db"
	"context"
	"time"
)

func Set(key string, value any, exp time.Duration) error {
	return db.RDB.Set(context.Background(), key, value, exp).Err()
}

func Get(key string) (result string, err error) {
	return db.RDB.Get(context.Background(), key).Result()
}

func Del(key string) error {
	return db.RDB.Del(context.Background(), key).Err()
}

func HasCache(key string) bool {
	return db.RDB.Exists(context.Background(), key).Val() > 0
}
