package insight

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_insight"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/model/mo_int"
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

type Scan struct {
	Peer              dbx_conn.ConnScopedTeam
	Database          mo_path.FileSystemPath
	MaxRetries        mo_int.RangeInt
	ScanMemberFolders bool
	Errors            rp_model.RowReport
	Conclusion        app_msg.Message
	SkipSummarize     bool
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
	)
	if err != nil {
		return err
	}
	scanErr := ts.Scan()

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
