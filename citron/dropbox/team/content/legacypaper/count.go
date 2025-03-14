package legacypaper

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_paper"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type PaperCount struct {
	MemberEmail string `json:"member_email"`
	Created     int    `json:"created"`
	Accessed    int    `json:"accessed"`
}

type Count struct {
	Peer     dbx_conn.ConnScopedTeam
	Stats    rp_model.RowReport
	BasePath mo_string.SelectString
}

func (z *Count) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeTeamDataMember,
	)
	z.Stats.SetModel(&PaperCount{})
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Count) countMember(member *mo_member.Member, c app_control.Control) error {
	l := c.Log().With(esl.String("memberEmail", member.Email))

	pc := PaperCount{
		MemberEmail: member.Email,
		Created:     0,
		Accessed:    0,
	}

	mc := z.Peer.Client().
		AsMemberId(member.TeamMemberId, dbx_filesystem.AsNamespaceType(z.BasePath.Value())).
		WithPath(dbx_client.Namespace(member.Profile().RootNamespaceId))
	err := sv_paper.NewLegacy(mc).ListCreated(func(docId string) {
		pc.Created++
	})
	if err != nil {
		l.Debug("Unable to retrieve list created", esl.Error(err))
		return err
	}
	err = sv_paper.NewLegacy(mc).ListAccessed(func(docId string) {
		pc.Accessed++
	})
	if err != nil {
		l.Debug("Unable to retrieve list accessed", esl.Error(err))
		return err
	}

	z.Stats.Row(&pc)
	return nil
}

func (z *Count) Exec(c app_control.Control) error {
	if err := z.Stats.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_member", z.countMember, c)
		scan := s.Get("scan_member")

		for _, member := range members {
			scan.Enqueue(member)
		}
	})

	return nil
}

func (z *Count) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Count{}, rc_recipe.NoCustomValues)
}
