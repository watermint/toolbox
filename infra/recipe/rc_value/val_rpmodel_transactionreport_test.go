package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueRpModelTransactionReportRecipeData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ValueRpModelTransactionReportRecipe struct {
	Member rp_model.TransactionReport
}

func (z *ValueRpModelTransactionReportRecipe) Preset() {
	z.Member.SetModel(&ValueRpModelTransactionReportRecipeData{}, nil)
}

func (z *ValueRpModelTransactionReportRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueRpModelTransactionReportRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueRpModelTransactionReport(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueRpModelTransactionReportRecipe{}
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
		mod2 := rcp2.(*ValueRpModelTransactionReportRecipe)

		if err := mod2.Member.Open(); err != nil {
			t.Error(err)
		}
		mod2.Member.Success(&ValueRpModelTransactionReportRecipeData{
			Email: "john@example.com",
			Name:  "John",
		}, nil)

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
