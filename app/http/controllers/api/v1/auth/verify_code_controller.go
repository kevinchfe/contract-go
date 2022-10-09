package auth

import "C"
import (
	v1 "contract/app/http/controllers/api/v1"
	"contract/pkg/captcha"
	"contract/pkg/logger"
	"contract/pkg/response"
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
