package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"testing"
)

type ValueDaTextInputRecipe struct {
	Text da_text.TextInput
}

func (z *ValueDaTextInputRecipe) Preset() {
}

func (z *ValueDaTextInputRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueDaTextInputRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueDaTextInput_Accept(t *testing.T) {
	testData, err := qt_file.MakeTestFile("text", `Lorem ipsum dolor sit amet
consectetur adipiscing elit
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex`)

	if err != nil {
		return
	}
	defer func() {
		_ = os.Remove(testData)
	}()

	err = qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueDaTextInputRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-text", testData}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueDaTextInputRecipe)
		if x := mod1.Text.FilePath(); x != testData {
			t.Error(x)
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueDaTextInputRecipe)
		if x := mod2.Text.FilePath(); x != testData {
			t.Error(x)
		}

		if err := repo.SpinDown(c); err != nil {
			t.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
