package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_insight"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Scan struct {
	rc_recipe.RemarkSecret
	Peer     dbx_conn.ConnScopedTeam
	Database mo_path.FileSystemPath
}

func (z *Scan) Preset() {
}

func (z *Scan) Exec(c app_control.Control) error {
	ts, err := uc_insight.NewTeamScanner(c, z.Peer.Client(), z.Database.Path())
	if err != nil {
		return err
	}
	return ts.ScanTeam()
}

func (z *Scan) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
