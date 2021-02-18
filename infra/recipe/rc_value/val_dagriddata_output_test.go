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

type ValueDaGridDataOutputRecipe struct {
	GridData da_griddata.GridDataOutput
}

func (z *ValueDaGridDataOutputRecipe) Preset() {
}

func (z *ValueDaGridDataOutputRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueDaGridDataOutputRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueDaGridDataOutput_Accept(t *testing.T) {
	testData, err := qt_file.MakeTestFile("griddata", "")
	if err != nil {
		return
	}
	defer func() {
		_ = os.Remove(testData)
	}()

	err = qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueDaGridDataOutputRecipe{}
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
		mod1 := rcp1.(*ValueDaGridDataOutputRecipe)
		if x := mod1.GridData.FilePath(); x != testData {
			t.Error(x)
		}

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueDaGridDataOutputRecipe)
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
