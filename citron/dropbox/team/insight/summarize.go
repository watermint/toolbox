package insight

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_insight"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Summarize struct {
	rc_recipe.RemarkSecret
	Peer     dbx_conn.ConnScopedTeam
	Database mo_path.FileSystemPath
}

func (z *Summarize) Preset() {
}

func (z *Summarize) Exec(c app_control.Control) error {
	ts, err := uc_insight.NewTeamScanner(c, z.Peer.Client(), z.Database.Path())
	if err != nil {
		return err
	}
	return ts.Summarize()
}

func (z *Summarize) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("summarize", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	return rc_exec.ExecMock(c, &Scan{}, func(r rc_recipe.Recipe) {
		m := r.(*Scan)
		m.Database = mo_path.NewFileSystemPath(filepath.Join(f, "summarize.db"))
	})
}
