package update

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Invisible struct {
	rc_recipe.RemarkIrreversible
	Peer         dbx_conn.ConnScopedTeam
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
}

func (z *Invisible) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeMembersWrite,
	)
	z.File.SetModel(&VisibilityRecord{})
	z.OperationLog.SetModel(&VisibilityRecord{}, &mo_member.Member{})
}

func (z *Invisible) invisible(r *VisibilityRecord, svm sv_member.Member, c app_control.Control) error {
	l := c.Log().With(esl.Any("record", r))

	updated, err := svm.UpdateVisibility(r.Email, false)
	if err != nil {
		l.Debug("Unable to update member visibility", esl.Error(err))
		z.OperationLog.Failure(err, r)
		return err
	}
	z.OperationLog.Success(r, updated)
	return nil
}

func (z *Invisible) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	svm := sv_member.NewCached(z.Peer.Client())

	var lastErr, feedErr error
	queueId := "invisible"
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define(queueId, z.invisible, svm, c)
		q := s.Get(queueId)

		feedErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	if lastErr != nil {
		return lastErr
	}
	if feedErr != nil {
		return feedErr
	}
	return nil
}

func (z *Invisible) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("visibility", "taro@example.com\nhanako@example.com\n")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.ExecMock(c, &Invisible{}, func(r rc_recipe.Recipe) {
		m := r.(*Invisible)
		m.File.SetFilePath(f)
	})
}
