package app_control_impl

import (
	"github.com/watermint/toolbox/app86/app_control"
	"github.com/watermint/toolbox/app86/app_ui"
	"go.uber.org/zap"
	"os"
)

func NewMock() app_control.Control {
	return &mockControl{
		logger: newConsoleLogger(),
		ui:     &app_ui.Mock{},
	}
}

type mockControl struct {
	logger *zap.Logger
	ui     app_ui.UI
}

func (z *mockControl) Startup(opts ...app_control.StartupOpt) error {
	z.logger.Debug("Mock startup")
	return nil
}

func (z *mockControl) Shutdown() {
	z.logger.Debug("Mock shutdown")
	z.logger.Sync()
}

func (z *mockControl) Fatal(opts ...app_control.FatalOpt) {
	z.logger.Debug("Mock fatal", zap.Any("opts", opts))
	z.logger.Sync()
	os.Exit(app_control.FatalMock)
}

func (z *mockControl) UI() app_ui.UI {
	return z.ui
}

func (z *mockControl) Log() *zap.Logger {
	return z.logger
}
