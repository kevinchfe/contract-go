package redis

import (
	"context"
	"contract/pkg/logger"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

// RedisClient Redis服务
type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局redis对象只实例一次
var once sync.Once

// Redis 全局redis
var Redis *RedisClient

// ConnectRedis 链接redis数据库 设置全局redis对象
func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

// NewClient创建一个新的redis连接
func NewClient(address string, username string, password string, db int) *RedisClient {
	// 初始化自定义RedisClient实例
	rds := &RedisClient{}
	// 使用默认的context
	rds.Context = context.Background()
	// 使用redis库里的NewClient初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})
	err := rds.Ping()
	logger.LogIf(err)
	return rds
}

// Ping 测试redis连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储key对应value
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get 获取key对应value
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

// Has 判断一个key是否存在
func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

// Del 删除存储在redis里的数据 支持多个key传参
func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

// FlushDB 情况redis db数据
func (rds RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

// Increment 当参数只有 1 个时，为 key，其值增加 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要增加的值 int64 类型。
func (rds RedisClient) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Increment", "参数过多")
		return false
	}
	return true
}

// Decrement 当参数只有 1 个时，为 key，其值减去 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要减去的值 int64 类型。
func (rds RedisClient) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decrement", "参数过多")
		return false
	}
	return true
}
