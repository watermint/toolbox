package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueRpModelRowReportRecipeData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ValueRpModelRowReportRecipe struct {
	Member rp_model.RowReport
}

func (z *ValueRpModelRowReportRecipe) Preset() {
	z.Member.SetModel(&ValueRpModelRowReportRecipeData{})
}

func (z *ValueRpModelRowReportRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueRpModelRowReportRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueRpModelRowReport(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueRpModelRowReportRecipe{}
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
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueRpModelRowReportRecipe)

		if err := mod2.Member.Open(); err != nil {
			t.Error(err)
		}
		mod2.Member.Row(&ValueRpModelRowReportRecipeData{
			Email: "john@example.com",
			Name:  "John",
		})

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
