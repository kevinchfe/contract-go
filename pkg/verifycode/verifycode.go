// Package verifycode 用以发送手机验证码和邮箱验证码
package verifycode

import (
	"contract/pkg/app"
	"contract/pkg/config"
	"contract/pkg/helpers"
	"contract/pkg/logger"
	"contract/pkg/mail"
	"contract/pkg/redis"
	"contract/pkg/sms"
	"fmt"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

// NewVerifyCode 单例获取
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				// 增加前缀
				KeyPrefix: config.GetString("app.name") + ":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

// SendSMS 发送短信验证码 verifycode.NewVerifyCode.SendSMS(request.Phone)
func (vc *VerifyCode) SendSMS(phone string) bool {
	// 生成验证码
	code := vc.generateVerifyCode(phone)
	if !app.IsProduction() && strings.HasSuffix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}
	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.qiniu.template_code"),
		Data:     map[string]interface{}{"code": code},
	})
}

// SendEmail 发送邮件验证码 verifycode.NewVerifyCode().SendEmail(request.Email)
func (vc *VerifyCode) SendEmail(email string) error {
	code := vc.generateVerifyCode(email)
	if !app.IsProduction() && strings.HasSuffix(email, config.GetString("verifycode.debug_email_suffix")) {
		return nil
	}
	content := fmt.Sprintf("<h1> 您的Email验证码是%v <h1>", code)
	mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: config.GetString("mail.from.address"),
			Name:    config.GetString("mail.from.name"),
		},
		To:      []string{email},
		Subject: "Email 验证码",
		HTML:    []byte(content),
	})
	return nil
}

// CheckAnswer 检查用户提交的验证码是否正确，key 可以是手机号或者 Email
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {
	logger.DebugJSON("验证码", "检查验证", map[string]string{key: answer})
	if !app.IsProduction() && strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) ||
		strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}
	return vc.Store.Verify(key, answer, false)
}

// generateVerifyCode 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {
	// 生成随机数
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}
	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})
	vc.Store.Set(key, code)
	return code
}
