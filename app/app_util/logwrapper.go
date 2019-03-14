package app_util

import (
	"bytes"
	"go.uber.org/zap"
)

var (
	lineSeparators = [][]byte{
		[]byte("\r"),
		[]byte("\n"),
		[]byte("\r\n"),
		[]byte("\n\r"),
	}
)

func NewLogWrapper(size int, logger *zap.Logger) *LogWrapper {
	return &LogWrapper{
		buf:     make([]byte, size),
		bufSize: size,
		logger:  logger,
	}
}

type LogWrapper struct {
	logger  *zap.Logger
	buf     []byte
	bufSize int
	offset  int
}

func (z *LogWrapper) SetLogger(logger *zap.Logger) {
	z.logger = logger
}

func (z *LogWrapper) Flush() {
	if z.offset == 0 {
		return
	}
	z.logger.Debug(string(z.buf[:z.offset]))
	z.offset = 0
}

func (z *LogWrapper) Write(p []byte) (n int, err error) {
	var q, s int
	q = 0
	s = len(p)
	for _, sp := range lineSeparators {
		i := bytes.Index(p[q:], sp)
		if i >= 0 {
			adv := len(sp) + i
			z.logger.Debug(string(p[q:i]))
			s = s - adv
			if s < 1 {
				return len(p), nil
			}
			q += adv
		}
	}

	for {
		bufLeft := z.bufSize - z.offset
		if bufLeft < s {
			copy(z.buf[z.offset:], p[q:q+bufLeft])
			z.offset += bufLeft
			z.Flush()
			s = s - bufLeft
			q = q + bufLeft
		} else {
			copy(z.buf[z.offset:], p[q:])
			z.offset += s
			return len(p), nil
		}
	}
}
