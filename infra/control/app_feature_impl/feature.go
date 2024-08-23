package app_feature_impl

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/infra/control/app_budget"
	app_definitions2 "github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/report/rp_artifact_feature"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	seqReplay    []nw_replay.Response
	hashReplay   kv_storage.Storage
}

func (z featureImpl) UIReportFilter() (filter string, enabled bool) {
	return z.com.OutputFilter.Value(), z.com.OutputFilter.IsExists() && z.com.OutputFilter.Value() != ""
}

func (z featureImpl) KvsEngine() kv_storage.KvsEngine {
	switch {
	case z.Experiment(app_definitions2.ExperimentKvsBadger):
		return kv_storage.KvsEngineBadger
	case z.Experiment(app_definitions2.ExperimentKvsBadgerTurnstile):
		return kv_storage.KvsEngineBadgerTurnstile
	case z.Experiment(app_definitions2.ExperimentKvsBitcask):
		return kv_storage.KvsEngineBitcask
	case z.Experiment(app_definitions2.ExperimentKvsBitcaskTurnstile):
		return kv_storage.KvsEngineBitcaskTurnstile
	case z.Experiment(app_definitions2.ExperimentKvsSqlite):
		return kv_storage.KvsEngineSqlite
	case z.Experiment(app_definitions2.ExperimentKvsSqliteTurnstile):
		return kv_storage.KvsEngineSqliteTurnstile
	default:
		return kv_storage.KvsEngineBadger
	}
}

func (z featureImpl) IsDefaultPathAuthRepository() bool {
	return !z.com.AuthDatabase.IsExists()
}

func (z featureImpl) PathAuthRepository() string {
	if z.com.AuthDatabase.IsExists() {
		return z.com.AuthDatabase.Value()
	} else {
		return filepath.Join(z.ws.Secrets(), app_definitions2.AuthDatabaseDefaultName)
	}
}

func (z featureImpl) IsSkipLogging() bool {
	return z.com.SkipLogging
}

func (z featureImpl) Extra() app_opt.ExtraOpts {
	return z.com.ExtraOpts()
}

func (z featureImpl) Experiment(name string) bool {
	if z.com.ExtraOpts().HasExperiment(name) {
		return true
	}

	experiments := strings.Split(z.com.Experiment, ",")
	for _, experiment := range experiments {
		if experiment == name {
			return true
		}
	}
	return false
}

func (z featureImpl) AsReplayTest(replays kv_storage.Storage) app_feature.Feature {
	z.hashReplay = replays
	z.test = true
	return z
}

func (z featureImpl) AsSeqReplayTest(replay []nw_replay.Response) app_feature.Feature {
	z.seqReplay = replay
	z.test = true
	return z
}

func (z featureImpl) IsTestWithReplay() (replay kv_storage.Storage, enabled bool) {
	if z.hashReplay != nil {
		return z.hashReplay, true
	} else {
		return nil, false
	}
}

func (z featureImpl) IsTestWithSeqReplay() (replay []nw_replay.Response, enabled bool) {
	if len(z.seqReplay) > 0 {
		return z.seqReplay, true
	} else {
		return nil, false
	}
}

func (z featureImpl) IsVerbose() bool {
	return z.com.Verbose
}

func (z featureImpl) IsTransient() bool {
	return z.transient
}

func (z featureImpl) wsPathConfig() string {
	return filepath.Join(z.ws.Home(), ConfigFileName)
}

func (z featureImpl) loadConfig() (values map[string]interface{}, err error) {
	// #586 : Load config $HOME/.config/watermint-toolbox/config.json
	configPath, err := app_workspace.GetOrCreateDefaultAppConfigPath()
	if err != nil {
		return nil, err
	}
	if values, err = z.loadConfigPath(filepath.Join(configPath, ConfigFileName)); values != nil {
		return values, nil
	}

	// Fallback to old config path: $HOME/.toolbox/config.json
	p := z.wsPathConfig()
	values, err = z.loadConfigPath(p)
	return
}

func (z featureImpl) loadConfigPath(path string) (values map[string]interface{}, err error) {
	values = make(map[string]interface{})
	l := esl.Default()

	_, err = os.Lstat(path)
	if err != nil {
		l.Debug("No file information; skip loading", esl.Error(err))
		return values, nil
	}

	l.Debug("load config", esl.String("path", path))
	b, err := ioutil.ReadFile(path)
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
	cp, err := app_workspace.GetOrCreateDefaultAppConfigPath()
	if err != nil {
		l.Debug("Unable to determine app config path", esl.Error(err))
		return err
	}
	configPath := filepath.Join(cp, ConfigFileName)

	l.Debug("load config", esl.String("path", configPath))
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
	if err := ioutil.WriteFile(configPath, b, 0644); err != nil {
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
	return app_definitions2.IsProduction()
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
	aof := &rp_artifact_feature.OptInFeatureAutoOpen{}
	if f, found := z.OptInGet(aof); found && f.OptInIsEnabled() {
		return true
	}
	return z.com.AutoOpen
}

func (z featureImpl) UIFormat() string {
	return z.com.Output.Value()
}
