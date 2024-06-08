package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
var Log *zap.Logger
func InitLogger() error {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", "log/app.log"} // 输出到控制台和文件
	config.ErrorOutputPaths = []string{"stderr"}             // 错误日志单独输出到 stderr
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式化
	logger, err := config.Build()
	if err != nil {
		return  err
	}
	Log = logger
	return nil
}