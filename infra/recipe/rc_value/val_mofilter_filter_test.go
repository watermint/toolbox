package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueMoFilterFilterRecipe struct {
	Filter mo_filter.Filter
}

func (z *ValueMoFilterFilterRecipe) Preset() {
	z.Filter.SetOptions(mo_filter.NewNameFilter())
}

func (z *ValueMoFilterFilterRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueMoFilterFilterRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueMoFilterFilterSuccess(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueMoFilterFilterRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-filter-name", "watermint"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueMoFilterFilterRecipe)
		if !mod1.Filter.Accept("watermint") {
			t.Error("does not accept name")
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueMoFilterFilterRecipe)
		if !mod2.Filter.Accept("watermint") {
			t.Error("does not accept name")
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
