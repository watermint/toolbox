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

type ValueMoStringSelectRecipe struct {
	LeaveDefault mo_string.SelectString
	UpdateByArg  mo_string.SelectString
}

func (z *ValueMoStringSelectRecipe) Preset() {
	z.LeaveDefault.SetOptions("orange", "orange", "apple", "grape")
	z.UpdateByArg.SetOptions("soba", "soba", "udon", "somen")
}

func (z *ValueMoStringSelectRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueMoStringSelectRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueMoStringSelectSuccess(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueMoStringSelectRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-update-by-arg", "udon"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueMoStringSelectRecipe)
		if mod1.LeaveDefault.Value() != "orange" || mod1.UpdateByArg.Value() != "udon" {
			t.Error(mod1)
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueMoStringSelectRecipe)
		if mod2.LeaveDefault.Value() != "orange" || mod2.UpdateByArg.Value() != "udon" {
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

func TestValueMoStringSelectInvalid(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueMoStringSelectRecipe{}
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
		mod1 := rcp1.(*ValueMoStringSelectRecipe)
		if mod1.LeaveDefault.Value() != "orange" || mod1.UpdateByArg.Value() != "ramen" {
			t.Error(mod1)
		}

		// Spin up
		_, err := repo.SpinUp(c)
		if err != ErrorInvalidValue {
			t.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestValueMoStringSelect_Capture(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		v := newValueSelectString()
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

		v2 := newValueSelectString()

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
