package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Suspend struct {
	Peer                 dbx_conn.ConnScopedTeam
	Email                string
	KeepData             bool
	SkipAlreadySuspended app_msg.Message
	SuccessSuspended     app_msg.Message
	ErrorMemberNotFound  app_msg.Message
	ErrorCantSuspend     app_msg.Message
}

func (z *Suspend) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeMembersWrite,
	)
}

func (z *Suspend) Exec(c app_control.Control) error {
	l := c.Log().With(esl.String("email", z.Email), esl.Bool("keepData", z.KeepData))
	svm := sv_member.New(z.Peer.Context())
	member, err := svm.ResolveByEmail(z.Email)
	if err != nil {
		l.Debug("Unable to resolve the member", esl.Error(err))
		c.UI().Error(z.ErrorMemberNotFound.With("Error", err))
		return err
	}
	if member.Status == "suspended" {
		l.Debug("Skip: the member already suspended")
		c.UI().Info(z.SkipAlreadySuspended)
		return nil
	}
	err = svm.Suspend(member, sv_member.SuspendWipeData(!z.KeepData))
	if err != nil {
		c.UI().Error(z.ErrorCantSuspend.With("Error", err))
		return err
	}
	c.UI().Info(z.SuccessSuspended.With("Email", z.Email))
	return nil
}

func (z *Suspend) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Suspend{}, func(r rc_recipe.Recipe) {
		m := r.(*Suspend)
		m.Email = "john@example.com"
	})
}
