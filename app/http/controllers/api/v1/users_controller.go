package v1

import (
	"contract/app/models/user"
	"contract/pkg/auth"
	"contract/pkg/response"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseApiController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// Index 全部用户
func (ctrl *UsersController) Index(c *gin.Context) {
	data := user.All()
	response.Data(c, data)
}
