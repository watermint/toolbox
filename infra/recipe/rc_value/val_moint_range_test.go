package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueMoIntRangeRecipe struct {
	DefaultValue mo_int.RangeInt
	UpdateByArg  mo_int.RangeInt
}

func (z *ValueMoIntRangeRecipe) Preset() {
	z.DefaultValue.SetRange(100, 200, 123)
	z.UpdateByArg.SetRange(1000, 2000, 1234)
}

func (z *ValueMoIntRangeRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueMoIntRangeRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueMoIntRangeSuccess(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueMoIntRangeRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-update-by-arg", "1999"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueMoIntRangeRecipe)
		if mod1.DefaultValue.IsValid() && mod1.DefaultValue.Value() != 123 {
			t.Error(mod1)
		}
		if mod1.UpdateByArg.IsValid() && mod1.UpdateByArg.Value() != 1999 {
			t.Error(mod1)
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueMoIntRangeRecipe)
		if mod2.DefaultValue.IsValid() && mod2.DefaultValue.Value() != 123 {
			t.Error(mod2)
		}
		if mod2.UpdateByArg.IsValid() && mod2.UpdateByArg.Value() != 1999 {
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

func TestValueMoIntRangeOutOfRange(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueMoIntRangeRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-update-by-arg", "9999"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueMoIntRangeRecipe)
		if mod1.DefaultValue.IsValid() && mod1.DefaultValue.Value() != 123 {
			t.Error(mod1)
		}
		if mod1.UpdateByArg.IsValid() && mod1.UpdateByArg.Value() != 9999 {
			t.Error(mod1)
		}

		// Spin up
		_, err := repo.SpinUp(c)
		if err == nil || err != ErrorInvalidValue {
			t.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
