package sheet

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/sheets/service/sv_sheet"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Export struct {
	Peer           goog_conn.ConnGoogleSheets
	Data           da_griddata.GridDataOutput
	Range          string
	Id             string
	ValueRender    mo_string.SelectString
	DateTimeRender mo_string.SelectString
}

func (z *Export) Preset() {
	z.Peer.SetScopes(goog_auth.ScopeSheetsReadOnly)
	z.ValueRender.SetOptions(
		sv_sheet.ValueRenderOptionAliasFormatted,
		sv_sheet.ValueRenderOptionAliases...,
	)
	z.DateTimeRender.SetOptions(
		sv_sheet.DateTimeRenderOptionAliasSerialNumber,
		sv_sheet.DateTimeRenderOptionAliases...,
	)
}

func (z *Export) Exec(c app_control.Control) error {
	vr, err := sv_sheet.New(z.Peer.Context()).Export(
		z.Id,
		z.Range,
		sv_sheet.DateTimeRenderOption(z.DateTimeRender.Value()),
		sv_sheet.ValueRenderOption(z.ValueRender.Value()),
	)
	if err != nil {
		return err
	}
	for _, row := range vr.Values {
		z.Data.Row(row)
	}
	return nil
}

func (z *Export) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Clear{}, func(r rc_recipe.Recipe) {
		m := r.(*Clear)
		m.Id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		m.Range = "Sheet1"
	})
}
