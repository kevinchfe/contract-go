// Package topic 模型
package topic

import (
	"contract/app/models"
	"contract/app/models/category"
	"contract/app/models/user"
	"contract/pkg/database"
)

type Topic struct {
	models.BaseModel
	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	// 通过user_id关联用户
	User user.User `json:"user"`
	// 通过category_id关联分类
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topics *Topic) Create() {
	database.DB.Create(&topics)
}

func (topics *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topics)
	return result.RowsAffected
}

func (topics *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topics)
	return result.RowsAffected
}
