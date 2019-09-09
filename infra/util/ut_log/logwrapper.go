package ut_log

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

func NewLogWrapper(logger *zap.Logger) *LogWrapper {
	lw := &LogWrapper{
		logger: logger,
	}
	lw.init()
	return lw
}

type LogWrapper struct {
	logger *zap.Logger
	line   *LineWriter
}

func (z *LogWrapper) init() {
	bufSize := 4096
	z.line = &LineWriter{
		buf:     make([]byte, bufSize),
		bufSize: bufSize,
		writer: func(line string) error {
			z.logger.Debug(line)
			return nil
		},
	}
}

func (z *LogWrapper) SetLogger(logger *zap.Logger) {
	z.logger = logger
}

func (z *LogWrapper) Flush() error {
	return z.line.Flush()
}

func (z *LogWrapper) Write(p []byte) (n int, err error) {
	return z.line.Write(p)
}

func NewLineWriter(writer func(line string) error) *LineWriter {
	bufSize := 4096
	lw := &LineWriter{
		buf:     make([]byte, bufSize),
		bufSize: bufSize,
		writer:  writer,
	}
	return lw
}

type LineWriter struct {
	writer  func(line string) error
	buf     []byte
	bufSize int
	offset  int
}

func (z *LineWriter) Flush() error {
	if z.offset == 0 {
		return nil
	}
	q := z.offset
	z.offset = 0
	if err := z.writer(string(z.buf[:q])); err != nil {
		return err
	}
	return nil
}

func (z *LineWriter) Write(p []byte) (n int, err error) {
	var q, s int
	q = 0
	s = len(p)
	for _, sp := range lineSeparators {
		i := bytes.Index(p[q:], sp)
		if i >= 0 {
			adv := len(sp) + i
			z.Flush()
			z.writer(string(p[q:i]))
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
