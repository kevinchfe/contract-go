package auth

import (
	"contract/app/models/user"
	"contract/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
)

// Attempt 尝试登录
func Attempt(loginID string, password string) (user.User, error) {
	userModel := user.GetByMulti(loginID)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}
	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}
	return userModel, nil
}

// LoginByPhone 登录
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return userModel, nil
}

// CurrentUser 返回当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	return userModel
}

// CurrentUID 从gin.Context 获取当前登录用户id
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
