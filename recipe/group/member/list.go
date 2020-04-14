package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type MsgList struct {
	ProgressScan app_msg.Message
}

var (
	MList = app_msg.Apply(&MsgList{}).(*MsgList)
)

type ListWorker struct {
	// job context
	group *mo_group.Group

	// recipe's context
	ctl  app_control.Control
	conn dbx_context.Context
	rep  rp_model.RowReport
}

func (z *ListWorker) Exec() error {
	l := z.ctl.Log()
	ui := z.ctl.UI()

	ui.Progress(MList.ProgressScan.With("Group", z.group.GroupName))
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

type List struct {
	Peer        dbx_conn.ConnBusinessInfo
	GroupMember rp_model.RowReport
}

func (z *List) Preset() {
	z.GroupMember.SetModel(
		&mo_group_member.GroupMember{},
		rp_model.HiddenColumns(
			"group_id",
			"account_id",
			"team_member_id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	gsv := sv_group.New(z.Peer.Context())
	groups, err := gsv.List()
	if err != nil {
		return err
	}

	if err := z.GroupMember.Open(); err != nil {
		return err
	}

	q := c.NewQueue()
	for _, group := range groups {
		w := &ListWorker{
			group: group,
			ctl:   c,
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
		if _, ok := cols["group_name"]; !ok {
			return errors.New("group_name is not found")
		}
		if _, ok := cols["email"]; !ok {
			return errors.New("email is not found")
		}
		return nil
	})
}
