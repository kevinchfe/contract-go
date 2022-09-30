package config

import (
	"contract/pkg/helpers"
	"github.com/spf13/cast"
	viper2 "github.com/spf13/viper"
	"os"
)

// viper 库实例
var viper *viper2.Viper

// ConfigFunc 动态加载配置
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 1. 初始化viper库
	viper = viper2.New()

	// 2. 配置类型
	viper.SetConfigType("env")

	// 3. 环境变量配置文件查找路劲
	viper.AddConfigPath(".")

	// 4. 设置环境变量前缀 用以区分go的系统环境变量
	viper.SetEnvPrefix("appenv")

	// 5. 读取环境变量
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	// 1. 加载环境变量
	loadEnv(env)

	// 2. 注册配置信息
	loadConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn)
	}
}

func loadEnv(envSuffix string) {
	// 默认加载 .env 文件 如果有传参 --env=name 则加载.env.name文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := ".env" + envSuffix
		if _, err := os.Stat(filePath); err == nil {
			envPath = filePath
		}
	}

	// 加载env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 监控 env 文件
	viper.WatchConfig()
}

// Env 读取环境变量 支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// Add 增加配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取 如: app.name
// 第二个参数允许传参默认值
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况下
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// GetString 获取 string 类型配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 int 类型配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
