package sheet

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/sheets/model/bo_sheet"
	"github.com/watermint/toolbox/domain/google/sheets/service/sv_spreadsheet"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer   goog_conn.ConnGoogleSheets
	Sheets rp_model.RowReport
	Id     string
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeSheetsReadOnly,
	)
	z.Sheets.SetModel(
		&bo_sheet.Sheet{},
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Sheets.Open(); err != nil {
		return err
	}

	spreadsheet, err := sv_spreadsheet.New(z.Peer.Client()).Resolve(z.Id)
	if err != nil {
		return err
	}
	sheets, err := spreadsheet.Sheets()
	if err != nil {
		return err
	}
	for _, sheet := range sheets {
		z.Sheets.Row(sheet)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.Id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	})
}
