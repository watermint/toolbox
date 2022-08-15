package sheet

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/sheets/model/bo_sheet"
	"github.com/watermint/toolbox/domain/google/sheets/service/sv_sheet"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Create struct {
	Peer    goog_conn.ConnGoogleSheets
	Id      string
	Title   string
	Cols    int
	Rows    int
	Created rp_model.RowReport
}

func (z *Create) Preset() {
	z.Peer.SetScopes(goog_auth.ScopeSheetsReadWrite)
	z.Created.SetModel(
		&bo_sheet.Sheet{},
	)
	z.Cols = 26
	z.Rows = 1000
}

func (z *Create) Exec(c app_control.Control) error {
	if err := z.Created.Open(); err != nil {
		return err
	}

	sheet, err := sv_sheet.New(z.Peer.Context()).Create(z.Id, z.Title, z.Cols, z.Rows)
	if err != nil {
		return err
	}
	z.Created.Row(sheet)
	return nil
}

func (z *Create) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Create{}, func(r rc_recipe.Recipe) {
		m := r.(*Create)
		m.Id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		m.Title = "New Sheet"
	})
}
