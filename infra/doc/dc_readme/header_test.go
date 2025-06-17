package dc_readme

import (
	"testing"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
)

func TestNewHeader(t *testing.T) {
	// Test creating header for publish
	h1 := NewHeader(true)
	if h1 == nil {
		t.Error("Expected non-nil header")
	}
	header1, ok := h1.(*Header)
	if !ok {
		t.Error("Expected Header type")
	}
	if !header1.publish {
		t.Error("Expected publish to be true")
	}
	
	// Test creating header not for publish
	h2 := NewHeader(false)
	if h2 == nil {
		t.Error("Expected non-nil header")
	}
	header2, ok := h2.(*Header)
	if !ok {
		t.Error("Expected Header type")
	}
	if header2.publish {
		t.Error("Expected publish to be false")
	}
}

func TestHeader_Title(t *testing.T) {
	h := &Header{
		HeaderTitle: app_msg.Raw("Test Title"),
	}
	
	title := h.Title()
	if title == nil {
		t.Error("Expected non-nil title")
	}
}

func TestHeader_Body(t *testing.T) {
	// Test with publish = true
	h1 := &Header{
		publish:    true,
		HeaderBody: app_msg.Raw("Test body"),
	}
	
	// Create a mock UI to test Body method
	mockUI := &mockUI{}
	h1.Body(mockUI)
	
	// Test with publish = false
	h2 := &Header{
		publish:    false,
		HeaderBody: app_msg.Raw("Test body"),
	}
	h2.Body(mockUI)
}

// Mock UI for testing
type mockUI struct {
	infoCalled  int
	breakCalled int
}

func (m *mockUI) Info(msg app_msg.Message) {
	m.infoCalled++
}

func (m *mockUI) Break() {
	m.breakCalled++
}

func (m *mockUI) Quote(msg app_msg.Message) {
	// Implement for testing
}

// Add other required methods to satisfy app_ui.UI interface
func (m *mockUI) Ask(msg app_msg.Message, defaultValue string) string { return "" }
func (m *mockUI) AskCont(msg app_msg.Message) bool { return true }
func (m *mockUI) AskProceed(msg app_msg.Message) {}
func (m *mockUI) AskSecure(msg app_msg.Message) (string, bool) { return "", false }
func (m *mockUI) AskText(msg app_msg.Message) (string, bool) { return "", false }
func (m *mockUI) Code(code string) {}
func (m *mockUI) Error(msg app_msg.Message) {}
func (m *mockUI) Exists(msg app_msg.Message) bool { return false }
func (m *mockUI) Header(msg app_msg.Message) {}
func (m *mockUI) IsConsoleUI() bool { return true }
func (m *mockUI) ItemOf(msg app_msg.Message, id string) {}
func (m *mockUI) KeyValue(key, value string) {}
func (m *mockUI) ProgressStart(count int) {}
func (m *mockUI) ProgressUpdate(done int) {}
func (m *mockUI) ProgressEnd() {}
func (m *mockUI) SubInfo(msg app_msg.Message) {}
func (m *mockUI) Success(msg app_msg.Message) {}
func (m *mockUI) Text(msg app_msg.Message) string { return "" }
func (m *mockUI) TextOrEmpty(msg app_msg.Message) string { return "" }
func (m *mockUI) Translate(text app_msg.Message) string { return "" }
func (m *mockUI) TreePut(path []string, name string, value app_msg.MessageOptional) {}
func (m *mockUI) TreeShow() {}
func (m *mockUI) Warn(msg app_msg.Message) {}
func (m *mockUI) SubHeader(msg app_msg.Message) {}
func (m *mockUI) InfoTable(name string) app_ui.Table { return nil }
func (m *mockUI) Failure(msg app_msg.Message) {}
func (m *mockUI) Progress(msg app_msg.Message) {}
func (m *mockUI) DefinitionList(definitions []app_ui.Definition) {}
func (m *mockUI) Link(artifact rp_artifact.Artifact) {}
func (m *mockUI) IsConsole() bool { return true }
func (m *mockUI) IsWeb() bool { return false }
func (m *mockUI) WithContainerSyntax(mc app_msg_container.Container) app_ui.Syntax { return m }
func (m *mockUI) Messages() app_msg_container.Container { return nil }
func (m *mockUI) WithTable(name string, f func(t app_ui.Table)) {}
func (m *mockUI) Id() string { return "mock" }
func (m *mockUI) WithContainer(mc app_msg_container.Container) app_ui.UI { return m }