package seeders

import "contract/pkg/seed"

func Initialize() {
	// 触发加载本目录下的init方法
	// 指定优先于同目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
