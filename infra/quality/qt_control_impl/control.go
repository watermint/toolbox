package qt_control_impl

import (
	"github.com/watermint/toolbox/infra/quality/qt_control"
	"strings"
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
	if strings.HasSuffix(key, ".peer") {
		z.t.Errorf("[Message not found, but suggested value] \"%s\":\"%s\"", key, "Account alias")
	} else {
		z.t.Errorf("[Message not found for key] \"%s\":\"\"", key)
	}
}

type messageMock struct {
}

func (z *messageMock) NotFound(key string) {
	// nop
}
