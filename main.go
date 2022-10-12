package main

import (
	"contract/bootstrap"
	config2 "contract/config"
	"contract/pkg/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下配置信息
	config2.Initialize()
}

// go2 api实战
func main() {
	// 配置初始化
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=dev 加载的是 .env.dev 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化logger
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// 初始化DB
	bootstrap.SetupDB()

	// 初始化Redis
	bootstrap.SetupRedis()

	// new 一个gin Engine实例
	r := gin.New()
	bootstrap.SetupRoute(r)
	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("aAeKkhd2MNEC0oNOXYjc", "993223"), "正确答案")
	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("aAeKkhd2MNEC0oNOXYjc", "111111"), "错误答案")
	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Printf(err.Error())
	}
}
