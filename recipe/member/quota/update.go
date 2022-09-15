package quota

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member_quota"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member_quota"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"math"
)

type UpdateMemberQuota struct {
	Member *mo_member.Member
	Quota  int
}

type Update struct {
	Peer         dbx_conn.ConnScopedTeam
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
	Quota        mo_int.RangeInt
}

func (z *Update) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeMembersWrite,
	)
	z.File.SetModel(&mo_member_quota.MemberQuota{})
	z.OperationLog.SetModel(&mo_member_quota.MemberQuota{}, &mo_member_quota.MemberQuota{})
	z.Quota.SetRange(0, math.MaxInt32, 0)
}

func (z *Update) updateQuota(mq *UpdateMemberQuota) error {
	q := &mo_member_quota.Quota{
		TeamMemberId: mq.Member.TeamMemberId,
		Quota:        mq.Quota,
	}
	in := mo_member_quota.MemberQuota{
		Email: mq.Member.Email,
		Quota: mq.Quota,
	}

	newQuota, err := sv_member_quota.NewQuota(z.Peer.Client()).Update(q)
	if err != nil {
		z.OperationLog.Failure(err, in)
	} else {
		z.OperationLog.Success(in, mo_member_quota.NewMemberQuota(mq.Member, newQuota))
	}
	return nil
}

func (z *Update) Exec(c app_control.Control) error {
	ctx := z.Peer.Client()

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	err = z.OperationLog.Open()
	if err != nil {
		return err
	}

	var lastErr error

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("update", z.updateQuota)
		q := s.Get("update")
		lastErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			mq := m.(*mo_member_quota.MemberQuota)
			member, ok := emailToMember[mq.Email]
			if !ok {
				z.OperationLog.Failure(errors.New("member not found for an email"), mq)
				return nil
			}
			quota := z.Quota.Value()
			if mq.Quota != 0 {
				quota = mq.Quota
			}

			q.Enqueue(&UpdateMemberQuota{
				Member: member,
				Quota:  quota,
			})
			return nil
		})
	})

	return lastErr
}

func (z *Update) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Update{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("update-quota", "john@example.com,10")
		if err != nil {
			return
		}
		m := r.(*Update)
		m.Quota.SetValue(150)
		m.File.SetFilePath(f)
	})
	if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}
