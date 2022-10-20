package routes

import (
	controller "contract/app/http/controllers/api/v1"
	"contract/app/http/controllers/api/v1/auth"
	"contract/app/http/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAppletsRoutes 注册小程序路由
func RegisterAppletsRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流 这里是所有api(ip)请求加起来
	// 同一个ip访问v1所有的api 每小时最多200个
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 判断手机是否被注册
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)
			// 限流中间件：每小时限流 这里是对单个路由请求限制
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.SignupUsingEmail)

			// 验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("20-H"), vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			// 登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
		}

		uc := new(controller.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)

		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "v1",
			})
		})
	}
}
