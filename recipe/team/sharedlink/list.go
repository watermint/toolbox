package sharedlink

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
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
	member     *mo_member.Member
	conn       dbx_context.Context
	rep        rp_model.RowReport
	ctl        app_control.Control
	visibility string
}

func (z *ListWorker) Exec() error {
	l := z.ctl.Log().With(esl.String("member", z.member.Email))
	z.ctl.UI().Progress(MList.ProgressScan.With("MemberEmail", z.member.Email))
	mc := z.conn.AsMemberId(z.member.TeamMemberId)
	links, err := sv_sharedlink.New(mc).List()
	if err != nil {
		return err
	}
	for _, link := range links {
		lm := mo_sharedlink.NewSharedLinkMember(link, z.member)
		if z.visibility != "" && lm.Visibility != z.visibility {
			l.Debug("Skipped from report", esl.Any("link", lm))
			continue
		}
		z.rep.Row(lm)
	}
	return nil
}

type List struct {
	Peer       dbx_conn.ConnBusinessFile
	SharedLink rp_model.RowReport
	Visibility mo_string.SelectString
}

func (z *List) Preset() {
	z.Visibility.SetOptions(
		"all",
		"all", "public", "team_only", "password", "team_and_password", "shared_folder_only",
	)
	z.SharedLink.SetModel(
		&mo_sharedlink.SharedLinkMember{},
		rp_model.HiddenColumns(
			"shared_link_id",
			"account_id",
			"team_member_id",
		),
	)
}

func (z *List) listMemberLinks(member *mo_member.Member, c app_control.Control) error {
	l := c.Log().With(esl.String("member", member.Email))
	mc := z.Peer.Context().AsMemberId(member.TeamMemberId)
	links, err := sv_sharedlink.New(mc).List()
	if err != nil {
		return err
	}
	for _, link := range links {
		lm := mo_sharedlink.NewSharedLinkMember(link, member)
		if z.Visibility.Value() != "all" && lm.Visibility != z.Visibility.Value() {
			l.Debug("Skipped from report", esl.Any("link", lm))
			continue
		}
		z.SharedLink.Row(lm)
	}
	return nil
}

func (z *List) Exec(c app_control.Control) error {
	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	if err := z.SharedLink.Open(); err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_member", z.listMemberLinks, c)
		q := s.Get("scan_member")
		for _, member := range members {
			q.Enqueue(member)
		}
	})

	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "shared_link", func(cols map[string]string) error {
		if _, ok := cols["name"]; !ok {
			return errors.New("`name` is not found")
		}
		if _, ok := cols["email"]; !ok {
			return errors.New("`email` is not found")
		}
		return nil
	})
}
