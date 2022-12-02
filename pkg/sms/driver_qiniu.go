package sms

import (
	config2 "contract/pkg/config"
	"contract/pkg/logger"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/sms"
)

// Qiniu 实现 sms.Driver interface
type Qiniu struct {
}

func (s *Qiniu) Send(phone string, message Message, config map[string]string) bool {
	logger.DebugJSON("短信[七牛]", "配置信息", config)

	mac := auth.New(config["access_key"], config["access_secret"])
	manager := sms.NewManager(mac)
	result, err := manager.SendMessage(sms.MessagesRequest{
		SignatureID: config2.GetString("sms.qiniu.signature_id"),
		TemplateID:  message.Template,
		Mobiles:     []string{phone},
		Parameters:  message.Data,
	})

	logger.DebugJSON("短信[七牛]", "接口响应", result)
	if err != nil {
		logger.ErrorString("短信[七牛]", "发送失败", err.Error())
		return false
	}
	
	return true
}
