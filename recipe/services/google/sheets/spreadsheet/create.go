package spreadsheet

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/sheets/model/bo_spreadsheet"
	"github.com/watermint/toolbox/domain/google/sheets/service/sv_spreadsheet"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Create struct {
	Peer       goog_conn.ConnGoogleSheets
	Sreadsheet rp_model.RowReport
	Title      string
}

func (z *Create) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeSheetsReadWrite,
	)
	z.Sreadsheet.SetModel(
		&bo_spreadsheet.Spreadsheet{},
	)
}

func (z *Create) Exec(c app_control.Control) error {
	if err := z.Sreadsheet.Open(); err != nil {
		return err
	}

	s, err := sv_spreadsheet.New(z.Peer.Client()).Create(z.Title)
	if err != nil {
		return err
	}
	z.Sreadsheet.Row(s)
	return nil
}

func (z *Create) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Create{}, func(r rc_recipe.Recipe) {
		m := r.(*Create)
		m.Title = "test"
	})
}
