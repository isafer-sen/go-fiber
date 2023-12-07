package db

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB

// RDB redis connector
var RDB *redis.Client
