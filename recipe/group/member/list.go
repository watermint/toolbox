package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_group_member"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
)

type ListVO struct {
	Peer app_conn.ConnBusinessInfo
}

type ListWorker struct {
	// job context
	group *mo_group.Group

	// recipe's context
	ctl  app_control.Control
	conn api_context.Context
	rep  rp_model.Report
}

func (z *ListWorker) Exec() error {
	l := z.ctl.Log()

	z.ctl.UI().Info("recipe.group.member.list.progress.scan", app_msg.P{"Group": z.group.GroupName})
	l.Debug("Scan group", zap.String("Routine", ut_runtime.GetGoRoutineName()), zap.Any("Group", z.group))

	msv := sv_group_member.New(z.conn, z.group)
	members, err := msv.List()
	if err != nil {
		return err
	}
	for _, m := range members {
		row := mo_group_member.NewGroupMember(z.group, m)
		z.rep.Row(row)
	}
	return nil
}

const (
	reportList = "group_member"
)

type List struct {
}

func (z *List) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportList, &mo_group_member.GroupMember{}),
	}
}

func (*List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*ListVO)
	connInfo, err := lvo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	gsv := sv_group.New(connInfo)
	groups, err := gsv.List()
	if err != nil {
		return err
	}

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportList)
	if err != nil {
		return err
	}
	defer rep.Close()

	q := k.NewQueue()
	for _, group := range groups {
		w := &ListWorker{
			group: group,
			ctl:   k.Control(),
			conn:  connInfo,
			rep:   rep,
		}
		q.Enqueue(w)
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
	return app_test.TestRows(c, "group_member", func(cols map[string]string) error {
		if _, ok := cols["group_id"]; !ok {
			return errors.New("group_id is not found")
		}
		if _, ok := cols["team_member_id"]; !ok {
			return errors.New("team_member_id is not found")
		}
		return nil
	})
}
