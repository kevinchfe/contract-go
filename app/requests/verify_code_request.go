package requests

import (
	"contract/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	Phone         string `json:"phone,omitempty" valid:"phone"`
}

type VerifyCodeEmailRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	Email         string `json:"email,omitempty" valid:"email"`
}

// VerifyCodePhone 验证表单
func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号必填",
			"digits:手机号长度必须为11",
		},
		"captcha_id": []string{
			"required:图片验证码id必填",
		},
		"captcha_answer": []string{
			"required:答案必填",
			"digits:答案必须为6位数字",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*VerifyCodePhoneRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
	return errs
}

// VerifyCodeEmail 验证表单，长度为0即通过
func VerifyCodeEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":          []string{"required", "min:4", "max:30", "email"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填",
			"min:Email 长度必须大于4",
			"max:Email 长度必须小于30",
			"email:Email 格式不正确",
		},
		"captcha_id": []string{
			"required: 图片验证码id必填",
		},
		"captcha_answer": []string{
			"required: 答案必填",
			"digits: 长度必须为6位的数字",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*VerifyCodeEmailRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
	return errs
}
