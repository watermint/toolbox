package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueIntRecipe struct {
	DefaultValue int
	FromArg      int
	UpdateByArg  int
}

func (z *ValueIntRecipe) Preset() {
	z.DefaultValue = 123
	z.UpdateByArg = 345
}

func (z *ValueIntRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueIntRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueInt(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueIntRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-from-arg", "234", "-update-by-arg", "567"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueIntRecipe)
		if mod1.DefaultValue != 123 || mod1.FromArg != 234 || mod1.UpdateByArg != 567 {
			t.Error(mod1)
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueIntRecipe)
		if mod2.DefaultValue != 123 || mod2.FromArg != 234 || mod2.UpdateByArg != 567 {
			t.Error(mod2)
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
