package bulk

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
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

type AddRecord struct {
	GroupName   string
	MemberEmail string
}

type Add struct {
	rc_recipe.RemarkIrreversible
	Peer                         dbx_conn.ConnScopedTeam
	File                         fd_file.RowFeed
	OperationLog                 rp_model.TransactionReport
	SkipTheUserAlreadyInTheGroup app_msg.Message
}

func (z *Add) Preset() {
	z.File.SetModel(&AddRecord{})
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeGroupsWrite,
	)
	z.OperationLog.SetModel(&AddRecord{}, nil)
}

func (z *Add) add(r *AddRecord, svg sv_group.Group, c app_control.Control) error {
	l := c.Log().With(esl.Any("record", r))
	group, err := svg.ResolveByName(r.GroupName)
	if err != nil {
		l.Debug("Unable to resolve the group", esl.Error(err))
		return err
	}

	updated, err := sv_group_member.New(z.Peer.Context(), group).Add(sv_group_member.ByEmail(r.MemberEmail))
	if err != nil {
		if dbx_error.NewErrors(err).IsDuplicateUser() {
			z.OperationLog.Skip(z.SkipTheUserAlreadyInTheGroup, r)
			l.Debug("The user is already in the member", esl.Error(err))
			return nil
		}

		l.Debug("Unable to update group member", esl.Error(err))
		z.OperationLog.Failure(err, r)
		return err
	}
	l.Debug("The member is successfully updated", esl.Any("updated", updated))
	z.OperationLog.Success(r, nil)
	return nil
}

func (z *Add) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	svg := sv_group.NewCached(z.Peer.Context())

	queueIdAdd := "add"

	var lastErr error

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define(queueIdAdd, z.add, svg, c)
		q := s.Get(queueIdAdd)

		lastErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	})

	return lastErr
}

func (z *Add) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("add", "Sales,taro@example.com\nSales,hanako@example.com\n")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.File.SetFilePath(f)
	})
}
