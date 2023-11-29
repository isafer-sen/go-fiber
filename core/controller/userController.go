package controller

import (
	"app/core/service"
	"app/helper/md5"
	"app/helper/response"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

type UserController struct {
	userService service.UserService
}

func (e *UserController) Index(c *fiber.Ctx) error {
	page := c.QueryInt("page", 0)
	pageSize := c.QueryInt("pageSize", 10)
	data, _ := e.userService.FindAll(page, pageSize)
	return response.Success(c, data)
}

func (e *UserController) Cache(c *fiber.Ctx) error {
	page := c.QueryInt("page", 0)
	pageSize := c.QueryInt("pageSize", 10)
	md5Query := md5.Sum("user_index#" + cast.ToString(page) + cast.ToString(pageSize))
	data, err := e.userService.FindAllCache(page, pageSize, md5Query)
	if err != nil {
		return response.Failed(c, err.Error())
	}
	return response.Success(c, data)
}

func (e *UserController) Json(c *fiber.Ctx) error {
	return c.Redirect("https://www.baidu.com", 301)
}
