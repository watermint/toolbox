package app_control_impl

import (
	"errors"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_log"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/quality/qt_control_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"os"
)

func NewMock() app_control.Control {
	mc := &app_msg_container_impl.Alt{}
	return &mockControl{
		logger: app_log.NewConsoleLogger(false),
		ui:     app_ui.NewConsole(mc, qt_control_impl.NewMessageMock(), false),
		ws:     app_workspace.NewTempAppWorkspace(),
	}
}

type mockControl struct {
	logger *zap.Logger
	ui     app_ui.UI
	ws     app_workspace.Workspace
}

func (z *mockControl) IsProduction() bool {
	return isProduction()
}

func (z *mockControl) Capture() *zap.Logger {
	return z.logger
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

func (z *mockControl) Up(opts ...app_control.UpOpt) error {
	name := app.Name
	ver := app.Version
	hash := app.Hash

	z.logger.Debug("Mock startup",
		zap.String("name", name),
		zap.String("ver", ver),
		zap.String("hash", hash),
	)

	return nil
}

func (z *mockControl) Down() {
	z.logger.Debug("Mock shutdown")
	z.logger.Sync()
}

func (z *mockControl) Abort(opts ...app_control.AbortOpt) {
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
