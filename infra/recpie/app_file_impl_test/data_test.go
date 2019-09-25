package app_file_impl_test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_file_impl"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"path/filepath"
	"testing"
)

func TestCsvData_EachRow(t *testing.T) {
	type DataRow struct {
		Name     string
		Email    string
		Count    int
		Verified bool
	}
	type DataRow2 struct {
		Name2     string
		Email2    string
		Count2    int
		Verified2 bool
	}
	type DataRowMissingField struct {
		Email    string
		Count    int
		Verified bool
	}
	type DataRowMoreField struct {
		Name     string
		Email    string
		Count    int
		Verified bool
		Quota    int
	}

	app_test.TestWithControl(t, func(ctl app_control.Control) {
		dataRows := make([]*DataRow, 0)
		dataRows = append(dataRows, &DataRow{
			Name:     "John",
			Email:    "john@example.com",
			Count:    30,
			Verified: false,
		})
		dataRows = append(dataRows, &DataRow{
			Name:     "Thomas",
			Email:    "thomas@example.com",
			Count:    2,
			Verified: true,
		})
		dataRows = append(dataRows, &DataRow{
			Name:     "Kevin",
			Email:    "kevin@example.com",
			Count:    20,
			Verified: true,
		})

		rep, err := app_report.New("data_row", &DataRow{}, ctl)
		if err != nil {
			t.Error(err)
			return
		}
		for _, d := range dataRows {
			rep.Row(d)
		}
		rep.Close()

		repPath := filepath.Join(ctl.Workspace().Job(), "reports")
		repFile := filepath.Join(repPath, "data_row.csv")

		// fieldName
		{
			cd := app_file_impl.CsvData{
				FilePath: repFile,
			}
			if err := cd.Model(ctl, &DataRow{}); err != nil {
				t.Error(err)
				return
			}

			err = cd.EachRow(func(m interface{}, rowIndex int) error {
				d := m.(*DataRow)
				orig := dataRows[rowIndex-1]

				if d.Email != orig.Email {
					t.Error("invalid email")
				}
				if d.Name != orig.Name {
					t.Error("invalid name")
				}
				if d.Count != orig.Count {
					t.Error("invalid count")
				}
				if d.Verified != orig.Verified {
					t.Error("invalid verified")
				}
				return nil
			})
			if err != nil {
				t.Error(err)
				return
			}
		}

		// order
		{
			cd := app_file_impl.CsvData{
				FilePath: repFile,
			}
			if err := cd.Model(ctl, &DataRow2{}); err != nil {
				t.Error(err)
				return
			}

			err = cd.EachRow(func(m interface{}, rowIndex int) error {
				d := m.(*DataRow2)
				orig := dataRows[rowIndex-1]

				if d.Email2 != orig.Email {
					t.Error("invalid email")
				}
				if d.Name2 != orig.Name {
					t.Error("invalid name")
				}
				if d.Count2 != orig.Count {
					t.Error("invalid count")
				}
				if d.Verified2 != orig.Verified {
					t.Error("invalid verified")
				}
				return nil
			})
			if err != nil {
				t.Error(err)
				return
			}
		}

		// missing field
		{
			cd := app_file_impl.CsvData{
				FilePath: repFile,
			}
			if err := cd.Model(ctl, &DataRowMissingField{}); err != nil {
				t.Error(err)
				return
			}

			err = cd.EachRow(func(m interface{}, rowIndex int) error {
				d := m.(*DataRowMissingField)
				orig := dataRows[rowIndex-1]

				if d.Email != orig.Email {
					t.Error("invalid email")
				}
				if d.Count != orig.Count {
					t.Error("invalid count")
				}
				if d.Verified != orig.Verified {
					t.Error("invalid verified")
				}
				return nil
			})
			if err == nil {
				t.Error("should fail")
				return
			}
		}

		// more field
		{
			cd := app_file_impl.CsvData{
				FilePath: repFile,
			}
			if err := cd.Model(ctl, &DataRowMoreField{}); err != nil {
				t.Error(err)
				return
			}

			err = cd.EachRow(func(m interface{}, rowIndex int) error {
				d := m.(*DataRowMoreField)
				orig := dataRows[rowIndex-1]

				if d.Email != orig.Email {
					t.Error("invalid email")
				}
				if d.Name != orig.Name {
					t.Error("invalid name")
				}
				if d.Count != orig.Count {
					t.Error("invalid count")
				}
				if d.Verified != orig.Verified {
					t.Error("invalid verified")
				}
				if d.Quota != 0 {
					t.Errorf("invalid quota")
				}
				return nil
			})
			if err != nil {
				t.Error(err)
				return
			}
		}
	})
}
