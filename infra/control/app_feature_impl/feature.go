package app_feature_impl

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	ConfigFileName = "config.json"
)

var (
	ErrorValueNotFound = errors.New("value not found")
)

func NewFeature(opts app_opt.CommonOpts, ws app_workspace.Workspace, transient bool) app_feature.Feature {
	return &featureImpl{
		com:       opts,
		ws:        ws,
		transient: transient,
	}
}

type featureImpl struct {
	com          app_opt.CommonOpts
	ws           app_workspace.Workspace
	test         bool
	testWithMock bool
	transient    bool
}

func (z featureImpl) IsTransient() bool {
	return z.transient
}

func (z featureImpl) pathConfig() string {
	return filepath.Join(z.ws.Home(), ConfigFileName)
}

func (z featureImpl) loadConfig() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})
	l := esl.Default()
	p := z.pathConfig()

	_, err = os.Lstat(p)
	if err != nil {
		l.Debug("No file information; skip loading", esl.Error(err))
		return values, nil
	}

	l.Debug("load config", esl.String("path", p))
	b, err := ioutil.ReadFile(p)
	if err != nil {
		l.Debug("Unable to read config", esl.Error(err))
		return
	}
	if err := json.Unmarshal(b, &values); err != nil {
		l.Debug("unable to unmarshal", esl.Error(err))
		return values, err
	}
	return
}

func (z featureImpl) getConfig(key string) (v interface{}, err error) {
	if values, err := z.loadConfig(); err != nil {
		return nil, err
	} else if v, ok := values[key]; ok {
		return v, nil
	} else {
		return nil, ErrorValueNotFound
	}
}

func (z featureImpl) saveConfig(key string, v interface{}) (err error) {
	l := esl.Default()
	p := z.pathConfig()
	l.Debug("load config", esl.String("path", p))
	values, err := z.loadConfig()
	if err != nil {
		return err
	}
	values[key] = v

	b, err := json.Marshal(values)
	if err != nil {
		l.Debug("Unable to marshal", esl.Error(err))
		return err
	}
	if err := ioutil.WriteFile(p, b, 0644); err != nil {
		l.Debug("Unable to write config", esl.Error(err))
		return err
	}
	return nil
}

func (z featureImpl) ConsoleLogLevel() esl.Level {
	return app_feature.ConsoleLogLevel(z.test, z.com.Debug)
}

func (z featureImpl) AsTest(useMock bool) app_feature.Feature {
	z.test = true
	z.testWithMock = useMock
	return &z
}

func (z featureImpl) AsQuiet() app_feature.Feature {
	z.com.Quiet = true
	return &z
}

func (z featureImpl) OptInGet(oi app_feature.OptIn) (f app_feature.OptIn, found bool) {
	l := esl.Default()
	key := app_feature.OptInName(oi)
	l.Debug("OptInGet", esl.String("key", key))
	if v, err := z.getConfig(key); err != nil {
		l.Debug("The key not found in the current config", esl.Error(err))
		return oi, false
	} else if mv, ok := v.(map[string]interface{}); ok {
		if err := app_feature.OptInFrom(mv, oi); err != nil {
			l.Debug("The value is not a opt-in format", esl.Error(err))
			return oi, false
		}
	}
	return oi, true
}

func (z featureImpl) OptInUpdate(oi app_feature.OptIn) error {
	l := esl.Default()
	key := app_feature.OptInName(oi)
	l = l.With(esl.String("key", key))
	l.Debug("OptInUpdate")
	if err := z.saveConfig(key, oi); err != nil {
		l.Debug("Failed to update opt-in", esl.Error(err))
		return err
	}
	return nil
}

func (z featureImpl) IsTestWithMock() bool {
	return z.testWithMock
}

func (z featureImpl) Home() string {
	return z.com.Workspace.Value()
}

func (z featureImpl) BudgetMemory() app_budget.Budget {
	return app_budget.Budget(z.com.BudgetMemory.Value())
}

func (z featureImpl) BudgetStorage() app_budget.Budget {
	return app_budget.Budget(z.com.BudgetStorage.Value())
}

func (z featureImpl) Concurrency() int {
	return z.com.Concurrency
}

func (z featureImpl) IsProduction() bool {
	return app.IsProduction()
}

func (z featureImpl) IsDebug() bool {
	return z.com.Debug
}

func (z featureImpl) IsTest() bool {
	return z.test
}

func (z featureImpl) IsQuiet() bool {
	return z.com.Quiet
}

func (z featureImpl) IsSecure() bool {
	return z.com.Secure
}

func (z featureImpl) IsAutoOpen() bool {
	return z.com.AutoOpen
}

func (z featureImpl) UIFormat() string {
	return z.com.Output.Value()
}
