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

	// Workspace
	Workspace() Workspace

	// Logger
	Logger() es_container.Logger

	// REST logger
	Capture() es_container.Logger

	// Summary logger
	Summary() es_container.Logger

	// Storage budget
	Budget() app_budget.Budget

	// Log level for console logs
	ConsoleLogLevel() es_log.Level
}

func ForkBundle(wb Bundle, name string) (bundle Bundle, err error) {
	return ForkBundleWithLevel(wb, name, wb.ConsoleLogLevel())
}

func ForkBundleWithLevel(wb Bundle, name string, consoleLevel es_log.Level) (bundle Bundle, err error) {
	nws, err := Fork(wb.Workspace(), name)
	if err != nil {
		return nil, err
	}
	l, c, s, err := es_container.NewAll(nws.Log(), wb.Budget(), consoleLevel)
	if err != nil {
		return nil, err
	}
	return newBundleInternal(nws, wb.Budget(), c, l, s, consoleLevel), nil
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
	l, c, s, err := es_container.NewAll(ws.Log(), budget, consoleLevel)
	if err != nil {
		return nil, err
	}
	return newBundleInternal(
		ws,
		budget,
		c,
		l,
		s,
		consoleLevel,
	), nil
}

func newBundleInternal(ws Workspace, budget app_budget.Budget, capture, logger, summary es_container.Logger, consoleLevel es_log.Level) Bundle {
	return &bdlImpl{
		budget:  budget,
		conLv:   consoleLevel,
		capture: capture,
		logger:  logger,
		summary: summary,
		ws:      ws,
	}
}

type bdlImpl struct {
	budget  app_budget.Budget
	conLv   es_log.Level
	capture es_container.Logger
	logger  es_container.Logger
	summary es_container.Logger
	ws      Workspace
}

func (z bdlImpl) Summary() es_container.Logger {
	return z.summary
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
	z.summary.Close()
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
