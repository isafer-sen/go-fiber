package router

import (
	"app/core/controller"
	"app/core/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(app *fiber.App) {
	//app.Use(logger.New())
	u := new(controller.UserController)
	app.Get("/index", u.Index)
	app.Get("/cache", u.Cache)
	app.Get("/json", u.Json)

	//app.Get("/ws/:id", service.WsConnect())
	app.Get("/ws/:id", service.WsConnect())

	wsc := new(controller.WsController)
	app.Get("/wss/send", wsc.SendMsg)

	// 设备管理
	dev := new(controller.DeviceController)
	app.Get("/device", dev.Index)

	// 上传文件
	up := new(controller.UploadController)
	app.Post("/upload", up.Upload)
}
