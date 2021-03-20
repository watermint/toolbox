package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_json"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"testing"
)

type ValueDaJsonInputSample struct {
	Name  string `json:"name" path:"name"`
	Price int    `json:"price" path:"price"`
}

type ValueDaJsonInputRecipe struct {
	Json da_json.JsonInput
}

func (z *ValueDaJsonInputRecipe) Preset() {
	z.Json.SetModel(&ValueDaJsonInputSample{})
}

func (z *ValueDaJsonInputRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueDaJsonInputRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueDaJsonInput_Accept(t *testing.T) {
	content := `{"name":"apple","price":3}`
	qt_file.TestWithTestFile(t, "json", content, func(path string) {
		wcErr := qt_control.WithControl(func(c app_control.Control) error {
			rcp0 := &ValueDaJsonInputRecipe{}
			repo := NewRepository(rcp0)

			// Parse flags
			flg := flag.NewFlagSet("value", flag.ContinueOnError)
			repo.ApplyFlags(flg, c.UI())
			if err := flg.Parse([]string{"-text", path}); err != nil {
				t.Error(err)
				return err
			}

			// Apply parsed values
			rcp1 := repo.Apply()
			mod1 := rcp1.(*ValueDaJsonInputRecipe)
			if x := mod1.Json.FilePath(); x != path {
				t.Error(x)
			}

			// Spin up
			rcp2, err := repo.SpinUp(c)
			if err != nil {
				t.Error(err)
				return err
			}
			mod2 := rcp2.(*ValueDaJsonInputRecipe)
			if x := mod2.Json.FilePath(); x != path {
				t.Error(x)
			}

			if err := repo.SpinDown(c); err != nil {
				t.Error(err)
				return err
			}
			return nil
		})
		if wcErr != nil {
			t.Error(wcErr)
		}
	})
}
