package es_container

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/essentials/log/es_rotate"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_budget"
)

type Logger interface {
	// Current logger
	Logger() es_log.Logger

	// Set rotate hook
	SetRotateHook(hook es_rotate.RotateHook)

	// Close logger
	Close()
}

func NewDual(basePath string, budget app_budget.Budget) (t, c Logger, err error) {
	t, err = NewToolbox(basePath, budget)
	if err != nil {
		return nil, nil, err
	}
	c, err = NewCapture(basePath, budget)
	if err != nil {
		t.Close()
		return nil, nil, err
	}
	return t, c, nil
}

func NewToolbox(basePath string, budget app_budget.Budget) (c Logger, err error) {
	return New(basePath, app.LogToolbox, budget, es_log.LevelDebug, es_log.FlavorFileStandard, true)
}

func NewCapture(basePath string, budget app_budget.Budget) (c Logger, err error) {
	return New(basePath, app.LogCapture, budget, es_log.LevelDebug, es_log.FlavorFileCompact, false)
}

func New(basePath, name string, budget app_budget.Budget, level es_log.Level, flavor es_log.Flavor, teeConsole bool) (c Logger, err error) {
	w := es_rotate.NewWriter(basePath, name)

	cs, qt, nb := app_budget.StorageBudget(budget)
	err = w.Open(
		es_rotate.Compress(),
		es_rotate.ChunkSize(cs),
		es_rotate.Quota(qt),
		es_rotate.NumBackup(nb),
	)
	if err != nil {
		return
	}
	l := es_log.NewLogCloser(level, flavor, w)
	if teeConsole {
		return newTee(w, l), nil
	} else {
		return &ctnImpl{
			w: w,
			l: l,
		}, nil
	}
}

type ctnImpl struct {
	w es_rotate.Writer
	l es_log.LogCloser
}

func (z ctnImpl) Logger() es_log.Logger {
	return z.l
}

func (z ctnImpl) SetRotateHook(hook es_rotate.RotateHook) {
	z.w.UpdateOpt(es_rotate.HookBeforeDelete(hook))
}

func (z ctnImpl) Close() {
	_ = z.l.Close()
	_ = z.w.Close()
}

func newTee(w es_rotate.Writer, l es_log.LogCloser) Logger {
	t := es_log.NewTee()
	t.AddSubscriber(l)
	t.AddSubscriber(es_log.DefaultConsole())
	return &teeImpl{
		w: w,
		l: l,
		t: t,
	}
}

type teeImpl struct {
	w es_rotate.Writer
	l es_log.LogCloser
	t es_log.Logger
}

func (z teeImpl) Logger() es_log.Logger {
	return z.t
}

func (z teeImpl) SetRotateHook(hook es_rotate.RotateHook) {
	z.w.UpdateOpt(es_rotate.HookBeforeDelete(hook))
}

func (z teeImpl) Close() {
	_ = z.l.Close()
	_ = z.w.Close()
}
