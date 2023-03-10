package routes

import (
	controller "contract/app/http/controllers/api/v1"
	"contract/app/http/controllers/api/v1/auth"
	"contract/app/http/middlewares"
	"contract/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAppletsRoutes 注册小程序路由
func RegisterAppletsRoutes(r *gin.Engine) {

	var v1 *gin.RouterGroup
	if len(config.GetString("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}

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
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("1000-H"), vcc.SendUsingPhone)
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
		usersGroup := v1.Group("/users")
		usersGroup.Use(middlewares.AuthJWT())
		{
			usersGroup.GET("", uc.Index)
			usersGroup.PUT("", uc.UpdateProfile)
			usersGroup.PUT("/email", uc.UpdateEmail)
			usersGroup.PUT("/phone", uc.UpdatePhone)
			usersGroup.PUT("/password", uc.UpdatePassword)
			usersGroup.PUT("/avatar", uc.UpdateAvatar)
		}

		cgc := new(controller.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			cgcGroup.GET("", cgc.Index)
			cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
			cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
		}

		tpc := new(controller.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			tpcGroup.GET("", tpc.Index)
			tpcGroup.GET("/:id", tpc.Show)
			tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)
			tpcGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
			tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
		}

		lsc := new(controller.LinksController)
		lscGroup := v1.Group("/links")
		{
			lscGroup.GET("", lsc.Index)
		}

		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "v1",
			})
		})
	}
}
