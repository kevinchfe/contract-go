package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create policy file, example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1), // 有且只有一个参数
}

func runMakePolicy(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	// MkdirAll确保父目录和子目录都创建
	os.MkdirAll("app/policies", 0777)
	filePath := fmt.Sprintf("app/policies/%s_policy.go", model.PackageName)
	createFileFromStub(filePath, "policy", model)
}
