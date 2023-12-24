package sendas

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_sendas"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_sendas"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer   goog_conn.ConnGoogleMail
	SendAs rp_model.RowReport
	UserId string
}

func (z *List) Preset() {
	z.SendAs.SetModel(
		&mo_sendas.SendAs{},
		rp_model.HiddenColumns(
			"signature",
		),
	)
	z.Peer.SetScopes(
		goog_auth.ScopeGmailReadonly,
	)
	z.UserId = "me"
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.SendAs.Open(); err != nil {
		return err
	}

	entries, err := sv_sendas.New(z.Peer.Client(), z.UserId).List()
	if err != nil {
		return err
	}
	for _, entry := range entries {
		z.SendAs.Row(entry)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
