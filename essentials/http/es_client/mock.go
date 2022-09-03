package es_client

import (
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewMock() Client {
	return &mockImpl{}
}

type mockImpl struct {
}

func (z mockImpl) Name() string {
	return ""
}

func (z mockImpl) ClientHash() string {
	return ""
}

func (z mockImpl) Log() esl.Logger {
	return esl.Default()
}

func (z mockImpl) Capture() esl.Logger {
	return esl.Capture()
}
