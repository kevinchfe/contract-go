package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	Sort    string `valid:"sort" form:"sort"`
	Order   string `valid:"order" form:"order"`
	PerPage int    `valid:"per_page" form:"per_page"`
}

func Pagination(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"sort":     []string{"in:id,created_at,updated_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
	}
	messages := govalidator.MapData{
		"sort": []string{
			"in:排序规则仅支持id,created_at,updated_at",
		},
		"order": []string{
			"in:排序规则仅支持asc,desc",
		},
		"per_page": []string{
			"numeric_between:每页条数介于2-100之间",
		},
	}
	return validate(data, rules, messages)
}
