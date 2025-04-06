package insight

import (
	"os"
	"path/filepath"

	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_insight"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type Scanretry struct {
	rc_recipe.RemarkSecret
	Peer          dbx_conn.ConnScopedTeam
	Database      mo_path.FileSystemPath
	Errors        rp_model.RowReport
	Conclusion    app_msg.Message
	SkipSummarize bool
	BasePath      mo_string.SelectString
}

func (z *Scanretry) Preset() {
	z.Errors.SetModel(&uc_insight.ApiErrorReport{})
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Scanretry) Exec(c app_control.Control) error {
	if err := z.Errors.Open(); err != nil {
		return err
	}
	ts, err := uc_insight.NewTeamScanner(
		c,
		z.Peer.Client(),
		z.Database.Path(),
		uc_insight.BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value())),
	)
	if err != nil {
		return err
	}

	scanErr := ts.RetryErrors()

	numErr, reportErr := ts.ReportLastErrors(func(errCategory string, errMessage string, errTag string, detail string) {
		z.Errors.Row(&uc_insight.ApiErrorReport{
			Category: errCategory,
			Message:  errMessage,
			Tag:      errTag,
			Detail:   detail,
		})
	})
	if reportErr == nil {
		c.UI().Info(z.Conclusion.With("Count", numErr))
		if numErr < 1 && !z.SkipSummarize {
			summarizer, err := uc_insight.NewSummary(c, z.Database.Path())
			if err != nil {
				return err
			}
			return summarizer.Summarize()
		}
	}
	return es_lang.NewMultiErrorOrNull(scanErr, reportErr)
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
		m := r.(*Scanretry)
		m.Database = mo_path.NewFileSystemPath(filepath.Join(f, "scan.db"))
	})
}
