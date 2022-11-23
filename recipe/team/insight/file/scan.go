package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Scan struct {
	rc_recipe.RemarkSecret
	Peer dbx_conn.ConnScopedTeam
}

func (z *Scan) Preset() {
	//TODO implement me
	panic("implement me")
}

func (z *Scan) Exec(c app_control.Control) error {

	//TODO implement me
	panic("implement me")
}

func (z *Scan) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
