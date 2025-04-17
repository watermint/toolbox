package da_json_test

import (
	"testing"

	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_json"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestNewJsonSpec(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		type S struct {
			Name  string `json:"name" path:"name"`
			Price int    `json:"price" path:"price"`
		}
		type K struct {
			StockId string `json:"stock_id" path:"stock_id"`
			Stock   []S    `json:"stock" path:"stock"`
		}

		// S
		{
			// single object
			content := `{"name":"apple", "price": 3}`
			qt_file.TestWithTestFile(t, "json", content, func(path string) {
				v := da_json.NewInput("s", &S{})
				v.SetFilePath(path)
				v.SetModel(&S{})
				if err := v.Open(ctl); err != nil {
					t.Error(err)
					return
				}

				lineNum := 0
				emErr := v.EachModel(func(m interface{}) error {
					c := m.(*S)
					if c.Name != "apple" || c.Price != 3 {
						t.Error(c.Name, c.Price)
					}
					lineNum++
					return nil
				})
				if emErr != nil {
					t.Error(emErr)
				}
				if lineNum != 1 {
					t.Error(lineNum)
				}
			})

			// single object multi lines
			content = `{
"name":"apple",
"price": 3
}`
			qt_file.TestWithTestFile(t, "json", content, func(path string) {
				v := da_json.NewInput("s", &S{})
				v.SetFilePath(path)
				v.SetModel(&S{})
				if err := v.Open(ctl); err != nil {
					t.Error(err)
					return
				}

				lineNum := 0
				emErr := v.EachModel(func(m interface{}) error {
					c := m.(*S)
					if c.Name != "apple" || c.Price != 3 {
						t.Error(c.Name, c.Price)
					}
					lineNum++
					return nil
				})
				if emErr != nil {
					t.Error(emErr)
				}
				if lineNum != 1 {
					t.Error(lineNum)
				}
			})

			// multi object array
			content = `[
	{
	  "name":"apple",
	  "price": 3
	},
	{
	  "name":"banana",
	  "price": 6
	}
]`
			qt_file.TestWithTestFile(t, "json", content, func(path string) {
				v := da_json.NewInput("s", &S{})
				v.SetFilePath(path)
				v.SetModel(&S{})
				if err := v.Open(ctl); err != nil {
					t.Error(err)
					return
				}

				lineNum := 0
				emErr := v.EachModel(func(m interface{}) error {
					c := m.(*S)
					switch lineNum {
					case 0:
						if c.Name != "apple" || c.Price != 3 {
							t.Error(c.Name, c.Price)
						}
					case 1:
						if c.Name != "banana" || c.Price != 6 {
							t.Error(c.Name, c.Price)
						}
					default:
						t.Error(lineNum)
					}
					lineNum++
					return nil
				})
				if emErr != nil {
					t.Error(emErr)
				}
				if lineNum != 2 {
					t.Error(lineNum)
				}
			})

			// JSON lines
			content = `
{"name":"apple", "price": 3}
{"name":"banana", "price": 6}
`
			qt_file.TestWithTestFile(t, "json", content, func(path string) {
				v := da_json.NewInput("s", &S{})
				v.SetFilePath(path)
				v.SetModel(&S{})
				if err := v.Open(ctl); err != nil {
					t.Error(err)
					return
				}

				lineNum := 0
				emErr := v.EachModel(func(m interface{}) error {
					c := m.(*S)
					switch lineNum {
					case 0:
						if c.Name != "apple" || c.Price != 3 {
							t.Error(c.Name, c.Price)
						}
					case 1:
						if c.Name != "banana" || c.Price != 6 {
							t.Error(c.Name, c.Price)
						}
					default:
						t.Error(lineNum)
					}
					lineNum++
					return nil
				})
				if emErr != nil {
					t.Error(emErr)
				}
				if lineNum != 2 {
					t.Error(lineNum)
				}
			})
		}
	})
}

// --- Unit tests for da_json/spec.go ---

type mockUI struct{}

func (m *mockUI) Syntax()                                                          {}
func (m *mockUI) Exists(app_msg.Message) bool                                      { return true }
func (m *mockUI) WithTable(string, func(app_ui.Table))                             {}
func (m *mockUI) Text(msg app_msg.Message) string                                  { return "desc-text" }
func (m *mockUI) TextOrEmpty(msg app_msg.Message) string                           { return "desc-text" }
func (m *mockUI) Id() string                                                       { return "mock" }
func (m *mockUI) WithContainer(mc app_msg_container.Container) app_ui.UI           { return m }
func (m *mockUI) AskCont(app_msg.Message) bool                                     { return true }
func (m *mockUI) AskProceed(app_msg.Message)                                       {}
func (m *mockUI) AskSecure(app_msg.Message) (string, bool)                         { return "", false }
func (m *mockUI) AskText(app_msg.Message) (string, bool)                           { return "", false }
func (m *mockUI) Break()                                                           {}
func (m *mockUI) Code(string)                                                      {}
func (m *mockUI) Error(app_msg.Message)                                            {}
func (m *mockUI) Failure(app_msg.Message)                                          {}
func (m *mockUI) Header(app_msg.Message)                                           {}
func (m *mockUI) SubHeader(app_msg.Message)                                        {}
func (m *mockUI) Info(app_msg.Message)                                             {}
func (m *mockUI) InfoTable(string) app_ui.Table                                    { return nil }
func (m *mockUI) Success(app_msg.Message)                                          {}
func (m *mockUI) Progress(app_msg.Message)                                         {}
func (m *mockUI) Quote(app_msg.Message)                                            {}
func (m *mockUI) Link(rp_artifact.Artifact)                                        {}
func (m *mockUI) IsConsole() bool                                                  { return false }
func (m *mockUI) IsWeb() bool                                                      { return false }
func (m *mockUI) Messages() app_msg_container.Container                            { return nil }
func (m *mockUI) WithContainerSyntax(mc app_msg_container.Container) app_ui.Syntax { return m }

func TestJsonSpec_NameDescDoc(t *testing.T) {
	type dummy struct{}
	spec := da_json.NewJsonSpec("TestName", &dummy{})

	if spec.Name() != "TestName" {
		t.Errorf("expected Name to be 'TestName', got '%s'", spec.Name())
	}

	desc := spec.Desc()
	if desc == nil || desc.Key() == "" {
		t.Error("expected Desc to have a key")
	}

	ui := &mockUI{}
	doc := spec.Doc(ui)
	if doc.Name != "TestName" {
		t.Errorf("expected Doc.Name to be 'TestName', got '%s'", doc.Name)
	}
	if doc.Desc != "desc-text" {
		t.Errorf("expected Doc.Desc to be 'desc-text', got '%s'", doc.Desc)
	}
}
