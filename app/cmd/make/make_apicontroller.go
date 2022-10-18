package make

import (
	"contract/pkg/console"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller,example: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1), // 只允许且必须传一个参数
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}
	// apiVersion 用来拼接目标路劲
	// name 用来生成cmd.Model 实例
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)
	// 组建目标目录
	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, model.TableName)
	createFileFromStub(filePath, "apicontroller", model)
}
