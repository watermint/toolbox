package thread

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_thread"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_thread"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer    goog_conn.ConnGoogleMail
	Threads rp_model.RowReport
	UserId  string
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailReadonly,
	)
	z.Threads.SetModel(&mo_thread.Thread{},
		rp_model.HiddenColumns(
			"id",
		),
	)
	z.UserId = "me"
}

func (z *List) Exec(c app_control.Control) error {
	threads, err := sv_thread.New(z.Peer.Context(), z.UserId).List()
	if err != nil {
		return err
	}
	if err := z.Threads.Open(); err != nil {
		return err
	}
	for _, thread := range threads {
		z.Threads.Row(thread)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &List{}, "recipe-services-google-mail-thread-list.json.gz", rc_recipe.NoCustomValues)
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
