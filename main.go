package main

import (
	"contract/app/cmd"
	"contract/bootstrap"
	config2 "contract/config"
	"contract/pkg/config"
	"contract/pkg/console"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	// 加载 config 目录下配置信息
	config2.Initialize()
}

// go2 api实战
func main() {
	// 应用主入口，默认调用cmd.CmdServe命令
	var rootCmd = &cobra.Command{
		Use:   "Contract",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		// rootCmd 的所有子命令都会执行一下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)
			// 初始化logger
			bootstrap.SetupLogger()
			// 初始化DB
			bootstrap.SetupDB()
			// 初始化redis
			bootstrap.SetupRedis()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
	)

	// 配置默认运行的web服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数 --env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

	// 配置初始化
	//var env string
	//flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=dev 加载的是 .env.dev 文件")
	//flag.Parse()
	//config.InitConfig(env)
	//
	//// 初始化logger
	//bootstrap.SetupLogger()
	//
	//// 设置 gin 的运行模式，支持 debug, release, test
	//// release 会屏蔽调试信息，官方建议生产环境中使用
	//// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	//// 故此设置为 release，有特殊情况手动改为 debug 即可
	//gin.SetMode(gin.ReleaseMode)
	//
	//// 初始化DB
	//bootstrap.SetupDB()
	//
	//// 初始化Redis
	//bootstrap.SetupRedis()
	//
	//// new 一个gin Engine实例
	//r := gin.New()
	//bootstrap.SetupRoute(r)
	//
	////logger.Dump(captcha.NewCaptcha().VerifyCaptcha("aAeKkhd2MNEC0oNOXYjc", "993223"), "正确答案")
	////logger.Dump(captcha.NewCaptcha().VerifyCaptcha("aAeKkhd2MNEC0oNOXYjc", "111111"), "错误答案")
	//err := r.Run(":" + config.Get("app.port"))
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
}
