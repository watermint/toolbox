package esl_process

import (
	"bufio"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

type Logger interface {
	Start()
	Close()
}

func NewLogger(cmd *exec.Cmd, logger esl.Logger) Logger {
	return &loggerImpl{
		l:   logger,
		cmd: cmd,
	}
}

type loggerImpl struct {
	l       esl.Logger
	cmd     *exec.Cmd
	so      io.ReadCloser
	se      io.ReadCloser
	running bool
}

func (z *loggerImpl) logger(r io.Reader, prefix string) {
	l := z.l
	sb := bufio.NewReader(r)
	w := sync.Once{}
	for {
		line, _, err := sb.ReadLine()
		switch err {
		case io.EOF, os.ErrClosed:
			return
		case nil:
			l.Info(prefix, esl.String("line", string(line)))
		default:
			w.Do(func() {
				l.Warn(prefix+": Read error", esl.Error(err))
			})
		}
		if !z.running {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func (z *loggerImpl) Start() {
	var err error
	l := z.l
	z.running = true
	z.so, err = z.cmd.StdoutPipe()
	if err != nil {
		l.Warn("Unable to create pipe of stdout", esl.Error(err))
		z.so = nil
	} else {
		go z.logger(z.so, "STDOUT")
	}
	z.se, err = z.cmd.StderrPipe()
	if err != nil {
		l.Warn("Unable to create pipe of stderr", esl.Error(err))
		z.se = nil
	} else {
		go z.logger(z.se, "STDERR")
	}
}

func (z *loggerImpl) Close() {
	if z.so != nil {
		z.so.Close()
		z.so = nil
	}
	if z.se != nil {
		z.se.Close()
		z.se = nil
	}
}
