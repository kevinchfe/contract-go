package user

import (
	"contract/app/models"
	"contract/pkg/database"
	"contract/pkg/hash"
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
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}
