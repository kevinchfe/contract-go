package config

import "contract/pkg/config"

func init() {
	config.Add("captcha", func() map[string]interface{} {
		return map[string]interface{}{
			// 验证码图片高度
			"height": 80,
			// 验证码图片宽度
			"width": 240,
			// 验证码长度
			"length": 6,
			// 数字最大倾斜度
			"maxskew": 0.7,
			// 图片背景里的混淆点数量
			"dotcount": 80,
			// 过期时间,单位分钟
			"expire_time": 15,
			// debug模式下的过期时间，方便本地开发调试
			"debug_expire_time": 10080,
			// 非production环境，使用次key可跳过验证，方便调试
			"testing_key": "captcha_skip_test",
		}
	})
}
