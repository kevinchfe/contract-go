package requests

import (
	"contract/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// SignupUsingPhoneRequest 通过手机注册请求信息
type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `json:"name,omitempty" valid:"name"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

type SignupUsingEmailRequest struct {
	Email           string `json:"email,omitempty" valid:"email"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `json:"name,omitempty" valid:"name"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
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

// SignupUsingPhone 通过手机号注册
func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号必填",
			"digits:手机号长度必须为11位数字",
		},
		"name": []string{
			"required:用户名必填",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度必须在3-20位之间",
		},
		//"password": []string{
		//	"required:密码必须",
		//	"min:密码长度必须大于6位",
		//},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码必须",
		},
		"verify_code": []string{
			"required:验证码必须",
			"digits:验证码长度必须为6位",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}

// SignupUsingEmail 邮箱注册
func SignupUsingEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email必填",
			"min:Email长度必须大于4",
			"max:Email长度必填小于30",
			"email:Email格式错误",
			"not_exists:Email已被占用",
		},
		"name": []string{
			"required:用户名必须",
			"alpha_num:用户名只能为数组和字母",
			"between:用户名长度在3-20位之间",
			"not_exists:用户名已被占用",
		},
		"password": []string{
			"required:密码必须",
			"min:密码至少6位",
		},
		"password_confirm": []string{
			"required:确认密码必须",
		},
		"verify_code": []string{
			"required:验证码必须",
			"digits:验证码长度位6位数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*SignupUsingEmailRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)
	return errs
}
