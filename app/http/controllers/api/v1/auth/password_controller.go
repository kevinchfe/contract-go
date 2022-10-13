package auth

import (
	v1 "contract/app/http/controllers/api/v1"
	"contract/app/models/user"
	"contract/app/requests"
	"contract/pkg/response"
	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseApiController
}

// ResetByPhone 手机重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}
	// 更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}

// ResetByEmail 邮箱充值密码
func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
