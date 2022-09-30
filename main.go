package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "1242",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		accetpString := c.Request.Header.Get("Accept")
		if strings.Contains(accetpString, "text/html") {
			c.String(http.StatusOK, "页面返回404")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error_code":    404,
				"error_message": "路由未定义",
			})
		}
	})

	r.Run(":8080")
}
