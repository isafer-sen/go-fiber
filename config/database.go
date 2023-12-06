package config

const (
	DriverName         = "mysql"
	DSN                = "root:root@tcp(localhost:3306)/wlw_2?charset=utf8&parseTime=True&loc=Local"
	MaxOpenConnections = 300 // 设置最大连接数
	MaxIdleConnections = 50  // 设置最大空闲连接数
	MaxLifeSeconds     = 500 // 设置每个链接的过期时间
)
