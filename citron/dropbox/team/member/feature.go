package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_user"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_user"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Feature struct {
	Peer     dbx_conn.ConnScopedTeam
	Features rp_model.RowReport
}

func (z *Feature) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeAccountInfoRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeTeamDataMember,
	)
	z.Features.SetModel(&mo_user.MemberFeature{})
}

func (z *Feature) memberFeature(member *mo_member.Member, c app_control.Control) error {
	l := c.Log().With(esl.String("scanMemberId", member.TeamMemberId), esl.String("scanMemberEmail", member.Email))
	l.Debug("scan")
	feature, err := sv_user.New(z.Peer.Client().AsMemberId(member.TeamMemberId)).Features()
	if err != nil {
		l.Debug("Unable to retrieve member features", esl.Error(err))
		return err
	}

	z.Features.Row(mo_user.NewMemberFeature(member.Email, feature))
	return nil
}

func (z *Feature) Exec(c app_control.Control) error {
	if err := z.Features.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_member", z.memberFeature, c)
		sm := s.Get("scan_member")

		for _, member := range members {
			sm.Enqueue(member)
		}
	})

	return nil
}

func (z *Feature) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Feature{}, rc_recipe.NoCustomValues)
}
