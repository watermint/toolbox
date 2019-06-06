package app_control_impl

import (
	"errors"
	"github.com/watermint/toolbox/experimental/app_control"
	"github.com/watermint/toolbox/experimental/app_log"
	"github.com/watermint/toolbox/experimental/app_msg_container_impl"
	"github.com/watermint/toolbox/experimental/app_ui"
	"github.com/watermint/toolbox/experimental/app_workspace"
	"go.uber.org/zap"
	"os"
)

func NewMock() app_control.Control {
	mc := &app_msg_container_impl.Alt{}
	return &mockControl{
		logger: app_log.NewConsoleLogger(false),
		ui:     app_ui.NewConsole(mc),
		ws:     app_workspace.NewTempWorkspace(),
	}
}

type mockControl struct {
	logger *zap.Logger
	ui     app_ui.UI
	ws     app_workspace.Workspace
}

func (z *mockControl) IsSecure() bool {
	return false
}

func (z *mockControl) IsQuiet() bool {
	return false
}

func (z *mockControl) IsTest() bool {
	return false
}

func (z *mockControl) Workspace() app_workspace.Workspace {
	return z.ws
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
	os.Exit(app_control.FatalGeneral)
}

func (z *mockControl) UI() app_ui.UI {
	return z.ui
}

func (z *mockControl) Log() *zap.Logger {
	return z.logger
}
