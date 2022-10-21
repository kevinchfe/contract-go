package app

import (
	"contract/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}

// URL 传参path拼接站点url
func URL(path string) string {
	return config.Get("app.url") + path
}

// V1URL 拼接带v1标示url
func V1URL(path string) string {
	return URL("/v1/" + path)
}
