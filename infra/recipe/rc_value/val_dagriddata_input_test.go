package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"testing"
)

type ValueDaGridDataInputRecipe struct {
	GridData da_griddata.GridDataInput
}

func (z *ValueDaGridDataInputRecipe) Preset() {
}

func (z *ValueDaGridDataInputRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueDaGridDataInputRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueDaGridDataInput_Accept(t *testing.T) {
	testData, err := qt_file.MakeTestFile("griddata", "alex@example.com,Alex\ndavid@example.com,David\n")
	if err != nil {
		return
	}
	defer func() {
		_ = os.Remove(testData)
	}()

	err = qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueDaGridDataInputRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-grid-data", testData}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueDaGridDataInputRecipe)
		if x := mod1.GridData.FilePath(); x != testData {
			t.Error(x)
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueDaGridDataInputRecipe)
		if x := mod2.GridData.FilePath(); x != testData {
			t.Error(x)
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
