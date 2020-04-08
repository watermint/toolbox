package app_control_impl

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_config"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workspace"
)

func NewFeature(opts *app_control.UpOpts, ws app_workspace.Workspace) app_control.Feature {
	return &Feature{
		opts: opts,
		cfg:  app_config.NewConfig(ws.Home()),
	}
}

type Feature struct {
	opts *app_control.UpOpts
	cfg  app_config.Config
}

func (z *Feature) IsConfigEnabled(key string) bool {
	if v, err := z.cfg.Get(key); err != nil {
		return false
	} else if b, ok := v.(bool); ok {
		return b
	} else {
		return false
	}
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
