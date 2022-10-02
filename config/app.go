package config

import "contract/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": config.Env("APP_NAME", "contract"),
			// 当前环境 local, dev,yfb,production
			"env": config.Env("APP_ENV", "production"),
			// 是否调试模式
			"debug": config.Env("APP_DEBUG", false),
			// 服务端口
			"port": config.Env("APP_PORT", "8080"),
			// 加密会话 JWT加密
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
			// 用以生成链接
			"url": config.Env("APP_URL", "http://localhost:8080"),
			// 设置时区
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
