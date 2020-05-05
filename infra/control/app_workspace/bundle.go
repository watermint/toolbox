package app_workspace

import (
	"github.com/watermint/toolbox/essentials/log/es_container"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"io"
)

// Workspace bundle
type Bundle interface {
	io.Closer
	Workspace() Workspace
	Logger() es_container.Logger
	Capture() es_container.Logger
	Budget() app_budget.Budget
	ConsoleLogLevel() es_log.Level
}

func ForkBundle(wb Bundle, name string) (bundle Bundle, err error) {
	nws, err := Fork(wb.Workspace(), name)
	if err != nil {
		return nil, err
	}
	l, c, err := es_container.NewDual(nws.Log(), wb.Budget(), wb.ConsoleLogLevel())
	if err != nil {
		return nil, err
	}
	return &bdlImpl{
		ws:      nws,
		logger:  l,
		capture: c,
	}, nil
}

func WithFork(wb Bundle, name string, f func(fwb Bundle) error) error {
	fwb, err := ForkBundle(wb, name)
	if err != nil {
		return err
	}
	defer func() {
		_ = fwb.Close()
	}()
	return f(fwb)
}

func NewBundle(home string, budget app_budget.Budget, consoleLevel es_log.Level) (bundle Bundle, err error) {
	ws, err := NewWorkspace(home)
	if err != nil {
		return nil, err
	}
	l, c, err := es_container.NewDual(ws.Log(), budget, consoleLevel)
	if err != nil {
		return nil, err
	}
	return &bdlImpl{
		conLv:   consoleLevel,
		budget:  budget,
		ws:      ws,
		logger:  l,
		capture: c,
	}, nil
}

type bdlImpl struct {
	budget  app_budget.Budget
	conLv   es_log.Level
	capture es_container.Logger
	logger  es_container.Logger
	ws      Workspace
}

func (z bdlImpl) ConsoleLogLevel() es_log.Level {
	return z.conLv
}

func (z bdlImpl) Budget() app_budget.Budget {
	return z.budget
}

func (z bdlImpl) Close() error {
	z.logger.Close()
	z.capture.Close()
	return nil
}

func (z bdlImpl) Workspace() Workspace {
	return z.ws
}

func (z bdlImpl) Logger() es_container.Logger {
	return z.logger
}

func (z bdlImpl) Capture() es_container.Logger {
	return z.capture
}
