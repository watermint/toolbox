package app_root

import (
	"github.com/watermint/toolbox/app86/app_control_impl"
	"go.uber.org/zap"
)

var (
	Root = app_control_impl.NewMock()
)

func Log() *zap.Logger {
	return Root.Log()
}
