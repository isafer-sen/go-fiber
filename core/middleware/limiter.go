package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func Limiter() fiber.Handler {
	// 创建一个限速器，限制每秒最多处理10个请求，过期时间为10秒
	return limiter.New(limiter.Config{
		Max:        10,
		Expiration: 10 * time.Second,
	})
}
