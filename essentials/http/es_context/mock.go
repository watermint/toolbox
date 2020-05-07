package es_context

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
)

func NewMock() Context {
	return &mockImpl{}
}

type mockImpl struct {
}

func (z mockImpl) ClientHash() string {
	return ""
}

func (z mockImpl) Log() es_log.Logger {
	return es_log.Default()
}

func (z mockImpl) Capture() es_log.Logger {
	return es_log.Capture()
}
