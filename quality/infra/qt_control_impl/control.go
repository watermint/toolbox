package qt_control_impl

import (
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"strings"
	"testing"
)

func NewMessageMemory() qt_control.Message {
	return &messageMemory{
		missing: make(map[string]bool),
	}
}

func NewMessageTest(t *testing.T) qt_control.Message {
	return &messageTest{
		t:       t,
		missing: make(map[string]bool),
	}
}

type messageTest struct {
	missing map[string]bool
	t       *testing.T
	n       int
}

func (z *messageTest) Missing() []string {
	m := make([]string, 0)
	for k := range z.missing {
		m = append(m, k)
	}
	return m
}

func (z *messageTest) NotFound(key string) {
	z.n++
	if strings.HasSuffix(key, ".peer") {
		z.t.Errorf("[Message not found, but suggested value] \"%s\":\"%s\"", key, "Account alias")
	} else {
		z.t.Errorf("[Message not found for key] \"%s\":\"\"", key)
	}
}

type messageMemory struct {
	missing map[string]bool
}

func (z *messageMemory) Missing() []string {
	m := make([]string, 0)
	for k := range z.missing {
		m = append(m, k)
	}
	return m
}

func (z *messageMemory) NotFound(key string) {
	z.missing[key] = true
}
