package context

import (
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
)

func NewMock() Context {
	return &mockImpl{}
}

type mockImpl struct {
}

func (z mockImpl) ClientHash() string {
	return ""
}

func (z mockImpl) Log() *zap.Logger {
	return app_root.Log()
}

func (z mockImpl) Capture() *zap.Logger {
	return app_root.Capture()
}
