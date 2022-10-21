package user

import (
	"contract/pkg/app"
	"contract/pkg/database"
	"contract/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// IsEmailExist 判断email是否被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email=?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号是否被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone=?", phone).Count(&count)
	return count > 0
}

func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone=?", phone).First(&userModel)
	return
}

func GetByEmail(email string) (userModel User) {
	database.DB.Where("email=?", email).First(&userModel)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel User) {
	database.DB.Where("phone=?", loginID).Or("email=?", loginID).Or("name=?", loginID).First(&userModel)
	return
}

// Get 通过id获取用户
func Get(id string) (userModel User) {
	database.DB.Where("id", id).First(&userModel)
	return
}

func All() (users []User) {
	database.DB.Find(&users)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
