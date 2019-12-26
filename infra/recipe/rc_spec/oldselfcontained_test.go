package rc_spec

import (
	"errors"
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_time"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type SelfContainedTestRow struct {
	Email string `json:"email"`
	Quota int    `json:"quota"`
}

type SelfContainedTestRecipe struct {
	ProgressStart app_msg.Message
	Start         mo_time.Time
	DbxPath       mo_path.DropboxPath
	CustomQuota   fd_file.RowFeed
	Enabled       bool
	Limit         int
	Name          string
	OperLog       rp_model.TransactionReport
	DataReport    rp_model.RowReport
}

func (z *SelfContainedTestRecipe) Exec(k rc_kitchen.Kitchen) error {
	ui := k.UI()
	ui.InfoM(z.ProgressStart)

	if !z.Enabled {
		return errors.New("!enabled")
	}
	if z.Limit != 20 {
		return errors.New("limit != 20")
	}
	if z.DbxPath.Path() != "/dropbox" {
		return errors.New("!= /dropbox")
	}
	if z.Name != "hey" {
		return errors.New("!= hey")
	}
	if z.Start.Iso8601() != "2010-11-12T13:14:15Z" {
		return errors.New("!= 2010-11-12T13:14:15Z")
	}
	err := z.OperLog.Open()
	if err != nil {
		return err
	}
	err = z.DataReport.Open()
	if err != nil {
		return err
	}
	if err := z.CustomQuota.EachRow(func(m interface{}, rowIndex int) error {
		row := m.(*SelfContainedTestRow)
		if row.Email != "orange@example.com" {
			return errors.New("!= orange@example.com")
		}
		z.OperLog.Success(row, nil)
		z.DataReport.Row(row)
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (z *SelfContainedTestRecipe) Test(c app_control.Control) error {
	return qt_recipe.NoTestRequired()
}

func (z *SelfContainedTestRecipe) Preset() {
	z.Limit = 10
	z.CustomQuota.SetModel(&SelfContainedTestRow{})
	z.OperLog.SetModel(&SelfContainedTestRow{}, &mo_file.ConcreteEntry{})
	z.DataReport.SetModel(&mo_file.ConcreteEntry{})
}

func TestSpecSelfContained_ApplyValues(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		scr := &SelfContainedTestRecipe{}
		spec := newOldSelfContained(scr)

		feedDir, err := ioutil.TempDir("", "feed")
		if err != nil {
			t.Error(err)
			return
		}
		feedPath := filepath.Join(feedDir, "fd_feed.csv")
		err = ioutil.WriteFile(feedPath, []byte("orange@example.com,10"), 0600)
		if err != nil {
			t.Error(err)
			return
		}

		f := flag.NewFlagSet("test", flag.ContinueOnError)
		spec.SetFlags(f, ctl.UI())
		err = f.Parse([]string{"-enabled",
			"-limit", "20",
			"-name", "hey",
			"-start", "2010-11-12T13:14:15Z",
			"-dbx-path", "/dropbox",
			"-custom-quota", feedPath,
		})
		if err != nil {
			t.Error(err)
			return
		}

		{
			rcp, k, err := spec.ApplyValues(ctl, rc_recipe.NoCustomValues)
			if err != nil {
				t.Error(err)
				return
			}
			if err = rcp.Exec(k); err != nil {
				t.Error(err)
			}
			if err = rcp.Test(ctl); err != nil {
				switch e := err.(type) {
				case *qt_recipe.ErrorNoTestRequired:
					ctl.Log().Debug("ok")
				default:
					t.Error(e)
				}
			}
		}
	})
}
