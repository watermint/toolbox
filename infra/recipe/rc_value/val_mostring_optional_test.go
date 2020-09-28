package rc_value

import (
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueMoStringOptionalRecipe struct {
	LeaveEmpty  mo_string.OptionalString
	FromArg     mo_string.OptionalString
	UpdateByArg mo_string.OptionalString
}

func (z *ValueMoStringOptionalRecipe) Preset() {
	z.UpdateByArg = mo_string.NewOptional("orange")
}

func (z *ValueMoStringOptionalRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueMoStringOptionalRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueMoStringOptional(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueMoStringOptionalRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-from-arg", "apple", "-update-by-arg", "grape"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueMoStringOptionalRecipe)
		if mod1.LeaveEmpty.IsExists() || mod1.FromArg.Value() != "apple" || mod1.UpdateByArg.Value() != "grape" {
			t.Error(mod1)
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueMoStringOptionalRecipe)
		if mod2.LeaveEmpty.IsExists() || mod2.FromArg.Value() != "apple" || mod2.UpdateByArg.Value() != "grape" {
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

func TestValueMoStringOptional_Capture(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		v := newValueOptionalString()
		vb := v.Bind().(*string)
		*vb = "ginger ale"

		vc, err := v.Capture(ctl)
		if err != nil {
			t.Error(err)
		}

		capData, err := json.Marshal(vc)
		if err != nil {
			t.Error(err)
		}

		capJson, err := es_json.Parse(capData)
		if err != nil {
			t.Error(err)
		}

		v2 := newValueOptionalString()

		err = v2.Restore(capJson, ctl)
		if err != nil {
			t.Error(err)
		}

		v2b := v2.Bind().(*string)
		if *v2b != "ginger ale" {
			t.Error(v2b)
		}
		return err
	})
	if err != nil {
		t.Error(err)
	}
}
