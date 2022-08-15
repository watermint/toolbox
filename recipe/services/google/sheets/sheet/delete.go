package sheet

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/sheets/service/sv_sheet"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Delete struct {
	Peer    goog_conn.ConnGoogleSheets
	Id      string
	SheetId string
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(goog_auth.ScopeSheetsReadWrite)
}

func (z *Delete) Exec(c app_control.Control) error {
	return sv_sheet.New(z.Peer.Context()).Delete(z.Id, z.SheetId)
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		m.SheetId = "1234567890"
	})
}
