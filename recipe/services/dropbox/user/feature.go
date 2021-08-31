package user

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_user"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_user"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Feature struct {
	Peer   dbx_conn.ConnScopedIndividual
	Report rp_model.RowReport
}

func (z *Feature) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeAccountInfoRead,
	)
	z.Report.SetModel(&mo_user.Feature{})
}

func (z *Feature) Exec(c app_control.Control) error {
	if err := z.Report.Open(); err != nil {
		return err
	}

	features, err := sv_user.New(z.Peer.Context()).Features()
	if err != nil {
		return err
	}

	z.Report.Row(features)
	return nil
}

func (z *Feature) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Feature{}, rc_recipe.NoCustomValues)
}
