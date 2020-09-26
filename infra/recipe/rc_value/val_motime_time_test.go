package rc_value

import (
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueMoTimeTimeRecipe struct {
	UpdateByArg          mo_time.Time
	LeaveDefaultOptional mo_time.TimeOptional
	UpdateByArgOptional  mo_time.TimeOptional
}

func (z *ValueMoTimeTimeRecipe) Preset() {
}

func (z *ValueMoTimeTimeRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueMoTimeTimeRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueMoTimeTimeSuccess(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueMoTimeTimeRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		args := []string{
			"-update-by-arg", "2020-05-07T17:10:55Z",
			"-update-by-arg-optional", "2020-05-07T17:20:55Z",
		}
		if err := flg.Parse(args); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		_ = repo.Apply()

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueMoTimeTimeRecipe)
		if mod2.UpdateByArg.Iso8601() != "2020-05-07T17:10:55Z" {
			t.Error(mod2)
		}
		if mod2.UpdateByArgOptional.Iso8601() != "2020-05-07T17:20:55Z" {
			t.Error(mod2)
		}
		if !mod2.LeaveDefaultOptional.IsZero() {
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

func TestValueMoTimeTimeMissing(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueMoTimeTimeRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		args := []string{}
		if err := flg.Parse(args); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		_ = repo.Apply()

		// Spin up
		_, err := repo.SpinUp(c)
		if err == nil {
			t.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestValueMoTimeTime_Capture(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		v := newValueMoTimeTime("mo")
		vb := v.Bind().(*string)
		*vb = "2020-09-26"
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

		v2 := newValueMoTimeTime("mo")

		err = v2.Restore(capJson, ctl)
		if err != nil {
			t.Error(err)
		}

		v2b := v2.Bind().(*string)
		if *v2b != "2020-09-26" {
			t.Error(v2b)
		}
		return err
	})
	if err != nil {
		t.Error(err)
	}
}
