package uc_team_content

import (
	"github.com/watermint/toolbox/domain/common/model/mo_filter"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_worker"
)

type TeamScanner struct {
	Ctx                 dbx_context.Context
	Ctl                 app_control.Control
	TeamOwnedNamespaces map[string]bool
	Scanner             ScanNamespace
	Queue               rc_worker.Queue
	Filter              mo_filter.Filter
}

func (z *TeamScanner) namespacesOfTeam() error {
	l := z.Ctx.Log()

	l.Debug("Scanning admin")
	admin, err := sv_profile.NewTeam(z.Ctx).Admin()
	if err != nil {
		return err
	}
	l = l.With(esl.String("admin", admin.Email))

	l.Debug("Scanning team folders")
	teamfolders, err := sv_teamfolder.New(z.Ctx).List()
	if err != nil {
		return err
	}

	l.Debug("Scanning namespaces")
	namespaces, err := sv_namespace.New(z.Ctx).List()
	if err != nil {
		return err
	}

	l.Debug("Computing duplicates")
	z.TeamOwnedNamespaces = make(map[string]bool)
	teamOwnedNamespaceWithName := make(map[string]string)
	for _, f := range teamfolders {
		if z.Filter.Accept(f.Name) {
			z.TeamOwnedNamespaces[f.TeamFolderId] = true
			teamOwnedNamespaceWithName[f.TeamFolderId] = f.Name
		}
	}
	for _, n := range namespaces {
		if !z.Filter.Accept(n.Name) {
			l.Debug("Skip folder that unmatched to filter condition", esl.String("name", n.Name))
			continue
		}

		switch n.NamespaceType {
		case "app_folder", "team_member_folder":
			l.Debug("Skip non-shared namespace", esl.Any("namespace", n))

		default:
			z.TeamOwnedNamespaces[n.NamespaceId] = true
			teamOwnedNamespaceWithName[n.NamespaceId] = n.Name
		}
	}

	l.Debug("Enqueue to metadata scan")
	for id, name := range teamOwnedNamespaceWithName {
		z.Scanner.Scan(z.Ctl, z.Ctx.AsAdminId(admin.TeamMemberId), name, id)
	}

	l.Debug("Metadata of teams finished")
	return nil
}

func (z *TeamScanner) namespaceOfMember(member *mo_member.Member) error {
	z.Queue.Enqueue(&MemberScannerWorker{
		Member:              member,
		Control:             z.Ctl,
		Context:             z.Ctx.AsMemberId(member.TeamMemberId),
		TeamOwnedNamespaces: z.TeamOwnedNamespaces,
		Scanner:             z.Scanner,
		Folder:              z.Filter,
	})
	return nil
}

func (z *TeamScanner) iterateMembers(f func(member *mo_member.Member) error) error {
	l := z.Ctl.Log()

	if z.TeamOwnedNamespaces == nil {
		l.Debug("Team owned namespaces is not initialized")
		return ErrorTeamOwnedNamespaceIsNotInitialized
	}

	l.Debug("Scanning members")
	members, err := sv_member.New(z.Ctx).List()
	if err != nil {
		return err
	}

	for _, member := range members {
		if err := f(member); err != nil {
			return err
		}
	}
	return nil
}

func (z *TeamScanner) namespacesOfMembers() error {
	return z.iterateMembers(z.namespaceOfMember)
}

func (z *TeamScanner) ScanAll() error {
	if err := z.namespacesOfTeam(); err != nil {
		return err
	}
	if err := z.namespacesOfMembers(); err != nil {
		return err
	}
	return nil
}

// Scan team namespaces
func (z *TeamScanner) ScanTeamOnly() error {
	return z.namespacesOfTeam()
}
