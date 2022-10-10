package auth

import "C"
import (
	v1 "contract/app/http/controllers/api/v1"
	"contract/app/requests"
	"contract/pkg/captcha"
	"contract/pkg/logger"
	"contract/pkg/response"
	"contract/pkg/verifycode"
	"github.com/gin-gonic/gin"
)

// VerifyCodeController 验证码控制器
type VerifyCodeController struct {
	v1.BaseApiController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码图片
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志
	logger.LogIf(err)
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// SendUsingPhone 发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {
	// 1. 表单验证
	request := requests.VerifyCodePhoneRequest{}
	logger.DebugString("测试", "1244", string(request.Phone))
	//logger.Info("测试", request.Phone)
	//fmt.Println(request.CaptchAnswer)
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// 2. 发送SMS
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败")
	} else {
		response.Success(c)
	}
}
