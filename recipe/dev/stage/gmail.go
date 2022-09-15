package stage

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

type Gmail struct {
	rc_recipe.RemarkSecret
	Peer   goog_conn.ConnGoogleMail
	Labels rp_model.RowReport
	UserId string
}

func (z *Gmail) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailReadonly,
	)
	z.Labels.SetModel(&mo_label.Label{},
		rp_model.HiddenColumns(
			"id",
		),
	)
	z.UserId = "me"
}

func (z *Gmail) Exec(c app_control.Control) error {
	labels, err := sv_label.New(z.Peer.Client(), z.UserId).List()
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

func (z *Gmail) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Gmail{}, rc_recipe.NoCustomValues)
}
