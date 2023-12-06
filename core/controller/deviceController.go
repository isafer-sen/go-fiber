package controller

import (
	"app/core/service"
	"app/helper/response"
	"github.com/gofiber/fiber/v2"
)

type DeviceController struct {
	deviceService service.DeviceService
}

func (e *DeviceController) Index(c *fiber.Ctx) error {
	data, _ := e.deviceService.FindAll()
	return response.Success(c, data)
}
