package logger

import (
	"github.com/allegro/bigcache/v3"
	"go.uber.org/zap"
)

var _ bigcache.Logger = (*IBigCacheLogger)(nil)

var BigCacheLogger = &IBigCacheLogger{}

type IBigCacheLogger struct {
	logger *zap.Logger
}

func (l *IBigCacheLogger) Printf(format string, v ...interface{}) {
	l.logger.Sugar().Infof(format, v...)
}
