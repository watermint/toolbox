package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
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
	conn dbx_client.Client
	rep  rp_model.RowReport
}

func (z *ListWorker) Exec() error {
	l := z.ctl.Log()
	ui := z.ctl.UI()

	ui.Progress(MList.ProgressScan.With("Group", z.group.GroupName))
	l.Debug("Scan group", esl.String("Routine", es_goroutine.GetGoRoutineName()), esl.Any("Group", z.group))

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
	Peer        dbx_conn.ConnScopedTeam
	GroupMember rp_model.RowReport
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsRead,
	)
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
	l := c.Log()

	gsv := sv_group.New(z.Peer.Client())
	groups, err := gsv.List()
	if err != nil {
		return err
	}

	if err := z.GroupMember.Open(); err != nil {
		return err
	}

	memberList := func(group *mo_group.Group) error {
		ll := l.With(esl.String("Routine", es_goroutine.GetGoRoutineName()), esl.Any("Group", group))
		ll.Debug("Scan group")

		msv := sv_group_member.New(z.Peer.Client(), group)
		members, err := msv.List()
		if err != nil {
			ll.Debug("Unable to list members", esl.Error(err))
			return err
		}
		for _, m := range members {
			row := mo_group_member.NewGroupMember(group, m)
			z.GroupMember.Row(row)
		}
		return nil
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("memberList", memberList)
		q := s.Get("memberList")
		for _, group := range groups {
			q.Enqueue(group)
		}
	})

	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "group_member", func(cols map[string]string) error {
		if _, ok := cols["group_name"]; !ok {
			return errors.New("group_name is not found")
		}
		if _, ok := cols["email"]; !ok {
			return errors.New("email is not found")
		}
		return nil
	})
}
