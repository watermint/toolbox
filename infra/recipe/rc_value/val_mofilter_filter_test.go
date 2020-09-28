package rc_value

import (
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
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

func TestValueMoFilterFilter_Capture(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		var capJson es_json.Json
		{
			v := newValueMoFilter("fi")
			filter := mo_filter.New("lt")

			fName := mo_filter.NewNameFilter()
			fEmail := mo_filter.NewEmailFilter()
			fPrefix := mo_filter.NewNamePrefixFilter()
			fSuffix := mo_filter.NewNameSuffixFilter()

			filter.SetOptions(fName, fEmail, fPrefix, fSuffix)
			fNameVal := fName.Bind().(*string)
			*fNameVal = "hello"
			fEmailVal := fEmail.Bind().(*string)
			*fEmailVal = "test@example.com"

			v.ApplyPreset(filter)

			vc, err := v.Capture(ctl)
			if err != nil {
				t.Error(err)
			}

			capData, err := json.Marshal(vc)
			if err != nil {
				t.Error(err)
			}

			capJson, err = es_json.Parse(capData)
			if err != nil {
				t.Error(err)
			}
		}

		{
			v2 := newValueMoFilter("fi")
			filter := mo_filter.New("lt")
			fName2 := mo_filter.NewNameFilter()
			fEmail2 := mo_filter.NewEmailFilter()
			fPrefix2 := mo_filter.NewNamePrefixFilter()
			fSuffix2 := mo_filter.NewNameSuffixFilter()
			filter.SetOptions(fName2, fEmail2, fPrefix2, fSuffix2)

			v2.ApplyPreset(filter)
			err := v2.Restore(capJson, ctl)
			if err != nil {
				t.Error(err)
			}

			fName2Val := fName2.Bind().(*string)
			fEmail2Val := fEmail2.Bind().(*string)

			if *fName2Val != "hello" {
				t.Error(*fName2Val)
			}
			if *fEmail2Val != "test@example.com" {
				t.Error(*fEmail2Val)
			}
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
