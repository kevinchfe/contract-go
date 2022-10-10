package verifycode

import (
	"contract/pkg/app"
	"contract/pkg/config"
	"contract/pkg/redis"
	"time"
)

// RedisStore 实现verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 实现verifycode.Store interface的Set方法
func (s *RedisStore) Set(key string, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}
	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

// Get 实现verifycode.Store interface的Get方法
func (s *RedisStore) Get(key string, clear bool) (value string) {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// Verify 实现verifycode.Store interface的Verify方法
func (s *RedisStore) Verify(key string, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
