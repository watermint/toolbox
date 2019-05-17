package app_root

import (
	"github.com/watermint/toolbox/app86/app_control"
	"go.uber.org/zap"
)

var (
	Root = app_control.NewMock()
)

func Log() *zap.Logger {
	return Root.Log()
}
