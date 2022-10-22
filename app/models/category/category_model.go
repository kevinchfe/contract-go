// Package category 模型
package category

import (
	"contract/app/models"
	"contract/pkg/database"
)

type Category struct {
	models.BaseModel
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	models.CommonTimestampsField
}

func (categories *Category) create() {
	database.DB.Create(&categories)
}

func (categories *Category) Save() (rowsAffected int64) {
	result := database.DB.Save(&categories)
	return result.RowsAffected
}

func (categories *Category) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&categories)
	return result.RowsAffected
}
