package orm_logger

import (
	"context"
	"github.com/watermint/toolbox/essentials/log/esl"
	"gorm.io/gorm/logger"
	"time"
)

func NewGormLogger(logger esl.Logger) logger.Interface {
	return &gormLoggerWrapper{
		logger: logger,
	}
}

// gormLoggerWrapper wraps esl.Logger to logger.Interface
type gormLoggerWrapper struct {
	logger esl.Logger
}

func (z gormLoggerWrapper) LogMode(level logger.LogLevel) logger.Interface {
	return z
}

func (z gormLoggerWrapper) Info(ctx context.Context, s string, i ...interface{}) {
	z.logger.Debug("GormInfo", esl.String("s", s), esl.Any("i", i))
}

func (z gormLoggerWrapper) Warn(ctx context.Context, s string, i ...interface{}) {
	z.logger.Debug("GormWarm", esl.String("s", s), esl.Any("i", i))
}

func (z gormLoggerWrapper) Error(ctx context.Context, s string, i ...interface{}) {
	z.logger.Debug("GormError", esl.String("s", s), esl.Any("i", i))
}

func (z gormLoggerWrapper) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//sql, affected := fc()
	//z.logger.Debug("GormTrace", esl.String("sql", sql), esl.Int64("affected", affected), esl.Error(err))
}
