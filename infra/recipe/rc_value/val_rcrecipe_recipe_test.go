package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueRcRecipeRecipeRecipe struct {
	SubRecipe *ValueBoolRecipe
}

func (z *ValueRcRecipeRecipeRecipe) Preset() {
}

func (z *ValueRcRecipeRecipeRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueRcRecipeRecipeRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueRcRecipeRecipe(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueRcRecipeRecipeRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		_ = repo.Apply()

		// Spin up
		_, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
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
