package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func JSON(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusOK).JSON(data)
}

func Success(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"code": 0, "msg": "success", "data": data})
}

func Failed(c *fiber.Ctx, msg any) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"code": 1, "msg": msg})
}
