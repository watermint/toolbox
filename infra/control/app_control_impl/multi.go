package app_control_impl

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_log"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recpie/app_worker"
	"github.com/watermint/toolbox/infra/recpie/app_worker_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_ui"
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

func (z *Multi) IsLowMemory() bool {
	return false
}

func (z *Multi) Messages() app_msg_container.Container {
	return z.mc
}

func (z *Multi) TestResource(key string) (data gjson.Result, found bool) {
	return gjson.Parse("{}"), false
}

func (z *Multi) NewQueue() app_worker.Queue {
	return app_worker_impl.NewQueue(z, 1)
}

func (z *Multi) Fork(ui app_ui.UI, ws app_workspace.Workspace) (ctl app_control.Control, err error) {
	ctl = &Multi{
		ui:     ui,
		flc:    z.flc,
		cap:    z.cap,
		box:    z.box,
		mc:     z.mc,
		ws:     ws,
		quiet:  z.quiet,
		secure: z.secure,
	}

	opts := make([]app_control.UpOpt, 0)
	if z.secure {
		opts = append(opts, app_control.Secure())
	}
	err = ctl.Up(opts...)

	if err != nil {
		return nil, err
	} else {
		return ctl, nil
	}
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
	name := app.Name
	ver := app.Version
	hash := app.Hash

	z.Log().Debug("Up completed",
		zap.String("name", name),
		zap.String("ver", ver),
		zap.String("hash", hash),
	)

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
