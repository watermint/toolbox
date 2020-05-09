package app_workspace

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/esl_container"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"io"
	"os"
)

// Workspace bundle
type Bundle interface {
	io.Closer

	// Workspace
	Workspace() Workspace

	// Logger
	Logger() esl_container.Logger

	// REST logger
	Capture() esl_container.Logger

	// Summary logger
	Summary() esl_container.Logger

	// Storage budget
	Budget() app_budget.Budget

	// Log level for console logs
	ConsoleLogLevel() esl.Level

	// True when the workspace bundle is transient.
	IsTransient() bool
}

func ForkBundle(wb Bundle, name string) (bundle Bundle, err error) {
	return ForkBundleWithLevel(wb, name, wb.ConsoleLogLevel())
}

func ForkBundleWithLevel(wb Bundle, name string, consoleLevel esl.Level) (bundle Bundle, err error) {
	nws, err := Fork(wb.Workspace(), name)
	if err != nil {
		return nil, err
	}
	l, c, s, err := esl_container.NewAll(nws.Log(), wb.Budget(), consoleLevel)
	if err != nil {
		return nil, err
	}
	return newBundleInternal(nws, wb.Budget(), c, l, s, consoleLevel, wb.IsTransient()), nil
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

func NewBundle(home string, budget app_budget.Budget, consoleLevel esl.Level, transient bool) (bundle Bundle, err error) {
	ws, err := NewWorkspace(home, transient)
	if err != nil {
		return nil, err
	}

	if transient {
		l, c, s := esl_container.NewTransient(consoleLevel)
		return newBundleInternal(ws, budget, c, l, s, consoleLevel, true), nil
	} else {
		l, c, s, err := esl_container.NewAll(ws.Log(), budget, consoleLevel)
		if err != nil {
			return nil, err
		}
		return newBundleInternal(ws, budget, c, l, s, consoleLevel, false), nil
	}
}

func newBundleInternal(ws Workspace, budget app_budget.Budget, capture, logger, summary esl_container.Logger, consoleLevel esl.Level, transient bool) Bundle {
	return &bdlImpl{
		budget:    budget,
		conLv:     consoleLevel,
		capture:   capture,
		logger:    logger,
		summary:   summary,
		ws:        ws,
		transient: transient,
	}
}

type bdlImpl struct {
	budget    app_budget.Budget
	conLv     esl.Level
	capture   esl_container.Logger
	logger    esl_container.Logger
	summary   esl_container.Logger
	ws        Workspace
	transient bool
}

func (z bdlImpl) IsTransient() bool {
	return z.transient
}

func (z bdlImpl) Summary() esl_container.Logger {
	return z.summary
}

func (z bdlImpl) ConsoleLogLevel() esl.Level {
	return z.conLv
}

func (z bdlImpl) Budget() app_budget.Budget {
	return z.budget
}

func (z bdlImpl) Close() error {
	z.logger.Close()
	z.capture.Close()
	z.summary.Close()

	if z.transient {
		return os.RemoveAll(z.ws.Job())
	}
	return nil
}

func (z bdlImpl) Workspace() Workspace {
	return z.ws
}

func (z bdlImpl) Logger() esl_container.Logger {
	return z.logger
}

func (z bdlImpl) Capture() esl_container.Logger {
	return z.capture
}
