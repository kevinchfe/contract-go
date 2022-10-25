package policies

import (
	"contract/app/models/topic"
	"contract/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, topicsModel topic.Topic) bool {
	return auth.CurrentUID(c) == topicsModel.UserID
}

// func CanViewTopic(c *gin.Context, topicsModel topic.Topic) bool {}
// func CanCreateTopic(c *gin.Context, topicsModel topic.Topic) bool {}
// func CanUpdateTopic(c *gin.Context, topicsModel topic.Topic) bool {}
// func CanDeleteTopic(c *gin.Context, topicsModel topic.Topic) bool {}
