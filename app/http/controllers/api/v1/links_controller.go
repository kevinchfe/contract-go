package v1

import (
	"contract/app/models/link"
	"contract/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseApiController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	response.Data(c, link.AllCached())
}
