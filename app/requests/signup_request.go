package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机必填",
			"digits:手机号长度为11位数字",
		},
	}

	return validate(data, rules, messages)
}

func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email必填",
			"min:Email长度需大于4",
			"max:EmaiL长度需小于30",
			"email:Email格式不正确",
		},
	}

	return validate(data, rules, messages)
}
