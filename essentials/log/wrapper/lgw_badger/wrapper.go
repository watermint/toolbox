package lgw_badger

import (
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewLogWrapper(logger esl.Logger) badger.Logger {
	return badgerLoggerWrapper{
		logger: logger,
	}
}

type badgerLoggerWrapper struct {
	logger esl.Logger
}

func (z badgerLoggerWrapper) Errorf(s string, i ...interface{}) {
	msg := fmt.Sprintf(s, i...)
	z.logger.Debug(msg, esl.String("level", "error"))
}

func (z badgerLoggerWrapper) Warningf(s string, i ...interface{}) {
	msg := fmt.Sprintf(s, i...)
	z.logger.Debug(msg, esl.String("level", "warn"))
}

func (z badgerLoggerWrapper) Infof(s string, i ...interface{}) {
	msg := fmt.Sprintf(s, i...)
	z.logger.Debug(msg, esl.String("level", "info"))
}

func (z badgerLoggerWrapper) Debugf(s string, i ...interface{}) {
	msg := fmt.Sprintf(s, i...)
	z.logger.Debug(msg, esl.String("level", "debug"))
}
