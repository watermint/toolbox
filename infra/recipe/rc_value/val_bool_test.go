package rc_value

import (
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueBoolRecipe struct {
	Enabled        bool
	EnabledByFlag  bool
	DisabledByFlag bool
}

func (z *ValueBoolRecipe) Preset() {
	z.Enabled = true
	z.DisabledByFlag = true
}

func (z *ValueBoolRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueBoolRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueBool(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueBoolRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-enabled-by-flag", "-disabled-by-flag=false"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueBoolRecipe)
		if !mod1.Enabled || !mod1.EnabledByFlag || mod1.DisabledByFlag {
			t.Error(mod1)
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueBoolRecipe)
		if !mod2.Enabled || !mod2.EnabledByFlag || mod2.DisabledByFlag {
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

func TestValueBool_Capture(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		v := newValueBool()
		vb := v.Bind().(*bool)
		*vb = true

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

		v2 := newValueBool()

		err = v2.Restore(capJson, ctl)
		if err != nil {
			t.Error(err)
		}

		v2b := v2.Bind().(*bool)
		if !*v2b {
			t.Error(*v2b)
		}
		return err
	})
	if err != nil {
		t.Error(err)
	}
}
