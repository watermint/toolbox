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

type Unsuspend struct {
	Peer                     dbx_conn.ConnScopedTeam
	Email                    string
	SkipMemberIsNotSuspended app_msg.Message
	SuccessUnsuspended       app_msg.Message
	ErrorMemberNotFound      app_msg.Message
	ErrorCantUnsuspend       app_msg.Message
}

func (z *Unsuspend) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeMembersWrite,
	)
}

func (z *Unsuspend) Exec(c app_control.Control) error {
	l := c.Log().With(esl.String("email", z.Email))
	svm := sv_member.New(z.Peer.Client())
	member, err := svm.ResolveByEmail(z.Email)
	if err != nil {
		c.UI().Error(z.ErrorMemberNotFound.With("Error", err))
		return err
	}
	if member.Status != "suspended" {
		l.Debug("Skip: the member is not suspended")
		c.UI().Info(z.SkipMemberIsNotSuspended)
		return nil
	}
	err = svm.Unsuspend(member)
	if err != nil {
		c.UI().Error(z.ErrorCantUnsuspend.With("Error", err))
		return err
	}
	c.UI().Info(z.SuccessUnsuspended.With("Email", z.Email))
	return nil
}

func (z *Unsuspend) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Unsuspend{}, func(r rc_recipe.Recipe) {
		m := r.(*Unsuspend)
		m.Email = "john@example.com"
	})
}
