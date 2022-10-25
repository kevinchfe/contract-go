package v1

import (
	"contract/app/models/topic"
	"contract/app/policies"

	"contract/app/requests"
	"contract/pkg/auth"
	"contract/pkg/response"
	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseApiController
}

func (ctrl *TopicsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := topic.Paginate(c, 10)
	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicsModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicsModel.Create()
	if topicsModel.ID > 0 {
		response.Created(c, topicsModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Update(c *gin.Context) {

	topicsModel := topic.Get(c.Param("id"))
	if topicsModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicsModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicsModel.Title = request.Title
	topicsModel.Body = request.Body
	topicsModel.CategoryID = request.CategoryID
	rowsAffected := topicsModel.Save()
	if rowsAffected > 0 {
		response.Data(c, topicsModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Delete(c *gin.Context) {

	topicsModel := topic.Get(c.Param("id"))
	if topicsModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicsModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := topicsModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
