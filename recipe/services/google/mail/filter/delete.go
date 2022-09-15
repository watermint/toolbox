package filter

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_filter"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Delete struct {
	Peer   goog_conn.ConnGoogleMail
	UserId string
	Id     string
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailSettingsBasic,
	)
	z.UserId = "me"
}

func (z *Delete) Exec(c app_control.Control) error {
	return sv_filter.New(z.Peer.Client(), z.UserId).Delete(z.Id)
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Id = "Label_1"
	})
}
