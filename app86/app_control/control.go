package app_control

import "github.com/watermint/toolbox/app86/app_ui"

type Control interface {
	Startup(opts ...StartupOpt) error
	Shutdown()
	Fatal(opts ...FatalOpt)
	UI() app_ui.UI
}

type StartupOpt func(opt *startupOpts) startupOpts
type startupOpts struct {
}

type FatalOpt func(opt *fatalOpts) fatalOpts
type fatalOpts struct {
}

type Workspace interface {
	SecretsPath() string
	JobPath() string
	JobId() string
}
