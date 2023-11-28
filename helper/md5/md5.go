package md5

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"sort"
	"strings"
)

func Sum(val string) string {
	hash := md5.Sum([]byte(val))
	return hex.EncodeToString(hash[:])
}

func SumQuery(c *fiber.Ctx, prefix string) string {
	params := c.AllParams()
	// 将所有请求参数按照 key 进行排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 将所有请求参数按照 key=value 拼接成一个字符串
	var combinedParams string
	for _, k := range keys {
		combinedParams += prefix + k + "=" + params[k] + "&"
	}
	combinedParams = strings.TrimRight(combinedParams, "&")
	log.Info(combinedParams)
	// 对组合后的参数进行MD5加密
	return Sum(combinedParams)
}
