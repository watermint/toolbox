package namespace

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type MemberNamespaceSummary struct {
	Email             string `json:"email"`
	TotalNamespaces   int    `json:"total_namespaces"`
	MountedNamespaces int    `json:"mounted_namespaces"`
	OwnerNamespaces   int    `json:"owner_namespaces"`
	TeamFolders       int    `json:"team_folders"`
	AppFolders        int    `json:"app_folders"`
	InsideTeamFolders int    `json:"inside_team_folders"`
	ExternalFolders   int    `json:"external_folders"`
}

type TeamNamespaceSummary struct {
	NamespaceType  string `json:"namespace_type"`
	NamespaceCount int    `json:"namespace_count"`
}

type TeamFolderSummary struct {
	Name                          string `json:"name"`
	NumNamespacesInsideTeamFolder int    `json:"numNamespacesInsideTeamFolder"`
}

type Summary struct {
	Peer   dbx_conn.ConnScopedTeam
	Member rp_model.RowReport
	Team   rp_model.RowReport
}

func (z *Summary) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.Member.SetModel(&MemberNamespaceSummary{})
	z.Team.SetModel(&TeamNamespaceSummary{})
}

func (z *Summary) scanTeam(dummy string, c app_control.Control) error {
	l := c.Log()
	svn := sv_namespace.New(z.Peer.Context())

	summaries := make(map[string]int)

	lastErr := svn.ListEach(func(entry *mo_namespace.Namespace) bool {
		if v, ok := summaries[entry.NamespaceType]; ok {
			summaries[entry.NamespaceType] = v + 1
		} else {
			summaries[entry.NamespaceType] = 1
		}
		return true
	})

	if lastErr != nil {
		l.Debug("Error during listing namespaces", esl.Error(lastErr))
		return lastErr
	}
	for nt, count := range summaries {
		z.Team.Row(&TeamNamespaceSummary{
			NamespaceType:  nt,
			NamespaceCount: count,
		})
	}
	return nil
}

func (z *Summary) scanMember(member *mo_member.Member, info *mo_team.Info, c app_control.Control) error {
	l := c.Log().With(esl.String("member", member.Email))
	svs := sv_sharedfolder.New(z.Peer.Context().AsMemberId(member.TeamMemberId))
	namespaces, err := svs.List()
	if err != nil {
		l.Debug("Unable to retrieve namespaces", esl.Error(err))
		return err
	}

	mounts, err := sv_sharedfolder_mount.New(z.Peer.Context().AsMemberId(member.TeamMemberId)).List()
	if err != nil {
		l.Debug("Unable to retrieve mount info", esl.Error(err))
		return err
	}

	summary := MemberNamespaceSummary{
		Email:             member.Email,
		TotalNamespaces:   len(namespaces),
		MountedNamespaces: len(mounts),
		OwnerNamespaces:   0,
		TeamFolders:       0,
		InsideTeamFolders: 0,
	}

	for _, ns := range namespaces {
		if ns.IsTeamFolder {
			summary.TeamFolders++
		}
		if ns.IsInsideTeamFolder {
			summary.InsideTeamFolders++
		}
		if ns.AccessType == "owner" {
			summary.OwnerNamespaces++
		}
		if ns.OwnerTeamId != info.TeamId {
			summary.ExternalFolders++
		}
	}

	z.Member.Row(&summary)
	return nil
}

func (z *Summary) Exec(c app_control.Control) error {
	//if err := z.Member.Open(rp_model.NoConsoleOutput()); err != nil {
	if err := z.Member.Open(); err != nil {
		return err
	}
	if err := z.Team.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	info, err := sv_team.New(z.Peer.Context()).Info()
	if err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_team", z.scanTeam, c)
		s.Define("scan_member", z.scanMember, info, c)

		qt := s.Get("scan_team")
		qm := s.Get("scan_member")
		qt.Enqueue("")

		for _, m := range members {
			qm.Enqueue(m)
		}
	})
	return nil
}

func (z *Summary) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Summary{}, rc_recipe.NoCustomValues)
}
