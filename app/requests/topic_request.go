package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	Title      string `json:"title,omitempty" valid:"title"`
	Body       string `json:"body,omitempty" valid:"body"`
	CategoryID string `json:"category_id,omitempty" valid:"category_id"`
}

func TopicSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title": []string{"required", "min_cn:3", "max_cn:40"},
		"body":  []string{"required", "min_cn:10", "max_cn:50000"},
		//"category_id": []string{"required", "exists:categories,id"},
		"category_id": []string{"required"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:话题标题为必填项",
			"min_cn:名称长度需至少 3 个字",
			"max_cn:名称长度不能超过 40 个字",
		},
		"body": []string{
			"required:话题内容必须",
			"min_cn:描述长度需至少 10 个字",
		},
		"category_id": []string{
			"required:帖子分类为必填项",
			"exists:帖子分类未找到",
		},
	}
	return validate(data, rules, messages)
}
