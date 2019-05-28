package app_control_impl

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app86/app_control"
	"github.com/watermint/toolbox/app86/app_log"
	"github.com/watermint/toolbox/app86/app_root"
	"github.com/watermint/toolbox/app86/app_ui"
	"github.com/watermint/toolbox/app86/app_workspace"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func NewControl(ui app_ui.UI, bx *rice.Box) app_control.Control {
	return &Control{
		ui:  ui,
		box: bx,
	}
}

type Control struct {
	ui  app_ui.UI
	flc *app_log.FileLogContext
	box *rice.Box
	ws  app_workspace.Workspace
}

func (z *Control) IsTest() bool {
	return false
}

func (z *Control) Workspace() app_workspace.Workspace {
	return z.ws
}

func (z *Control) Resource(key string) (bin []byte, err error) {
	return z.box.Bytes(key)
}

func (z *Control) Startup(opts ...app_control.StartupOpt) (err error) {
	opt := &app_control.StartupOpts{}
	for _, o := range opts {
		o(opt)
	}

	z.ws, err = app_workspace.NewWorkspace(opt.WorkspacePath)
	if err != nil {
		return err
	}

	z.flc, err = app_log.NewFileLogger(filepath.Join(z.ws.Log()), opt.Debug)
	if err != nil {
		return err
	}
	// Overwrite logger
	app_root.SetLogger(z.flc.Logger)

	z.Log().Debug("Startup completed")

	return nil
}

func (z *Control) Shutdown() {
	z.Log().Debug("Shutdown")
	app_root.Flush()
	z.flc.Close()
}

func (z *Control) Fatal(opts ...app_control.FatalOpt) {
	opt := &app_control.FatalOpts{}
	for _, o := range opts {
		o(opt)
	}
	z.Log().Debug("Fatal shutdown", zap.Any("opt", opt))
	app_root.Flush()
	z.flc.Close()

	if opt.Reason == nil {
		os.Exit(app_control.FatalGeneral)
	} else {
		os.Exit(*opt.Reason)
	}
}

func (z *Control) UI() app_ui.UI {
	return z.ui
}

func (z *Control) Log() *zap.Logger {
	return z.flc.Logger
}
