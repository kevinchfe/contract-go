package middlewares

import (
	"contract/pkg/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取user-agent标头信息
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"), "请求必须附带 User-Agent 标头")
			return
		}
		c.Next()
	}
}
