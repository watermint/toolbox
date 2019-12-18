package app_control_impl

import (
	"encoding/json"
	"errors"
	"github.com/GeertJohan/go.rice"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/control/app_log"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recpie/rc_recipe"
	"github.com/watermint/toolbox/infra/recpie/rc_worker"
	"github.com/watermint/toolbox/infra/recpie/rc_worker_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_template"
	"github.com/watermint/toolbox/infra/ui/app_template_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func NewSingle(ui app_ui.UI, bx, web *rice.Box, mc app_msg_container.Container, quiet bool, catalogue []rc_recipe.Recipe) app_control.Control {
	return &Single{
		ui:           ui,
		box:          bx,
		web:          web,
		mc:           mc,
		quiet:        quiet,
		catalogue:    catalogue,
		testResource: gjson.Parse("{}"),
	}
}

type Single struct {
	ui           app_ui.UI
	flc          *app_log.FileLogContext
	cap          *app_log.CaptureContext
	box          *rice.Box
	web          *rice.Box
	mc           app_msg_container.Container
	ws           app_workspace.Workspace
	opts         *app_control.UpOpts
	quiet        bool
	catalogue    []rc_recipe.Recipe
	testResource gjson.Result
}

func (z *Single) IsLowMemory() bool {
	return z.opts.LowMemory
}

func Fork(ctl app_control.Control, name string) (app_control.Control, error) {
	if fc, ok := ctl.(app_control_launcher.ControlFork); ok {
		return fc.Fork(name)
	}
	return nil, errors.New("fork is not supported on this control")
}

func (z *Single) Fork(name string) (ctl app_control.Control, err error) {
	ws, err := app_workspace.Fork(z.ws, name)
	if err != nil {
		return nil, err
	}
	s := &Single{
		ui:           z.ui,
		box:          z.box,
		web:          z.web,
		mc:           z.mc,
		ws:           ws,
		opts:         z.opts,
		quiet:        z.quiet,
		catalogue:    z.catalogue,
		testResource: z.testResource,
	}
	if err := s.upWithWorkspace(ws); err != nil {
		return nil, err
	}
	return s, nil
}

func (z *Single) With(mc app_msg_container.Container) app_control.Control {
	ui := app_ui.CloneConsole(z.ui, mc)

	return &Single{
		ui:           ui,
		flc:          z.flc,
		cap:          z.cap,
		box:          z.box,
		web:          z.web,
		mc:           mc,
		ws:           z.ws,
		opts:         z.opts,
		quiet:        z.quiet,
		catalogue:    z.catalogue,
		testResource: z.testResource,
	}
}

func (z *Single) Messages() app_msg_container.Container {
	return z.mc
}

func (z *Single) NewControl(user app_workspace.MultiUser) (ctl app_control.Control, err error) {
	ws, err := app_workspace.NewMultiJob(user)
	ctl = NewMulti(ws, z.ui, z.box, z.mc, z.quiet)
	opts := make([]app_control.UpOpt, 0)
	if z.opts.Debug {
		opts = append(opts, app_control.Debug())
	}
	if z.opts.Test {
		opts = append(opts, app_control.Test())
	}
	if z.opts.Secure {
		opts = append(opts, app_control.Secure())
	}
	err = ctl.Up(opts...)
	if err != nil {
		return nil, err
	}
	return ctl, nil
}

func (z *Single) NewTestControl(testResource gjson.Result) (ctl app_control.Control, err error) {
	ctl = &Single{
		ui:           z.ui,
		box:          z.box,
		web:          z.web,
		mc:           z.mc,
		quiet:        z.quiet,
		catalogue:    z.catalogue,
		testResource: testResource,
	}
	opts := make([]app_control.UpOpt, 0)
	opts = append(opts, app_control.Test())
	opts = append(opts, app_control.Concurrency(runtime.NumCPU()))
	err = ctl.Up(opts...)
	if err != nil {
		return nil, err
	}
	return ctl, nil
}

func (z *Single) NewQueue() rc_worker.Queue {
	return rc_worker_impl.NewQueue(z, z.opts.Concurrency)
}

func (z *Single) Catalogue() []rc_recipe.Recipe {
	return z.catalogue
}

func (z *Single) Template() app_template.Template {
	return app_template_impl.NewDev(z.HttpFileSystem(), z)
}

func (z *Single) HttpFileSystem() http.FileSystem {
	return z.web.HTTPBox()
}

func (z *Single) IsProduction() bool {
	return isProduction()
}

func (z *Single) IsSecure() bool {
	return z.opts.Secure
}

func (z *Single) IsQuiet() bool {
	return z.quiet
}

func (z *Single) IsTest() bool {
	return false
}

func (z *Single) Workspace() app_workspace.Workspace {
	return z.ws
}

func (z *Single) Resource(key string) (bin []byte, err error) {
	return z.box.Bytes(key)
}

func (z *Single) upWithWorkspace(ws app_workspace.Workspace) (err error) {
	rl, err := os.Create(filepath.Join(z.ws.Log(), "recipe.log"))
	if err != nil {
		return err
	}
	type RecipeLog struct {
		Name        string                 `json:"name"`
		ValueObject map[string]interface{} `json:"value_object"`
		CommonOpts  map[string]interface{} `json:"common_opts"`
	}
	rr := &RecipeLog{
		Name:        z.opts.RecipeName,
		ValueObject: z.opts.RecipeOptions,
		CommonOpts:  z.opts.CommonOptions,
	}
	rb, err := json.Marshal(rr)
	if err != nil {
		return err
	}
	rl.Write(rb)
	rl.Close()

	z.flc, err = app_log.NewFileLogger(z.ws.Log(), z.opts.Debug)
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

func (z *Single) upWithHome(homePath string) (err error) {
	z.ws, err = app_workspace.NewSingleUser(homePath)
	if err != nil {
		return err
	}
	return z.upWithWorkspace(z.ws)
}

func (z *Single) Up(opts ...app_control.UpOpt) (err error) {
	z.opts = &app_control.UpOpts{}
	for _, o := range opts {
		o(z.opts)
	}
	return z.upWithHome(z.opts.WorkspacePath)
}

func (z *Single) Down() {
	z.Log().Debug("Down")
	app_root.Flush()
	z.cap.Close()
	z.flc.Close()
}

func (z *Single) Abort(opts ...app_control.AbortOpt) {
	opt := &app_control.AbortOpts{}
	for _, o := range opts {
		o(opt)
	}
	z.Log().Debug("Abort shutdown", zap.Any("opt", opt))
	app_root.Flush()
	z.cap.Close()
	z.flc.Close()

	if opt.Reason == nil {
		os.Exit(app_control.FatalGeneral)
	} else {
		os.Exit(*opt.Reason)
	}
}

func (z *Single) UI() app_ui.UI {
	return z.ui
}

func (z *Single) Log() *zap.Logger {
	return z.flc.Logger
}

func (z *Single) Capture() *zap.Logger {
	return z.cap.Logger
}

func (z *Single) TestResource(key string) (data gjson.Result, found bool) {
	data = z.testResource.Get(key)
	return data, data.Exists()
}
