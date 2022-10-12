package middlewares

import (
	"contract/app/models/user"
	"contract/pkg/config"
	"contract/pkg/jwt"
	"contract/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取jwt 并验证
		claims, err := jwt.NewJWT().ParserToken(c)
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关接口认证文档", config.GetString("app.name")))
			return
		}

		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户")
			return
		}

		// 将用户信息存入 gin.Context 后续auth包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)
		c.Next()
	}
}
