package app_control_impl

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app86/app_control"
	"github.com/watermint/toolbox/app86/app_ui"
	"go.uber.org/zap"
)

type Control struct {
	ui  app_ui.UI
	log *zap.Logger
	box *rice.Box
}

func (z *Control) Resource(key string) (bin []byte, err error) {
	return z.box.Bytes(key)
}

func (z *Control) Startup(opts ...app_control.StartupOpt) error {
	return nil
}

func (z *Control) Shutdown() {

}

func (z *Control) Fatal(opts ...app_control.FatalOpt) {

}

func (z *Control) UI() app_ui.UI {
	return z.ui
}

func (z *Control) Log() *zap.Logger {
	return z.log
}
