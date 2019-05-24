package app_control

import (
	"github.com/watermint/toolbox/app86/app_ui"
	"go.uber.org/zap"
)

type Control interface {
	Startup(opts ...StartupOpt) error
	Shutdown()
	Fatal(opts ...FatalOpt)
	UI() app_ui.UI
	Log() *zap.Logger
	Resource(key string) (bin []byte, err error)
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

const (
	FatalMock = iota + 1
)
