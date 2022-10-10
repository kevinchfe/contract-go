package user

import (
	"contract/app/models"
	"contract/pkg/database"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户 通过User.ID判断是否成功
func (u *User) Create() {
	database.DB.Create(&u)
}
