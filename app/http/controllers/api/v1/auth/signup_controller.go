package auth

import (
	v1 "contract/app/http/controllers/api/v1"
	"contract/app/models/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}

	// 解析json请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败 返回422和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
