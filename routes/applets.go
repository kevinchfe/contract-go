package routes

import (
	"contract/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAppletsRoutes 注册小程序路由
func RegisterAppletsRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否被注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
		}

		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "v1",
			})
		})
	}
}
