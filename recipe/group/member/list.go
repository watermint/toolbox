package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_group_member"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type ListWorker struct {
	// job context
	group *mo_group.Group

	// recipe's context
	ctl  app_control.Control
	conn api_context.Context
	rep  rp_model.RowReport
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
	Peer        rc_conn.ConnBusinessInfo
	GroupMember rp_model.RowReport
}

func (z *List) Preset() {
	z.GroupMember.SetModel(&mo_group_member.GroupMember{})
}

func (z *List) Exec(k rc_kitchen.Kitchen) error {
	gsv := sv_group.New(z.Peer.Context())
	groups, err := gsv.List()
	if err != nil {
		return err
	}

	if err := z.GroupMember.Open(); err != nil {
		return err
	}

	q := k.NewQueue()
	for _, group := range groups {
		w := &ListWorker{
			group: group,
			ctl:   k.Control(),
			conn:  z.Peer.Context(),
			rep:   z.GroupMember,
		}
		q.Enqueue(w)
	}
	q.Wait()

	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "group_member", func(cols map[string]string) error {
		if _, ok := cols["group_id"]; !ok {
			return errors.New("group_id is not found")
		}
		if _, ok := cols["team_member_id"]; !ok {
			return errors.New("team_member_id is not found")
		}
		return nil
	})
}
