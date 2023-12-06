package controller

import (
	"app/helper/fastdfs"
	"app/helper/response"
	"github.com/gofiber/fiber/v2"
)

type UploadController struct {
}

// Upload 上传文件
func (e *UploadController) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	err, res := fastdfs.Upload(file)
	if err != nil {
		return err
	}
	// 返回文件上传成功的响应
	return response.Success(c, res)
}
