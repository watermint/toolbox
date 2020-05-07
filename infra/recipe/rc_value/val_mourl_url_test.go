package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_url"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueMoUrlUrlRecipe struct {
	LeaveDefault mo_url.Url
	UpdateByArg  mo_url.Url
}

func (z *ValueMoUrlUrlRecipe) Preset() {
	u, err := mo_url.NewUrl("https://www.dropbox.com")
	if err != nil {
		panic(err)
	}
	z.LeaveDefault = u
}

func (z *ValueMoUrlUrlRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueMoUrlUrlRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueMoUrlUrlSuccess(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueMoUrlUrlRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-update-by-arg", "https://www.dropbox.com/business"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueMoUrlUrlRecipe)
		if mod1.LeaveDefault.Value() != "https://www.dropbox.com" {
			t.Error(mod1)
		}
		if mod1.UpdateByArg.Value() != "https://www.dropbox.com/business" {
			t.Error(mod1)
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueMoUrlUrlRecipe)
		if mod2.LeaveDefault.Value() != "https://www.dropbox.com" {
			t.Error(mod2)
		}
		if mod2.UpdateByArg.Value() != "https://www.dropbox.com/business" {
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

func TestValueMoUrlUrlMissing(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueMoUrlUrlRecipe{}
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
		if err != ErrorMissingRequiredOption {
			t.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
