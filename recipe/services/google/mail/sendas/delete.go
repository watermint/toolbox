package sendas

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_sendas"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Delete struct {
	Peer   goog_conn.ConnGoogleMail
	UserId string
	Email  string
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailSettingsSharing,
	)
	z.UserId = "me"
}

func (z *Delete) Exec(c app_control.Control) error {
	return sv_sendas.New(z.Peer.Context(), z.UserId).Remove(z.Email)
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Email = "send_as@example.com"
	})
}
