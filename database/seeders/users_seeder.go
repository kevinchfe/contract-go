package seeders

import (
	"contract/database/factories"
	"contract/pkg/console"
	"contract/pkg/logger"
	"contract/pkg/seed"
	"fmt"
	"gorm.io/gorm"
)

func init() {
	// 添加 Seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		// 创建10个用户对象
		users := factories.MakeUsers(10)
		// 批量创建用户
		result := db.Table("users").Create(&users)
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}
		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
