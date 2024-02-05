package insight

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_insight"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Scanretry struct {
	rc_recipe.RemarkSecret
	Peer       dbx_conn.ConnScopedTeam
	Database   mo_path.FileSystemPath
	Errors     rp_model.RowReport
	Conclusion app_msg.Message
}

func (z *Scanretry) Preset() {
	z.Errors.SetModel(&uc_insight.ApiErrorReport{})
}

func (z *Scanretry) Exec(c app_control.Control) error {
	if err := z.Errors.Open(); err != nil {
		return err
	}
	ts, err := uc_insight.NewTeamScanner(
		c,
		z.Peer.Client(),
		z.Database.Path(),
	)
	if err != nil {
		return err
	}
	defer func() {
		numErr, err := ts.ReportLastErrors(func(errCategory string, errMessage string, errTag string, detail string) {
			z.Errors.Row(&uc_insight.ApiErrorReport{
				Category: errCategory,
				Message:  errMessage,
				Tag:      errTag,
				Detail:   detail,
			})
		})
		if err == nil {
			c.UI().Info(z.Conclusion.With("Count", numErr))
		}
	}()

	return ts.RetryErrors()
}

func (z *Scanretry) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("scan", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()
	// create a database
	err = rc_exec.ExecMock(c, &Scan{}, func(r rc_recipe.Recipe) {
		m := r.(*Scan)
		m.Database = mo_path.NewFileSystemPath(filepath.Join(f, "scan.db"))
	})
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &Scanretry{}, func(r rc_recipe.Recipe) {
		m := r.(*Scan)
		m.Database = mo_path.NewFileSystemPath(filepath.Join(f, "scan.db"))
	})
}
