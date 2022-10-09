// Package requests 处理请求数据和表单验证
package requests

import (
	"contract/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// 1. 解析请求，支持json数据 表单和url query
	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")

		//c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		//	"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
		//	"error":   err.Error(),
		//})
		//fmt.Println(err.Error())
		return false
	}
	// 2. 表单验证
	errs := handler(obj, c)
	// 3. 判断验证是否通过
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		//c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		//	"message": "请求验证不通过, 具体请查看errors",
		//	"errors":  errs,
		//})
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的struct标签标识符
		Messages:      messages,
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
