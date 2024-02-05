package insight

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_insight"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Scan struct {
	rc_recipe.RemarkSecret
	Peer              dbx_conn.ConnScopedTeam
	Database          mo_path.FileSystemPath
	MaxRetries        mo_int.RangeInt
	ScanMemberFolders bool
	Errors            rp_model.RowReport
}

func (z *Scan) Preset() {
	z.MaxRetries.SetRange(0, 10, 3)
	z.Errors.SetModel(&uc_insight.ApiErrorReport{})
}

func (z *Scan) Exec(c app_control.Control) error {
	if err := z.Errors.Open(); err != nil {
		return err
	}
	ts, err := uc_insight.NewTeamScanner(
		c,
		z.Peer.Client(),
		z.Database.Path(),
		uc_insight.MaxRetries(z.MaxRetries.Value()),
		uc_insight.ScanMemberFolders(z.ScanMemberFolders),
		uc_insight.OnErrorRecords(func(errCategory, errMessage, errTag, detail string) {
			z.Errors.Row(&uc_insight.ApiErrorReport{
				Category: errCategory,
				Message:  errMessage,
				Tag:      errTag,
				Detail:   detail,
			})
		}),
	)
	if err != nil {
		return err
	}
	return ts.Scan()
}

func (z *Scan) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("scan", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	return rc_exec.ExecMock(c, &Scan{}, func(r rc_recipe.Recipe) {
		m := r.(*Scan)
		m.Database = mo_path.NewFileSystemPath(filepath.Join(f, "scan.db"))
	})
}
