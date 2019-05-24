package app_control_impl

import (
	"errors"
	"github.com/watermint/toolbox/app86/app_control"
	"github.com/watermint/toolbox/app86/app_log"
	"github.com/watermint/toolbox/app86/app_msg_container_impl"
	"github.com/watermint/toolbox/app86/app_ui"
	"go.uber.org/zap"
	"os"
)

func NewMock() app_control.Control {
	mc := &app_msg_container_impl.Alt{}
	return &mockControl{
		logger: app_log.newConsoleLogger(),
		ui:     app_ui.NewConsole(mc),
	}
}

type mockControl struct {
	logger *zap.Logger
	ui     app_ui.UI
}

func (z *mockControl) Resource(key string) (bin []byte, err error) {
	return nil, errors.New("no resource on the mock")
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
