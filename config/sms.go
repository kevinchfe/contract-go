package config

import "contract/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认阿里云
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ALIYUN_ACCESS_ID"),
				"access_key_secret": config.Env("SMS_ALIYUN_ACCESS_SECRET"),
				"sign_name":         config.Env("SMS_ALIYUN_SIGN_NAME", "阿里云短信测试"),
				"template_code":     config.Env("SMS_ALIYUN_TEMPLATE_CODE", "SMS_154950909"),
			},
			// 七牛云
			"qiniu": map[string]interface{}{
				"access_key":    config.Env("SMS_QINIUYUN_ACCESS_ID"),
				"access_secret": config.Env("SMS_QINIUYUN_ACCESS_SECRET"),
				"signature_id":  config.Env("SMS_QINIUYUN_SIGNATUREID"),
				"template_code": config.Env("SMS_QINIUYUN_TEMPLATE_CODE"),
			},
		}
	})
}
