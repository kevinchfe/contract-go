package bootstrap

import (
	"contract/pkg/config"
	"contract/pkg/redis"
	"fmt"
)

// SetupRedis 初始化redis
func SetupRedis() {
	// 建立redis连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
