package config

import "contract/pkg/config"

func init() {
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{
			"code_length": config.Env("VERIFY_CODE_LENGTH", 6),
			// 分钟
			"expire_time":        config.Env("VERIFY_CODE_EXPIRE", 15),
			"debug_expire_time":  10080,
			"debug_code":         123456,
			"debug_phone_prefix": "000",
			"debug_email_suffix": "@testing.com",
		}
	})
}
