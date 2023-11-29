package controller

import (
	"app/core/service"
	"github.com/gofiber/fiber/v2"
)

type WsController struct {
}

func (e *WsController) SendMsg(c *fiber.Ctx) error {
	return service.SendMsg("你好")
}
