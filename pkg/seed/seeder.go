// Package seed 处理数据库填充相关逻辑
package seed

import "gorm.io/gorm"

// 存放所有Seerder
var seeders []Seeder

// 按顺序执行的Seeder数组
// 支持一些必须按顺序执行的 seeder，例如topic创建时依赖于user,所有 TopicSeeder 应该在 UserSeeder 后执行
var orderedSeederNames []string

type SeederFunc func(db *gorm.DB)

// Seeder 对应每一个database/seeders 目录下的Seeder文件
type Seeder struct {
	Func SeederFunc
	Name string
}

// Add 注册到 seeders 数组
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

// SetRunOrder 设置【按顺序执行的 Seeder 数组】
func SetRunOrder(names []string) {
	orderedSeederNames = names
}
