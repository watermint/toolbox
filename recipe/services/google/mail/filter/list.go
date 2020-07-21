package filter

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/model/mo_filter"
	"github.com/watermint/toolbox/domain/google/service/sv_filter"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer    goog_conn.ConnGoogleMail
	Filters rp_model.RowReport
	UserId  string
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailReadonly,
	)
	z.Filters.SetModel(&mo_filter.Filter{},
		rp_model.HiddenColumns(
			"id",
		),
	)
	z.UserId = "me"
}

func (z *List) Exec(c app_control.Control) error {
	filters, err := sv_filter.New(z.Peer.Context(), z.UserId).List()
	if err != nil {
		return err
	}
	if err := z.Filters.Open(); err != nil {
		return err
	}
	for _, filter := range filters {
		z.Filters.Row(filter)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
