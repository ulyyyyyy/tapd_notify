package logger

import (
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/gin-gonic/gin"
	"github.com/ulyyyyyy/tapd_notify/internal/helper/ginresp"
	"go.uber.org/zap"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var (
	level  zap.AtomicLevel
	logger *zap.Logger
)

// Initialize 初始化日志组件 ( 根据当前加载配置环境 )
func Initialize() (err error) {
	level = zap.NewAtomicLevel()
	if err != nil {
		return err
	}

	// 设置MQ消息等级
	rlog.SetLogLevel("error")

	GormLogger = &IGormLogger{
		logger:                    logger,
		level:                     gormLogger.Info,
		slowThreshold:             400 * time.Millisecond,
		ignoreRecordNotFoundError: false,
	}

	BigCacheLogger = &IBigCacheLogger{
		logger: logger,
	}

	return nil
}

func Error(msg string, fields ...zap.Field) {
	// defer func() {
	// 	_ = logger.Sync()
	// }()
	logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	// defer func() {
	// 	_ = logger.Sync()
	// }()
	logger.Warn(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	// defer func() {
	// 	_ = logger.Sync()
	// }()
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	// defer func() {
	// 	_ = logger.Sync()
	// }()
	logger.Debug(msg, fields...)
}

func Logger() *zap.Logger {
	return logger
}

func GetLevel(c *gin.Context) {
	ginresp.NewSuccess(c, nil)
}

func EditLevel(c *gin.Context) {
	ginresp.NewSuccess(c, nil)
}

func newProductionConfig(level zap.AtomicLevel) zap.Config {
	return zap.Config{
		Level:       level,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func newDevelopment(level zap.AtomicLevel) zap.Config {
	return zap.Config{
		Level:            level,
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}
