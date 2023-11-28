package database

import (
	"app/config"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func SetupDB() {
	// 使用 gorm.Open 连接数据库
	var err error
	var SQL *sql.DB
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: config.DriverName,
		DSN:        config.DSN,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// 处理错误
	if err != nil {
		log.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SQL, err = DB.DB()
	// 设置最大连接数
	SQL.SetMaxOpenConns(config.MaxOpenConnections)
	// 设置最大空闲连接数
	SQL.SetMaxIdleConns(config.MaxIdleConnections)
	// 设置每个链接的过期时间
	SQL.SetConnMaxLifetime(time.Duration(config.MaxLifeSeconds) * time.Second)
	if err != nil {
		log.Println(err.Error())
	}
}
