package label

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_label"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_label"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer   goog_conn.ConnGoogleMail
	Labels rp_model.RowReport
	UserId string
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailLabels,
	)
	z.Labels.SetModel(&mo_label.Label{},
		rp_model.HiddenColumns(
			"id",
		),
	)
	z.UserId = "me"
}

func (z *List) Exec(c app_control.Control) error {
	labels, err := sv_label.New(z.Peer.Context(), z.UserId).List()
	if err != nil {
		return err
	}
	if err := z.Labels.Open(); err != nil {
		return err
	}
	for _, label := range labels {
		z.Labels.Row(label)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
