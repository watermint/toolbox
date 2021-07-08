package auth

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type All struct {
	rc_recipe.RemarkSecret
	Peer dbx_conn.ConnScopedTeam
}

func (z *All) Preset() {
}

func (z *All) Exec(c app_control.Control) error {
	return nil
}

func (z *All) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &All{}, rc_recipe.NoCustomValues)
}
