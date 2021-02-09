package esl_container

import (
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/esl_rotate"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_budget"
)

type Logger interface {
	// Current logger. This logger could be wrapped.
	Logger() esl.Logger

	// Core logger.
	Core() esl.Logger

	// Set rotate hook
	SetRotateHook(hook esl_rotate.RotateHook)

	// Close logger
	Close()
}

func NewTransient(consoleLevel esl.Level) (t, c, s Logger) {
	t = newTransient(consoleLevel, esl.FlavorConsole)
	c = newTransient(esl.LevelQuiet, esl.FlavorFileCapture)
	s = newTransient(esl.LevelQuiet, esl.FlavorConsole)
	return
}

func NewAll(basePath string, budget app_budget.Budget, consoleLevel esl.Level) (t, c, s Logger, err error) {
	t, err = NewToolbox(basePath, budget, consoleLevel)
	if err != nil {
		return nil, nil, nil, err
	}
	c, err = NewCapture(basePath, budget)
	if err != nil {
		t.Close()
		return nil, nil, nil, err
	}
	s, err = LogSummary(basePath)
	if err != nil {
		t.Close()
		c.Close()
		return nil, nil, nil, err
	}
	return
}

func NewToolbox(basePath string, budget app_budget.Budget, consoleLevel esl.Level) (c Logger, err error) {
	return New(basePath, app.LogToolbox, budget, esl.LevelDebug, consoleLevel, esl.FlavorFileStandard, true, true)
}

func NewCapture(basePath string, budget app_budget.Budget) (c Logger, err error) {
	return New(basePath, app.LogCapture, budget, esl.LevelDebug, esl.LevelInfo, esl.FlavorFileCapture, false, true)
}

func LogSummary(basePath string) (c Logger, err error) {
	return New(basePath, app.LogSummary, app_budget.BudgetUnlimited, esl.LevelDebug, esl.LevelInfo, esl.FlavorFileStandard, false, false)
}

func New(basePath, name string, budget app_budget.Budget, fileLevel, consoleLevel esl.Level, flavor esl.Flavor, teeConsole, compress bool) (c Logger, err error) {
	w := esl_rotate.NewWriter(basePath, name)

	cs, qt, nb := app_budget.StorageBudget(budget)
	err = w.Open(
		esl_rotate.CompressEnabled(compress),
		esl_rotate.ChunkSize(cs),
		esl_rotate.Quota(qt),
		esl_rotate.NumBackup(nb),
	)
	if err != nil {
		return
	}
	l := esl.NewLogCloser(fileLevel, flavor, w)
	if teeConsole {
		return newTee(w, l, consoleLevel), nil
	} else {
		return &ctnImpl{
			w: w,
			l: l,
		}, nil
	}
}

type ctnImpl struct {
	w esl_rotate.Writer
	l esl.LogCloser
}

func (z ctnImpl) Core() esl.Logger {
	return z.l
}

func (z ctnImpl) Logger() esl.Logger {
	return z.l
}

func (z ctnImpl) SetRotateHook(hook esl_rotate.RotateHook) {
	z.w.UpdateOpt(esl_rotate.HookBeforeDelete(hook))
}

func (z ctnImpl) Close() {
	_ = z.l.Close()
	_ = z.w.Close()
}

func newTee(w esl_rotate.Writer, l esl.LogCloser, consoleLevel esl.Level) Logger {
	t := esl.NewTee()
	t.AddSubscriber(l)
	if consoleLevel != esl.LevelQuiet {
		t.AddSubscriber(esl.New(consoleLevel, esl.FlavorConsole, es_stdout.NewDirectErr()))
	}
	return &teeImpl{
		w: w,
		l: l,
		t: t,
	}
}

type teeImpl struct {
	w esl_rotate.Writer
	l esl.LogCloser
	t esl.Logger
}

func (z teeImpl) Core() esl.Logger {
	return z.l
}

func (z teeImpl) Logger() esl.Logger {
	return z.t
}

func (z teeImpl) SetRotateHook(hook esl_rotate.RotateHook) {
	z.w.UpdateOpt(esl_rotate.HookBeforeDelete(hook))
}

func (z teeImpl) Close() {
	_ = z.l.Close()
	_ = z.w.Close()
}

func newTransient(level esl.Level, flavor esl.Flavor) Logger {
	l := esl.New(level, flavor, es_stdout.NewDirectErr())
	return &transientImpl{
		l: l,
	}
}

type transientImpl struct {
	l esl.Logger
}

func (z transientImpl) Logger() esl.Logger {
	return z.l
}

func (z transientImpl) Core() esl.Logger {
	return z.l
}

func (z transientImpl) SetRotateHook(hook esl_rotate.RotateHook) {
}

func (z transientImpl) Close() {
}
