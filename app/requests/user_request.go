package requests

import (
	"contract/app/requests/validators"
	"contract/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"mime/multipart"
)

type UserUpdateProfileRequest struct {
	Name         string `valid:"name" json:"name"`
	City         string `valid:"city" json:"city"`
	Introduction string `valid:"introduction" json:"introduction"`
}

type UserUpdateEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

type UserUpdatePhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {
	// 查询用户名重复时，过滤掉当前用户id
	uid := auth.CurrentUID(c)

	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_num", "between:3,20", "not_exists:users,name," + uid},
		"introduction": []string{"min_cn:4", "max_cn:240"},
		"city":         []string{"min_cn:2", "max_cn:20"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度为3-20个字符",
			"not_exists:名称已存在",
		},
		"introduction": []string{
			"min_cn:描述至少4个字",
			"max_cn:描述最多不能超过240个字",
		},
		"city": []string{
			"min_cn:城市至少2个字",
			"max_cn:城市最多不能超过20个字",
		},
	}
	return validate(data, rules, messages)
}

func UserUpdateEmail(data interface{}, c *gin.Context) map[string][]string {
	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"email": []string{
			"required", "min:4", "max:30", "email", "not_exists:users,email" + currentUser.GetStringID(), "not_in:" + currentUser.Email,
		},
		"verify_code": []string{
			"required", "digits:6",
		},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email必填",
			"min:长度大于4",
			"max:长度必须小于30",
			"email:格式不正常",
			"not_exists:邮箱已被占用",
			"not_in:新旧邮箱一致",
		},
		"verify_code": []string{
			"required:验证码必须",
			"digits:验证码必须为6位数字",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*UserUpdateEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)
	return errs
}

func UserUpdatePhone(data interface{}, c *gin.Context) map[string][]string {

	currentUser := auth.CurrentUser(c)

	rules := govalidator.MapData{
		"phone": []string{
			"required",
			"digits:11",
			"not_exists:users,phone," + currentUser.GetStringID(),
			"not_in:" + currentUser.Phone,
		},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:手机号已被占用",
			"not_in:新的手机与老手机号一致",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

type UserUpdatePasswordRequest struct {
	Password           string `json:"password,omitempty" valid:"password"`
	NewPassword        string `json:"new_password,omitempty" valid:"new_password"`
	NewPasswordConfirm string `json:"new_password_confirm,omitempty" valid:"new_password_confirm"`
}

type UserUpdateAvatarRequest struct {
	Avatar *multipart.FileHeader `from:"avatar" valid:"avatar"`
}

func UserUpdatePassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"password":             []string{"required", "min:6"},
		"new_password":         []string{"required", "min:6"},
		"new_password_confirm": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"new_password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"new_password_confirm": []string{
			"required:确认密码框为必填项",
			"min:确认密码长度需大于 6",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePasswordRequest)
	errs = validators.ValidatePasswordConfirm(_data.NewPassword, _data.NewPasswordConfirm, errs)
	return errs
}

func UserUpdateAvatar(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"file:avatar": []string{"ext:png,jpg,jpeg", "size:20971520", "required"},
	}
	messages := govalidator.MapData{
		"file:avatar": []string{
			"required:图片必须",
			"size:文件大小不能超过20M",
			"ext:必须为png,jpg,jpeg格式",
		},
	}
	return ValidateFile(c, data, rules, messages)
}
