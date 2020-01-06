package sharedlink

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type ListWorker struct {
	member *mo_member.Member
	conn   api_context.Context
	rep    rp_model.RowReport
	ctl    app_control.Control
}

func (z *ListWorker) Exec() error {
	z.ctl.UI().InfoK("recipe.team.sharedlink.list.scan", app_msg.P{"MemberEmail": z.member.Email})
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
	Peer       rc_conn.ConnBusinessFile
	SharedLink rp_model.RowReport
}

func (z *List) Preset() {
	z.SharedLink.SetModel(&mo_sharedlink.SharedLinkMember{})
}

func (z *List) Exec(c app_control.Control) error {
	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	if err := z.SharedLink.Open(); err != nil {
		return err
	}

	q := c.NewQueue()
	for _, member := range members {
		q.Enqueue(&ListWorker{
			member: member,
			conn:   z.Peer.Context(),
			rep:    z.SharedLink,
			ctl:    c,
		})
	}
	q.Wait()

	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "shared_link", func(cols map[string]string) error {
		if _, ok := cols["shared_link_id"]; !ok {
			return errors.New("`shared_link_id` is not found")
		}
		if _, ok := cols["team_member_id"]; !ok {
			return errors.New("`team_member_id` is not found")
		}
		return nil
	})
}
