package make

import (
	"contract/pkg/console"
	"contract/pkg/file"
	"contract/pkg/str"
	"embed"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"strings"
)

// Model 参数解释
// 单个词，用户命令传参。以User模型为例
// - user
// - User
// - users
// - Users
// 整理好的数据
//{
//		"TableName": "Users",
//		"Struct": "User",
// 		"StructNamePlural": "Users",
//      "VariableName": "user",
//      "VariableNamePlural": "users,
//      "PackageName": "user"
//}
// -
// 两个词或以上，用户命令传参，以TopicComment 模型为例
// - topic_comment
// - topic_comments
// - TopicComment
// -TopicComments
// 整理好的数据
//{
//     "TableName": "topic_comments",
//     "Struct": "TopicComment",
//     "StructNamePlural": "TopicComments",
//     "VariableName": "topicComment",
//     "VariableNamePlural": "topicComments",
//     "PackageName": "topic_comment"
//}

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

// stubsFS 方便我们后面打包这些 .stub 为后缀的文件

//go:embed stubs
var stubsFS embed.FS

// CmdMake 说明cobra命令
var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	// 注册make子命令
	CmdMake.AddCommand(
		CmdMakeCMD,
		CmdMakeModel,
	)
}

// makeModelFromString 格式化用户输入的内容
func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = str.Singular(strcase.ToCamel(name))
	model.StructNamePlural = str.Plural(model.StructName)
	model.TableName = str.Snake(model.StructNamePlural)
	model.VariableName = str.LowerCamel(model.TableName)
	model.PackageName = str.Snake(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)
	return model
}

// createFileFromStub 读取stub文件并进行变量替换
// 最后一个选项可选，如若传参，应传map[string]string 类型，作为附件的变量搜索替换
func createFileFromStub(filePath string, stubName string, model Model, valiables ...interface{}) {
	// 实现最后一个参数可选
	replaces := make(map[string]string)
	if len(valiables) > 0 {
		replaces = valiables[0].(map[string]string)
	}
	// 目标文件已存在
	if file.Exists(filePath) {
		console.Exit(filePath + " already exists")
	}
	// 读取stub模板文件
	modelData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}
	modelStub := string(modelData)

	// 添加默认的替换变量
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{TableName}}"] = model.TableName

	// 对模板内容做替换
	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}

	// 存储到目标文件
	err = file.Put([]byte(modelStub), filePath)
	if err != nil {
		console.Exit(err.Error())
	}
	console.Success(fmt.Sprintf("[%s] created.", filePath))
}
