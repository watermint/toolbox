package app_feature_impl

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_config"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"go.uber.org/zap"
)

func NewFeature(opts *app_control.UpOpts, ws app_workspace.Workspace) app_feature.Feature {
	return &Feature{
		opts: opts,
		cfg:  app_config.NewConfig(ws.Home()),
	}
}

type Feature struct {
	opts *app_control.UpOpts
	cfg  app_config.Config
}

func (z *Feature) OptInGet(oi app_feature.OptIn) (f app_feature.OptIn, found bool) {
	l := app_root.Log()
	key := oi.OptInName(oi)
	l.Debug("OptInGet", zap.String("key", key))
	if v, err := z.cfg.Get(key); err != nil {
		l.Debug("The key not found in the current config", zap.Error(err))
		return oi, false
	} else if mv, ok := v.(map[string]interface{}); ok {
		if err := app_feature.OptInFrom(mv, oi); err != nil {
			l.Debug("The value is not a opt-in format", zap.Error(err))
			return oi, false
		}
	}
	return oi, true
}

func (z *Feature) OptInUpdate(oi app_feature.OptIn) error {
	l := app_root.Log()
	key := oi.OptInName(oi)
	l = l.With(zap.String("key", key))
	l.Debug("OptInUpdate")
	if err := z.cfg.Put(key, oi); err != nil {
		l.Debug("Failed to update opt-in", zap.Error(err))
		return err
	}
	return nil
}

func (z *Feature) Config() app_config.Config {
	return z.cfg
}

func (z *Feature) IsProduction() bool {
	return app.IsProduction()
}

func (z *Feature) IsDebug() bool {
	return z.opts.Debug
}

func (z *Feature) IsTest() bool {
	return z.opts.Test
}

func (z *Feature) IsQuiet() bool {
	return z.opts.Quiet
}

func (z *Feature) IsSecure() bool {
	return z.opts.Secure
}

func (z *Feature) IsLowMemory() bool {
	return z.opts.LowMemory
}

func (z *Feature) IsAutoOpen() bool {
	return z.opts.AutoOpen
}

func (z *Feature) UIFormat() string {
	return z.opts.UIFormat
}
