package sharedlink

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type List struct {
	Peer       dbx_conn.ConnScopedIndividual
	SharedLink rp_model.RowReport
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeSharingRead,
	)
	z.SharedLink.SetModel(
		&mo_sharedlink.Metadata{},
		rp_model.HiddenColumns(
			"id",
			"account_id",
			"team_member_id",
		),
	)
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "shared_link", func(cols map[string]string) error {
		if _, ok := cols["name"]; !ok {
			return errors.New("`name` is not found")
		}
		return nil
	})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.SharedLink.Open(); err != nil {
		return err
	}

	links, err := sv_sharedlink.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	for _, link := range links {
		z.SharedLink.Row(link)
	}

	return nil
}
