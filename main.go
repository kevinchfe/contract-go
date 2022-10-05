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

func main() {
	// 配置初始化
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=dev 加载的是 .env.dev 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化logger
	bootstrap.SetupLogger()

	// 初始化DB
	bootstrap.SetupDB()

	// new 一个gin Engine实例
	r := gin.New()
	bootstrap.SetupRoute(r)
	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Printf(err.Error())
	}
}
