package qt_control_impl

import (
	"github.com/watermint/toolbox/infra/quality/qt_control"
	"testing"
)

func NewMessageMock() qt_control.Message {
	return &messageMock{}
}

func NewMessageTest(t *testing.T) qt_control.Message {
	return &messageTest{t: t}
}

type messageTest struct {
	t *testing.T
}

func (z *messageTest) NotFound(key string) {
	z.t.Errorf("Message not found for key [%s]", key)
}

type messageMock struct {
}

func (z *messageMock) NotFound(key string) {
	// nop
}
