package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterApiRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "v1",
			})
		})
	}
}
