// Package link 模型
package link

import (
	"contract/app/models"
	"contract/pkg/database"
)

type Link struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`

	models.CommonTimestampsField
}

func (links *Link) Create() {
	database.DB.Create(&links)
}

func (links *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&links)
	return result.RowsAffected
}

func (links *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&links)
	return result.RowsAffected
}
