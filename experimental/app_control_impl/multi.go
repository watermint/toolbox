package app_control_impl

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/experimental/app_control"
	"github.com/watermint/toolbox/experimental/app_log"
	"github.com/watermint/toolbox/experimental/app_msg_container"
	"github.com/watermint/toolbox/experimental/app_ui"
	"github.com/watermint/toolbox/experimental/app_workspace"
	"go.uber.org/zap"
)

func NewMulti(ws app_workspace.Workspace, ui app_ui.UI, bx *rice.Box, mc app_msg_container.Container, quiet bool) app_control.Control {
	return &Multi{
		ws:    ws,
		ui:    ui,
		box:   bx,
		mc:    mc,
		quiet: quiet,
	}
}

type Multi struct {
	ui     app_ui.UI
	flc    *app_log.FileLogContext
	cap    *app_log.CaptureContext
	box    *rice.Box
	mc     app_msg_container.Container
	ws     app_workspace.Workspace
	quiet  bool
	secure bool
}

func (z *Multi) Up(opts ...app_control.UpOpt) (err error) {
	opt := &app_control.UpOpts{}
	for _, o := range opts {
		o(opt)
	}
	z.secure = opt.Secure
	if z.ws == nil {
		z.ws = opt.Workspace
	}

	z.flc, err = app_log.NewFileLogger(z.ws.Log(), opt.Debug)
	if err != nil {
		return err
	}

	z.cap, err = app_log.NewCaptureLogger(z.ws.Log())
	if err != nil {
		return err
	}

	z.Log().Debug("Up completed")

	return nil
}

func (z *Multi) Down() {
	z.Log().Debug("Down")
	z.cap.Close()
	z.flc.Close()
}

func (z *Multi) Abort(opts ...app_control.AbortOpt) {
	opt := &app_control.AbortOpts{}
	for _, o := range opts {
		o(opt)
	}
	z.Log().Debug("Abort shutdown", zap.Any("opt", opt))
	z.cap.Close()
	z.flc.Close()
}

func (z *Multi) UI() app_ui.UI {
	return z.ui
}

func (z *Multi) Log() *zap.Logger {
	return z.flc.Logger
}

func (z *Multi) Capture() *zap.Logger {
	return z.cap.Logger
}

func (z *Multi) Resource(key string) (bin []byte, err error) {
	return z.box.Bytes(key)
}

func (z *Multi) Workspace() app_workspace.Workspace {
	return z.ws
}

func (z *Multi) IsProduction() bool {
	return isProduction()
}

func (z *Multi) IsTest() bool {
	return false
}

func (z *Multi) IsQuiet() bool {
	return z.quiet
}

func (z *Multi) IsSecure() bool {
	return z.secure
}
