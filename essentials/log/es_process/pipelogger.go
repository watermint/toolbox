package es_process

import (
	"bufio"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
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

func NewLogger(cmd *exec.Cmd, ctl app_control.Control) Logger {
	return &loggerImpl{
		ctl: ctl,
		cmd: cmd,
	}
}

type loggerImpl struct {
	ctl     app_control.Control
	cmd     *exec.Cmd
	so      io.ReadCloser
	se      io.ReadCloser
	running bool
}

func (z *loggerImpl) logger(r io.Reader, prefix string) {
	l := z.ctl.Log()
	sb := bufio.NewReader(r)
	w := sync.Once{}
	for {
		line, _, err := sb.ReadLine()
		switch err {
		case io.EOF, os.ErrClosed:
			return
		case nil:
			l.Info(prefix, es_log.String("line", string(line)))
		default:
			w.Do(func() {
				l.Warn(prefix+": Read error", es_log.Error(err))
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
	l := z.ctl.Log()
	z.running = true
	z.so, err = z.cmd.StdoutPipe()
	if err != nil {
		l.Warn("Unable to create pipe of stdout", es_log.Error(err))
		z.so = nil
	} else {
		go z.logger(z.so, "STDOUT")
	}
	z.se, err = z.cmd.StderrPipe()
	if err != nil {
		l.Warn("Unable to create pipe of stderr", es_log.Error(err))
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
