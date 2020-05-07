package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueStringRecipe struct {
	DefaultValue string
	UpdateByArg  string
}

func (z *ValueStringRecipe) Preset() {
	z.DefaultValue = "soba"
	z.UpdateByArg = "udon"
}

func (z *ValueStringRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueStringRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueString(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueStringRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-update-by-arg", "ramen"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueStringRecipe)
		if mod1.DefaultValue != "soba" || mod1.UpdateByArg != "ramen" {
			t.Error(mod1)
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueStringRecipe)
		if mod2.DefaultValue != "soba" || mod2.UpdateByArg != "ramen" {
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
