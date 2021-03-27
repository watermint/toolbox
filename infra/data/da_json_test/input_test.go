package da_json_test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_json"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
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
