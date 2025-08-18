package dc_readme

import (
    "testing"

    "github.com/watermint/toolbox/infra/report/rp_artifact"
    "github.com/watermint/toolbox/infra/ui/app_msg"
    "github.com/watermint/toolbox/infra/ui/app_msg_container"
    "github.com/watermint/toolbox/infra/ui/app_ui"
)

func TestNewLicense(t *testing.T) {
	l := NewLicense()
	if l == nil {
		t.Error("Expected non-nil license")
	}

	license, ok := l.(*License)
	if !ok {
		t.Error("Expected License type")
	}
	if license == nil {
		t.Error("Expected non-nil license instance")
	}
}

func TestLicense_Title(t *testing.T) {
	l := &License{
		HeaderTitle: app_msg.Raw("License Title"),
	}

	title := l.Title()
	if title == nil {
		t.Error("Expected non-nil title")
	}
}

func TestLicense_Body(t *testing.T) {
	l := &License{
		BodyLicense:        app_msg.Raw("License body"),
		BodyLicenseRemarks: app_msg.Raw("License remarks"),
		BodyLicenseQuote:   app_msg.Raw("License quote"),
	}

	// Create a mock UI to test Body method
	mockUI := &mockLicenseUI{
		infoCount:  0,
		breakCount: 0,
		quoteCount: 0,
	}

	l.Body(mockUI)

	// Verify the expected calls were made
	if mockUI.infoCount != 2 {
		t.Errorf("Expected 2 Info calls, got %d", mockUI.infoCount)
	}
	if mockUI.breakCount != 1 {
		t.Errorf("Expected 1 Break call, got %d", mockUI.breakCount)
	}
	if mockUI.quoteCount != 1 {
		t.Errorf("Expected 1 Quote call, got %d", mockUI.quoteCount)
	}
}

// Mock UI for license testing
type mockLicenseUI struct {
	infoCount  int
	breakCount int
	quoteCount int
}

func (m *mockLicenseUI) Info(msg app_msg.Message) {
	m.infoCount++
}

func (m *mockLicenseUI) Break() {
	m.breakCount++
}

func (m *mockLicenseUI) Quote(msg app_msg.Message) {
	m.quoteCount++
}

// Add other required methods to satisfy app_ui.UI interface
func (m *mockLicenseUI) Ask(msg app_msg.Message, defaultValue string) string               { return "" }
func (m *mockLicenseUI) AskCont(msg app_msg.Message) bool                                  { return true }
func (m *mockLicenseUI) AskProceed(msg app_msg.Message)                                    {}
func (m *mockLicenseUI) AskSecure(msg app_msg.Message) (string, bool)                      { return "", false }
func (m *mockLicenseUI) AskText(msg app_msg.Message) (string, bool)                        { return "", false }
func (m *mockLicenseUI) Code(code string)                                                  {}
func (m *mockLicenseUI) Error(msg app_msg.Message)                                         {}
func (m *mockLicenseUI) Exists(msg app_msg.Message) bool                                   { return false }
func (m *mockLicenseUI) Header(msg app_msg.Message)                                        {}
func (m *mockLicenseUI) IsConsoleUI() bool                                                 { return true }
func (m *mockLicenseUI) ItemOf(msg app_msg.Message, id string)                             {}
func (m *mockLicenseUI) KeyValue(key, value string)                                        {}
func (m *mockLicenseUI) ProgressStart(count int)                                           {}
func (m *mockLicenseUI) ProgressUpdate(done int)                                           {}
func (m *mockLicenseUI) ProgressEnd()                                                      {}
func (m *mockLicenseUI) SubInfo(msg app_msg.Message)                                       {}
func (m *mockLicenseUI) Success(msg app_msg.Message)                                       {}
func (m *mockLicenseUI) Text(msg app_msg.Message) string                                   { return "" }
func (m *mockLicenseUI) TextOrEmpty(msg app_msg.Message) string                            { return "" }
func (m *mockLicenseUI) Translate(text app_msg.Message) string                             { return "" }
func (m *mockLicenseUI) TreePut(path []string, name string, value app_msg.MessageOptional) {}
func (m *mockLicenseUI) TreeShow()                                                         {}
func (m *mockLicenseUI) Warn(msg app_msg.Message)                                          {}
func (m *mockLicenseUI) SubHeader(msg app_msg.Message)                                     {}
func (m *mockLicenseUI) InfoTable(name string) app_ui.Table                                { return nil }
func (m *mockLicenseUI) Failure(msg app_msg.Message)                                       {}
func (m *mockLicenseUI) Progress(msg app_msg.Message)                                      {}
func (m *mockLicenseUI) DefinitionList(definitions []app_ui.Definition)                    {}
func (m *mockLicenseUI) Link(artifact rp_artifact.Artifact)                                {}
func (m *mockLicenseUI) IsConsole() bool                                                   { return true }
func (m *mockLicenseUI) IsWeb() bool                                                       { return false }
func (m *mockLicenseUI) WithContainerSyntax(mc app_msg_container.Container) app_ui.Syntax  { return m }
func (m *mockLicenseUI) Messages() app_msg_container.Container                             { return nil }
func (m *mockLicenseUI) WithTable(name string, f func(t app_ui.Table))                     {}
func (m *mockLicenseUI) Id() string                                                        { return "mock" }
func (m *mockLicenseUI) WithContainer(mc app_msg_container.Container) app_ui.UI            { return m }
