package app_feature_impl

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_config"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
)

func NewFeature(opts app_opt.CommonOpts, ws app_workspace.Workspace) app_feature.Feature {
	return &Feature{
		com: opts,
		cfg: app_config.NewConfig(ws.Home()),
	}
}

type Feature struct {
	com          app_opt.CommonOpts
	cfg          app_config.Config
	test         bool
	testWithMock bool
}

func (z Feature) AsTest(useMock bool) app_feature.Feature {
	z.test = true
	z.testWithMock = useMock
	return &z
}

func (z Feature) AsQuiet() app_feature.Feature {
	z.com.Quiet = true
	return &z
}

func (z Feature) OptInGet(oi app_feature.OptIn) (f app_feature.OptIn, found bool) {
	l := es_log.Default()
	key := oi.OptInName(oi)
	l.Debug("OptInGet", es_log.String("key", key))
	if v, err := z.cfg.Get(key); err != nil {
		l.Debug("The key not found in the current config", es_log.Error(err))
		return oi, false
	} else if mv, ok := v.(map[string]interface{}); ok {
		if err := app_feature.OptInFrom(mv, oi); err != nil {
			l.Debug("The value is not a opt-in format", es_log.Error(err))
			return oi, false
		}
	}
	return oi, true
}

func (z Feature) OptInUpdate(oi app_feature.OptIn) error {
	l := es_log.Default()
	key := oi.OptInName(oi)
	l = l.With(es_log.String("key", key))
	l.Debug("OptInUpdate")
	if err := z.cfg.Put(key, oi); err != nil {
		l.Debug("Failed to update opt-in", es_log.Error(err))
		return err
	}
	return nil
}

func (z Feature) IsTestWithMock() bool {
	return z.testWithMock
}

func (z Feature) Home() string {
	return z.com.Workspace.Value()
}

func (z Feature) BudgetMemory() string {
	return z.com.BudgetMemory.Value()
}

func (z Feature) BudgetStorage() string {
	return z.com.BudgetStorage.Value()
}

func (z Feature) Concurrency() int {
	return z.com.Concurrency
}

func (z Feature) Config() app_config.Config {
	return z.cfg
}

func (z Feature) IsProduction() bool {
	return app.IsProduction()
}

func (z Feature) IsDebug() bool {
	return z.com.Debug
}

func (z Feature) IsTest() bool {
	return z.test
}

func (z Feature) IsQuiet() bool {
	return z.com.Quiet
}

func (z Feature) IsSecure() bool {
	return z.com.Secure
}

func (z Feature) IsLowMemory() bool {
	return z.com.BudgetMemory.Value() == app_opt.BudgetLow
}

func (z Feature) IsAutoOpen() bool {
	return z.com.AutoOpen
}

func (z Feature) UIFormat() string {
	return z.com.Output.Value()
}
