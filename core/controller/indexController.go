package controller

import "github.com/gofiber/fiber/v2"

type IndexController struct {
}

func (e IndexController) Index(c *fiber.Ctx) error {
	return c.SendString("string")
}
