package v1

import (
	"contract/app/models/category"
	"contract/app/requests"
	"contract/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseApiController
}

func (ctrl *CategoriesController) Store(c *gin.Context) {
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoriesModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoriesModel.Create()
	if categoriesModel.ID > 0 {
		response.Created(c, categoriesModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}
