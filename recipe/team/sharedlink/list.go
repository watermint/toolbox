package sharedlink

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type ListVO struct {
	Peer app_conn.ConnBusinessFile
}

type ListWorker struct {
	member *mo_member.Member
	conn   api_context.Context
	rep    app_report.Report
	ctl    app_control.Control
}

func (z *ListWorker) Exec() error {
	z.ctl.UI().Info("recipe.team.sharedlink.list.scan", app_msg.P{"MemberEmail": z.member.Email})
	mc := z.conn.AsMemberId(z.member.TeamMemberId)
	links, err := sv_sharedlink.New(mc).List()
	if err != nil {
		return err
	}
	for _, link := range links {
		lm := mo_sharedlink.NewSharedLinkMember(link, z.member)
		z.rep.Row(lm)
	}
	return nil
}

type List struct {
}

func (*List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (*List) Exec(k app_kitchen.Kitchen) error {
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
	rep, err := k.Report("sharedlink", &mo_sharedlink.SharedLinkMember{})
	if err != nil {
		return err
	}
	defer rep.Close()

	q := k.NewQueue()
	for _, member := range members {
		q.Enqueue(&ListWorker{
			member: member,
			conn:   conn,
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
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "sharedlink", func(cols map[string]string) error {
		if _, ok := cols["shared_link_id"]; !ok {
			return errors.New("`shared_link_id` is not found")
		}
		if _, ok := cols["team_member_id"]; !ok {
			return errors.New("`team_member_id` is not found")
		}
		return nil
	})
}
