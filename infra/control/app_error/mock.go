package app_error

import (
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewMock() ErrorReport {
	return mockImpl{}
}

type mockImpl struct {
}

func (z mockImpl) Up(ctl app_control.Control) error {
	return nil
}

func (z mockImpl) Down() {
}

func (z mockImpl) ErrorHandler(err error, mouldId, batchId string, p interface{}) {
}
