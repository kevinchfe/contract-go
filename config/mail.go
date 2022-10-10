package config

import "contract/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{
			"smtp": map[string]interface{}{
				"host":     config.Env("MAIL_HOST", "localhost"),
				"port":     config.Env("MAIL_PORT"),
				"username": config.Env("MAIL_USERNAME", ""),
				"password": config.Env("MAIL_PASSWORD", ""),
			},
			"from": map[string]interface{}{
				"address": config.Env("MAIL_FROM_ADDRESS", "546598819@qq.com"),
				"name":    config.Env("MAIL_FROM_NAME", "kevin"),
			},
		}
	})
}
