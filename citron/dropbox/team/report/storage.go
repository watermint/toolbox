package report

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Storage struct {
	rc_recipe.RemarkSecret
	Peer      dbx_conn.ConnScopedTeam
	Report    rp_model.RowReport
	StartDate mo_time.TimeOptional
	EndDate   mo_time.TimeOptional
}

func (z *Storage) Preset() {
	z.Report.SetModel(&mo_team.Storage{})
	z.Peer.SetScopes(
		dbx_auth.ScopeTeamInfoRead,
	)
}

func (z *Storage) Exec(c app_control.Control) error {
	if err := z.Report.Open(); err != nil {
		return err
	}

	storage, err := sv_team.NewReport(z.Peer.Client()).Storage(sv_team.NewSpan(z.StartDate, z.EndDate))
	if err != nil {
		return err
	}
	z.Report.Row(&storage)
	return nil
}

func (z *Storage) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Storage{}, rc_recipe.NoCustomValues)
}
