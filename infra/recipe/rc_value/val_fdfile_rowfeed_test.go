package rc_value

import (
	"encoding/json"
	"flag"
	"os"
	"testing"

	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type ValueFdFileRowTestModel struct {
	Email string `json:"email"`
}

type ValueFdFileRowFeedRecipe struct {
	File fd_file.RowFeed
}

func (z ValueFdFileRowFeedRecipe) Preset() {
	z.File.SetModel(&ValueFdFileRowTestModel{})
}

func (z ValueFdFileRowFeedRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z ValueFdFileRowFeedRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueFdFileRowFeedSuccess(t *testing.T) {
	qt_file.TestWithTestFile(t, "email.csv", "john@example.com", func(path string) {
		err := qt_control.WithControl(func(c app_control.Control) error {
			rcp0 := &ValueFdFileRowFeedRecipe{}
			repo := NewRepository(rcp0)

			// Parse flags
			flg := flag.NewFlagSet("value", flag.ContinueOnError)
			repo.ApplyFlags(flg, c.UI())
			if err := flg.Parse([]string{"-file", path}); err != nil {
				t.Error(err)
				return err
			}

			// Apply parsed values
			rcp1 := repo.Apply()
			mod1 := rcp1.(*ValueFdFileRowFeedRecipe)
			if mod1.File.FilePath() != path {
				t.Error(mod1)
			}

			// Spin up
			rcp2, err := repo.SpinUp(c)
			if err != nil {
				t.Error(err)
				return err
			}
			mod2 := rcp2.(*ValueFdFileRowFeedRecipe)
			if mod2.File.FilePath() != path {
				t.Error(mod2)
			}
			err = mod2.File.EachRow(func(m interface{}, rowIndex int) error {
				dm := m.(*ValueFdFileRowTestModel)
				if dm.Email != "john@example.com" {
					t.Error(dm)
				}
				return nil
			})
			if err != nil {
				t.Error(err)
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
	})
}

func TestValueFdFileRowFeedNotFound(t *testing.T) {
	path := "/tmp/no_existent/data.csv"
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueFdFileRowFeedRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-file", path}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueFdFileRowFeedRecipe)
		if mod1.File.FilePath() != path {
			t.Error(mod1)
		}

		// Spin up; should fail
		_, err := repo.SpinUp(c)
		if err == nil {
			t.Error(err)
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestValueFdFileRowFeedEmpty(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueFdFileRowFeedRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueFdFileRowFeedRecipe)
		if mod1.File.FilePath() != "" {
			t.Error(mod1)
		}

		// Spin up; should fail
		_, err := repo.SpinUp(c)
		if err != ErrorMissingRequiredOption {
			t.Error(err)
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestValueFdFileRowFeed_Capture(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		// Prepare test data
		testDataFile, err := os.CreateTemp("", "row_feed*.csv")
		if err != nil {
			t.Error(err)
			return err
		}
		testDataPath := testDataFile.Name()
		defer func() {
			_ = os.Remove(testDataPath)
		}()
		testData := `test@example.com,Test,Example`
		_, _ = testDataFile.WriteString(testData)
		_ = testDataFile.Close()

		var capJson es_json.Json

		{
			v := newValueFdFileRowFeed("123")
			vrf := v.Init().(fd_file.RowFeed)
			vrf.SetFilePath(testDataPath)

			vc, err := v.Capture(c)
			if err != nil {
				t.Error(err)
				return err
			}

			capData, err := json.Marshal(vc)
			if err != nil {
				t.Error(err)
				return err
			}

			capJson, err = es_json.Parse(capData)
			if err != nil {
				t.Error(err)
				return err
			}
		}

		{
			v2 := newValueFdFileRowFeed("123")
			if err := v2.Restore(capJson, c); err != nil {
				t.Error(err)
			}

			vrf2 := v2.Init().(fd_file.RowFeed)
			if vrf2.FilePath() == "" {
				t.Error("empty file path")
			}

			restoreData, err := os.ReadFile(vrf2.FilePath())
			if err != nil {
				t.Error(err)
				return err
			}

			if string(restoreData) != testData {
				t.Error(restoreData)
			}
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
