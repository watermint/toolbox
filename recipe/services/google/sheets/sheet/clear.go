package sheet

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/sheets/service/sv_sheet"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Clear struct {
	Peer    goog_conn.ConnGoogleSheets
	Range   string
	Id      string
	Success app_msg.Message
}

func (z *Clear) Preset() {
	z.Peer.SetScopes(goog_auth.ScopeSheetsReadWrite)
}

func (z *Clear) Exec(c app_control.Control) error {
	clearedRange, err := sv_sheet.New(z.Peer.Client()).Clear(z.Id, z.Range)
	if err != nil {
		return err
	}
	c.UI().Success(z.Success.With("Range", clearedRange))
	return nil
}

func (z *Clear) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Clear{}, func(r rc_recipe.Recipe) {
		m := r.(*Clear)
		m.Id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		m.Range = "Sheet1"
	})
}
