package router

import (
	"app/core/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func RegisterRouter(app *fiber.App) {
	//app.Use(logger.New())
	u := new(controller.UserController)
	app.Get("/index", u.Index)
	app.Get("/cache", cache.New(), u.Cache)
}
