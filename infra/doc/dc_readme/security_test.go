package dc_readme

import (
	"testing"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
)

func TestNewSecurity(t *testing.T) {
	s := NewSecurity()
	if s == nil {
		t.Error("Expected non-nil security document")
	}
	
	sec, ok := s.(*docSecurity)
	if !ok {
		t.Error("Expected docSecurity type")
	}
	if sec == nil {
		t.Error("Expected non-nil security instance")
	}
}

func TestDocSecurity_DocId(t *testing.T) {
	s := &docSecurity{}
	
	id := s.DocId()
	if id != dc_index.DocRootSecurityAndPrivacy {
		t.Errorf("Expected DocId to be DocRootSecurityAndPrivacy, got %v", id)
	}
}

func TestDocSecurity_DocDesc(t *testing.T) {
	s := &docSecurity{
		Desc: app_msg.Raw("Security description"),
	}
	
	desc := s.DocDesc()
	if desc == nil {
		t.Error("Expected non-nil description")
	}
}

func TestDocSecurity_Sections(t *testing.T) {
	s := &docSecurity{}
	
	sections := s.Sections()
	if len(sections) != 1 {
		t.Errorf("Expected 1 section, got %d", len(sections))
	}
}

func TestNewSecuritySection(t *testing.T) {
	s := NewSecuritySection()
	if s == nil {
		t.Error("Expected non-nil security section")
	}
	
	sec, ok := s.(*SecurityDesc)
	if !ok {
		t.Error("Expected SecurityDesc type")
	}
	if sec == nil {
		t.Error("Expected non-nil security desc instance")
	}
}

func TestSecurityDesc_Title(t *testing.T) {
	s := &SecurityDesc{
		HeaderTitle: app_msg.Raw("Security Title"),
	}
	
	title := s.Title()
	if title == nil {
		t.Error("Expected non-nil title")
	}
}

func TestSecurityDesc_Body(t *testing.T) {
	s := &SecurityDesc{
		BodyOverview:         app_msg.Raw("Overview"),
		HeaderDataProtection: app_msg.Raw("Data Protection"),
		BodyDataProtection:   app_msg.Raw("Data protection body"),
		HeaderUse:            app_msg.Raw("Use"),
		BodyUse:              app_msg.Raw("Use body"),
		HeaderSharing:        app_msg.Raw("Sharing"),
		BodySharing:          app_msg.Raw("Sharing body"),
	}
	
	// Create a mock UI to test Body method
	mockUI := &mockSecurityUI{
		infoCount:      0,
		subHeaderCount: 0,
	}
	
	s.Body(mockUI)
	
	// Verify the expected calls were made
	if mockUI.infoCount != 4 { // 1 overview + 3 body sections
		t.Errorf("Expected 4 Info calls, got %d", mockUI.infoCount)
	}
	if mockUI.subHeaderCount != 3 {
		t.Errorf("Expected 3 SubHeader calls, got %d", mockUI.subHeaderCount)
	}
}

// Mock UI for security testing
type mockSecurityUI struct {
	infoCount      int
	subHeaderCount int
}

func (m *mockSecurityUI) Info(msg app_msg.Message) {
	m.infoCount++
}

func (m *mockSecurityUI) SubHeader(msg app_msg.Message) {
	m.subHeaderCount++
}

// Add other required methods to satisfy app_ui.UI interface
func (m *mockSecurityUI) Ask(msg app_msg.Message, defaultValue string) string { return "" }
func (m *mockSecurityUI) AskCont(msg app_msg.Message) bool { return true }
func (m *mockSecurityUI) AskProceed(msg app_msg.Message) {}
func (m *mockSecurityUI) AskSecure(msg app_msg.Message) (string, bool) { return "", false }
func (m *mockSecurityUI) AskText(msg app_msg.Message) (string, bool) { return "", false }
func (m *mockSecurityUI) Code(code string) {}
func (m *mockSecurityUI) Break() {}
func (m *mockSecurityUI) Error(msg app_msg.Message) {}
func (m *mockSecurityUI) Exists(msg app_msg.Message) bool { return false }
func (m *mockSecurityUI) Header(msg app_msg.Message) {}
func (m *mockSecurityUI) IsConsoleUI() bool { return true }
func (m *mockSecurityUI) ItemOf(msg app_msg.Message, id string) {}
func (m *mockSecurityUI) KeyValue(key, value string) {}
func (m *mockSecurityUI) ProgressStart(count int) {}
func (m *mockSecurityUI) ProgressUpdate(done int) {}
func (m *mockSecurityUI) ProgressEnd() {}
func (m *mockSecurityUI) Quote(msg app_msg.Message) {}
func (m *mockSecurityUI) SubInfo(msg app_msg.Message) {}
func (m *mockSecurityUI) Success(msg app_msg.Message) {}
func (m *mockSecurityUI) Text(msg app_msg.Message) string { return "" }
func (m *mockSecurityUI) TextOrEmpty(msg app_msg.Message) string { return "" }
func (m *mockSecurityUI) Translate(text app_msg.Message) string { return "" }
func (m *mockSecurityUI) TreePut(path []string, name string, value app_msg.MessageOptional) {}
func (m *mockSecurityUI) TreeShow() {}
func (m *mockSecurityUI) Warn(msg app_msg.Message) {}
func (m *mockSecurityUI) InfoTable(name string) app_ui.Table { return nil }
func (m *mockSecurityUI) Failure(msg app_msg.Message) {}
func (m *mockSecurityUI) Progress(msg app_msg.Message) {}
func (m *mockSecurityUI) DefinitionList(definitions []app_ui.Definition) {}
func (m *mockSecurityUI) Link(artifact rp_artifact.Artifact) {}
func (m *mockSecurityUI) IsConsole() bool { return true }
func (m *mockSecurityUI) IsWeb() bool { return false }
func (m *mockSecurityUI) WithContainerSyntax(mc app_msg_container.Container) app_ui.Syntax { return m }
func (m *mockSecurityUI) Messages() app_msg_container.Container { return nil }
func (m *mockSecurityUI) WithTable(name string, f func(t app_ui.Table)) {}
func (m *mockSecurityUI) Id() string { return "mock" }
func (m *mockSecurityUI) WithContainer(mc app_msg_container.Container) app_ui.UI { return m }