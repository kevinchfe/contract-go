package bootstrap

import (
	"contract/pkg/cache"
	"contract/pkg/config"
	"fmt"
)

func SetupCache() {
	// 初始化缓存专用的redis client
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.user_name"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)
	cache.InitWithCacheStore(rds)
}
