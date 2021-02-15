package sheet

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Export struct {
	Peer  goog_conn.ConnGoogleSheets
	Range string
	Id    string
}

func (z *Export) Preset() {
	z.Peer.SetScopes(goog_auth.ScopeSheetsReadWrite)
}

func (z *Export) Exec(c app_control.Control) error {
	return nil
}

func (z *Export) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Clear{}, func(r rc_recipe.Recipe) {
		m := r.(*Clear)
		m.Id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		m.Range = "Sheet1"
	})
}
