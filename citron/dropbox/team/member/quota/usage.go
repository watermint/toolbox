package quota

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_usage"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_usage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Usage struct {
	Peer     dbx_conn.ConnScopedTeam
	Usage    rp_model.RowReport
	BasePath mo_string.SelectString
}

func (z *Usage) scanMember(member *mo_member.Member, ctl app_control.Control, ctx dbx_client.Client) error {
	l := ctl.Log().With(esl.Any("member", member))
	l.Debug("Scanning")

	usage, err := sv_usage.New(ctx.AsMemberId(member.TeamMemberId, dbx_filesystem.AsNamespaceType(z.BasePath.Value()))).Resolve()
	if err != nil {
		l.Debug("Unable to scan usage data", esl.Error(err))
		return err
	}

	z.Usage.Row(mo_usage.NewMemberUsage(member, usage))
	return nil
}

func (z *Usage) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeAccountInfoRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.Usage.SetModel(&mo_usage.MemberUsage{})
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Usage) Exec(c app_control.Control) error {
	members, err := sv_member.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	if err := z.Usage.Open(); err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_member", z.scanMember, c, z.Peer.Client())
		q := s.Get("scan_member")
		for _, member := range members {
			q.Enqueue(member)
		}
	})
	return nil
}

func (z *Usage) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &Usage{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "usage", func(cols map[string]string) error {
		if _, ok := cols["email"]; !ok {
			return errors.New("`email` is not found")
		}
		if _, ok := cols["used_bytes"]; !ok {
			return errors.New("`used_bytes` is not found")
		}
		return nil
	})
}
