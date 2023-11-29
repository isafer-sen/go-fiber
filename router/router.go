package router

import (
	"app/core/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(app *fiber.App) {
	//app.Use(logger.New())
	u := new(controller.UserController)
	app.Get("/index", u.Index)
	app.Get("/cache", u.Cache)
	app.Get("/json", u.Json)

	app.Get("/ws/:id", controller.WS())
}
