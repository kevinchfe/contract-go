package middlewares

import (
	"bytes"
	"contract/pkg/helpers"
	"contract/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"io"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// Logger 记录请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取response内容
		w := &responseBodyWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = w
		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body是一个buffer对象，只能读取一次
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 读取后，重新赋值c.Request.Body,以供后续操作
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 设置开始时间
		start := time.Now()
		c.Next()
		// 开始记录日志的逻辑
		cost := time.Since(start)
		responStatus := c.Writer.Status()
		logFields := []zap.Field{
			zap.Int("status", responStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			// 请求内容
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))
			// 响应内容
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if responStatus > 400 && responStatus <= 499 {
			// 除了StatusBadRequest以外，warning提示一下，常见的有403，404
			logger.Warn("HTTP Warning "+cast.ToString(responStatus), logFields...)
		} else if responStatus > 500 && responStatus <= 599 {
			// 除了内部错误，记录error
			logger.Error("HTTP Error "+cast.ToString(responStatus), logFields...)
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}
	}
}
