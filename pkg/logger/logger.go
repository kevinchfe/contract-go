package logger

import (
	"contract/pkg/app"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

// Logger 全局Logger对象
var Logger *zap.Logger

// InitLogger 日志初始化
func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {
	// 获取日志写入介质
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)
	// 日志等级
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}

	// 初始化core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)
	//初始化logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // 调用文件和行号内部使用runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error时才会显示stacktrace
	)
	// 将自定义的logger替换为全局的logger
	zap.ReplaceGlobals(Logger)
}

// getEncoder 是指日志存储格式
func getEncoder() zapcore.Encoder {
	// 日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // 代码调用
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR/INFO
		EncodeTime:     customTimeEncoder,              // 时间格式 自定义
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller短格式
	}
	if app.IsLocal() {
		// 本地终端输出关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		// 本地设置内置的Console解码器
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	// 线上环境 JSON 编码
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// getLogWriter 日志记录介质
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {
	// 如果配置了按照日期记录日志文件
	if logType == "daily" {
		logname := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logname)
	}

	// 滚动日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	// 配置输出介质
	if app.IsLocal() {
		// 本地开发终端打印并记录文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		// 生产环境只记录文件
		return zapcore.AddSync(lumberJackLogger)
	}
}

// Dump 调试专用 不会中断程序，在终端打印除warning消息
// 第一个参数会使用 json.Marshal渲染 ，第二个参数信息可选
// logger.Dump(user.User{Name:"test"})
// logger.Dump(user.User{Name:"test"}, "用户信息")
func Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("data", valueString))
	}
}

// LogIf 当err!=nil时记录error等级错误
func LogIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred:", zap.Error(err))
	}
}

// LogWarnIf 当err!=nil时记录warning等级日志
func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

// LogInfoIf 当err!=nil时记录info等级日志
func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}

// Debug 调试日志
// logger.Debug("Database", zap.String("sql", sql))
func Debug(ModuleName string, fields ...zap.Field) {
	Logger.Debug(ModuleName, fields...)
}

// Info 告知类日志
func Info(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

// Warn 警告类日志
func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

// Error 错误时记录
func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

// Fatal 级别同Error,写完log后调用os.Exit(1)退出程序
func Fatal(moduleName string, fields ...zap.Field) {
	Logger.Fatal(moduleName, fields...)
}

// DebugString 记录一条字符串类型的debug日志
// logger.DebugString("SMS", "短信内容", string(result.RawResponse))
func DebugString(moduleName, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON 记录对象类型的 debug 日志，使用 json.Marshal 进行编码
// logger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser())
func DebugJSON(moduleName, name string, value interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName, name string, value interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}
