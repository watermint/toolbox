package app_control_impl

import (
	"encoding/json"
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

func NewControl(ui app_ui.UI, bx *rice.Box, quiet bool) app_control.Control {
	return &Control{
		ui:    ui,
		box:   bx,
		quiet: quiet,
	}
}

type Control struct {
	ui     app_ui.UI
	flc    *app_log.FileLogContext
	cap    *app_log.CaptureContext
	box    *rice.Box
	ws     app_workspace.Workspace
	quiet  bool
	secure bool
}

func (z *Control) IsSecure() bool {
	return z.secure
}

func (z *Control) IsQuiet() bool {
	return z.quiet
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
	z.secure = opt.Secure

	z.ws, err = app_workspace.NewWorkspace(opt.WorkspacePath)
	if err != nil {
		return err
	}

	rl, err := os.Create(filepath.Join(z.ws.Log(), "recipe.log"))
	if err != nil {
		return err
	}
	type RecipeLog struct {
		Name string `json:"name"`
	}
	rr := &RecipeLog{
		Name: opt.RecipeName,
	}
	rb, err := json.Marshal(rr)
	if err != nil {
		return err
	}
	rl.Write(rb)
	rl.Close()

	z.flc, err = app_log.NewFileLogger(z.ws.Log(), opt.Debug)
	if err != nil {
		return err
	}

	z.cap, err = app_log.NewCaptureLogger(z.ws.Log())
	if err != nil {
		return err
	}

	// Overwrite logger
	app_root.SetLogger(z.flc.Logger)
	app_root.SetCapture(z.cap.Logger)

	z.Log().Debug("Startup completed")

	return nil
}

func (z *Control) Shutdown() {
	z.Log().Debug("Shutdown")
	app_root.Flush()
	z.cap.Close()
	z.flc.Close()
}

func (z *Control) Fatal(opts ...app_control.FatalOpt) {
	opt := &app_control.FatalOpts{}
	for _, o := range opts {
		o(opt)
	}
	z.Log().Debug("Fatal shutdown", zap.Any("opt", opt))
	app_root.Flush()
	z.cap.Close()
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
