package logger

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var _ gormLogger.Interface = (*IGormLogger)(nil)

var GormLogger = &IGormLogger{}

type IGormLogger struct {
	level                     gormLogger.LogLevel
	slowThreshold             time.Duration
	ignoreRecordNotFoundError bool
	logger                    *zap.Logger
}

func (l *IGormLogger) LogMode(logLevel gormLogger.LogLevel) gormLogger.Interface {
	return &IGormLogger{
		logger:                    logger,
		level:                     logLevel,
		slowThreshold:             400 * time.Millisecond,
		ignoreRecordNotFoundError: false,
	}
}

func (l *IGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level < gormLogger.Info {
		return
	}
	l.logger.Sugar().Infof(msg, data)
}

func (l *IGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level < gormLogger.Warn {
		return
	}
	l.logger.Sugar().Warnf(msg, data)
}

func (l *IGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level < gormLogger.Error {
		return
	}
	l.logger.Sugar().Errorf(msg, data)
}

func (l *IGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.level >= gormLogger.Error && (l.ignoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		l.logger.Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.slowThreshold != 0 && elapsed > l.slowThreshold && l.level >= gormLogger.Warn:
		sql, rows := fc()
		l.logger.Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.level >= gormLogger.Info:
		sql, rows := fc()
		l.logger.Debug("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	}
}
