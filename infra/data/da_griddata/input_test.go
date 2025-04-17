package da_griddata

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type mockLogger struct{ esl.Logger }

func (m *mockLogger) With(fields ...esl.Field) esl.Logger   { return m }
func (m *mockLogger) AddCallerSkip(n int) esl.Logger        { return m }
func (m *mockLogger) Debug(msg string, fields ...esl.Field) {}
func (m *mockLogger) Info(msg string, fields ...esl.Field)  {}
func (m *mockLogger) Warn(msg string, fields ...esl.Field)  {}
func (m *mockLogger) Error(msg string, fields ...esl.Field) {}
func (m *mockLogger) Sync() error                           { return nil }

type mockUI struct{ app_ui.UI }

func (m *mockUI) Error(_ app_msg.Message) {}

// Only implement the methods used in the test
func (m *mockUI) Text(_ app_msg.Message) string { return "" }

// mockControl implements only the methods used by gdInput
// (UI and Log)
type mockControl struct {
	app_control.Control
}

func (m *mockControl) UI() app_ui.UI   { return &mockUI{} }
func (m *mockControl) Log() esl.Logger { return &mockLogger{} }

func TestNewInput_SetFilePathAndFilePath(t *testing.T) {
	input := NewInput(nil, "testInput")
	input.SetFilePath("/tmp/test.csv")
	assert.Equal(t, "/tmp/test.csv", input.FilePath())
}

func TestNewInput_Debug(t *testing.T) {
	input := NewInput(nil, "testInput")
	input.SetFilePath("/tmp/test.csv")
	dbg := input.Debug()
	m, ok := dbg.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "testInput", m["Name"])
	assert.Equal(t, "/tmp/test.csv", m["FilePath"])
}

func TestGdInput_Open_FileNotFound(t *testing.T) {
	input := NewInput(nil, "testInput")
	input.SetFilePath("/notfound/file.csv")
	err := input.Open(&mockControl{})
	assert.NotNil(t, err)
	assert.True(t, os.IsNotExist(err))
}

func TestGdInput_Open_NotAFile(t *testing.T) {
	// This test is skipped as it requires a directory path, which is environment dependent.
}
