package quota

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_member_quota"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_member_quota"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
)

type ListVO struct {
	Peer app_conn.ConnBusinessMgmt
}

type ListWorker struct {
	member *mo_member.Member
	ctx    api_context.Context
	rep    app_report.Report
	ctl    app_control.Control
}

func (z *ListWorker) Exec() error {
	l := z.ctl.Log()

	z.ctl.UI().Info("recipe.member.quota.list.scan", app_msg.P{"MemberEmail": z.member.Email})
	l.Debug("Scan member", zap.String("Routine", ut_runtime.GetGoRoutineName()), zap.Any("Member", z.member))

	q, err := sv_member_quota.NewQuota(z.ctx).Resolve(z.member.TeamMemberId)
	if err != nil {
		return err
	}
	z.rep.Row(mo_member_quota.NewMemberQuota(z.member, q))
	return nil
}

type List struct {
}

func (z *List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*ListVO)
	conn, err := lvo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(conn).List()
	if err != nil {
		return err
	}

	// Write report
	rep, err := k.Report("member_quota", &mo_member_quota.MemberQuota{})
	if err != nil {
		return err
	}
	defer rep.Close()

	q := k.NewQueue()
	for _, member := range members {
		q.Enqueue(&ListWorker{
			member: member,
			ctx:    conn,
			rep:    rep,
			ctl:    k.Control(),
		})
	}
	q.Wait()
	return nil
}

func (z *List) Test(c app_control.Control) error {
	lvo := &ListVO{}
	if !app_test.ApplyTestPeers(c, lvo) {
		return qt_test.NotEnoughResource()
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "member_quota", func(cols map[string]string) error {
		if _, ok := cols["email"]; !ok {
			return errors.New("`email` is not found")
		}
		if _, ok := cols["quota"]; !ok {
			return errors.New("`quota` is not found")
		}
		return nil
	})
}
