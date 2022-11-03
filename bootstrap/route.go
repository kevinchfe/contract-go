// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"contract/app/http/middlewares"
	"contract/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// SetupRoute 初始化路由
func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册小程序路由
	routes.RegisterAppletsRoutes(router)

	// 配置404路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
		//gin.Logger(),
		//gin.Recovery(),
	)
}

func setup404Handler(routr *gin.Engine) {
	routr.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusOK, "页面返回404")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error_code":    404,
				"error_message": "路由未定义",
			})
		}
	})
}
