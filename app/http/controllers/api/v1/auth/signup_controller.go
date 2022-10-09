// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "contract/app/http/controllers/api/v1"
	"contract/app/models/user"
	"contract/app/requests"
	"contract/pkg/response"
	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseApiController
}

// IsPhoneExist 检查手机是否注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数 并做验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	// 初始化请求对象
	/*request := requests.SignupPhoneExistRequest{}
	if ok :=  {

	}

	// 解析json请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败 返回422和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		fmt.Println(err.Error())
		return
	}

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}*/

	// 检查数据库并返回响应
	//response.JSON(c, user.IsPhoneExist(request.Phone))

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检查邮箱是否注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 获取请求参数 并做验证
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
