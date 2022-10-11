// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "contract/app/http/controllers/api/v1"
	"contract/app/models/user"
	"contract/app/requests"
	"contract/pkg/jwt"
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

// SignupUsingPhone 手机号注册
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	// 验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}
	// 验证成功，创建数据
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		// 返回token
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  _user,
		})
	} else {
		response.Abort500(c, "创建用户失败")
	}
}

// SignupUsingEmail 邮箱注册用户
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	// 表单验证
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}
	// 验证成功，创建用户
	_user := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  _user,
		})
	} else {
		response.Abort500(c, "创建用户失败")
	}
}
