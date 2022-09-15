package batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Suspend struct {
	Peer                 dbx_conn.ConnScopedTeam
	File                 fd_file.RowFeed
	KeepData             bool
	OperationLog         rp_model.TransactionReport
	SkipAlreadySuspended app_msg.Message
	ErrorCantSuspend     app_msg.Message
}

func (z *Suspend) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeMembersWrite,
	)
	z.File.SetModel(&User{})
	z.OperationLog.SetModel(
		&User{},
		nil,
	)
}

func (z *Suspend) suspend(user *User, c app_control.Control) error {
	l := c.Log().With(esl.String("email", user.Email))
	svm := sv_member.New(z.Peer.Client())
	member, err := svm.ResolveByEmail(user.Email)
	if err != nil {
		l.Debug("Unable to resolve the user", esl.Error(err))
		z.OperationLog.Failure(err, user)
		return err
	}
	if member.Status == "suspended" {
		l.Debug("Skip: the member already suspended")
		z.OperationLog.Skip(z.SkipAlreadySuspended, user)
		return nil
	}
	err = svm.Suspend(member, sv_member.SuspendWipeData(!z.KeepData))
	if err != nil {
		c.UI().Error(z.ErrorCantSuspend.With("Error", err))
		return err
	}
	l.Debug("Suspended")
	z.OperationLog.Success(user, nil)
	return nil
}

func (z *Suspend) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}
	var lastErr, fileErr error
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("suspend", z.suspend, c)
		q := s.Get("suspend")
		fileErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			u := m.(*User)
			q.Enqueue(u)
			return nil
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	return lang.NewMultiErrorOrNull(lastErr, fileErr)
}

func (z *Suspend) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("suspend", "kevin@example.com\ndavid@example.com\n")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecMock(c, &Suspend{}, func(r rc_recipe.Recipe) {
		m := r.(*Suspend)
		m.File.SetFilePath(f)
	})
}
