package namespace

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"strconv"
)

type MemberNamespaceSummary struct {
	Email             string `json:"email"`
	TotalNamespaces   int    `json:"total_namespaces"`
	MountedNamespaces int    `json:"mounted_namespaces"`
	OwnerNamespaces   int    `json:"owner_namespaces"`
	TeamFolders       int    `json:"team_folders"`
	InsideTeamFolders int    `json:"inside_team_folders"`
	ExternalFolders   int    `json:"external_folders"`
	AppFolders        int    `json:"app_folders"`
}

type TeamNamespaceSummary struct {
	NamespaceType  string `json:"namespace_type"`
	NamespaceCount int    `json:"namespace_count"`
}

type TeamFolderSummary struct {
	Name                string `json:"name"`
	NumNamespacesInside int    `json:"num_namespaces_inside"`
}

type FolderWithoutParent mo_sharedfolder.SharedFolder

type Summary struct {
	Peer                dbx_conn.ConnScopedTeam
	Member              rp_model.RowReport
	Team                rp_model.RowReport
	TeamFolder          rp_model.RowReport
	FolderWithoutParent rp_model.RowReport
	SkipMemberSummary   bool
	Namespace           kv_storage.Storage
	AppFolderByMember   kv_storage.Storage
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
	z.TeamFolder.SetModel(&TeamFolderSummary{})
	z.FolderWithoutParent.SetModel(
		&FolderWithoutParent{},
		rp_model.HiddenColumns(),
	)
}

func (z *Summary) scanTeam(dummy string, stage eq_sequence.Stage, c app_control.Control) error {
	l := c.Log()
	svn := sv_namespace.New(z.Peer.Client())
	qn := stage.Get("scan_namespace")

	summaries := make(map[string]int)
	appFoldersByMember := make(map[string]int)

	lastErr := svn.ListEach(func(entry *mo_namespace.Namespace) bool {
		switch entry.NamespaceType {
		case "app_folder", "team_member_folder", "team_member_root":
			if v, ok := summaries[entry.NamespaceType]; ok {
				summaries[entry.NamespaceType] = v + 1
			} else {
				summaries[entry.NamespaceType] = 1
			}

			if entry.NamespaceType == "app_folder" {
				if v, ok := appFoldersByMember[entry.TeamMemberId]; ok {
					appFoldersByMember[entry.TeamMemberId] = v + 1
				} else {
					appFoldersByMember[entry.TeamMemberId] = 1
				}
			}

		default:
			qn.Enqueue(entry)
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

	for teamMemberId, count := range appFoldersByMember {
		err := z.AppFolderByMember.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutString(teamMemberId, strconv.FormatInt(int64(count), 10))
		})
		if err != nil {
			l.Debug("Unable to store the app folder data", esl.Error(err))
			return err
		}
	}
	return nil
}

func (z *Summary) scanNamespace(namespace *mo_namespace.Namespace, admin *mo_profile.Profile, c app_control.Control) error {
	l := c.Log().With(esl.Any("namespace", namespace))
	meta, err := sv_sharedfolder.New(z.Peer.Client().AsAdminId(admin.TeamMemberId)).Resolve(namespace.NamespaceId)
	if err != nil {
		l.Debug("Unable to retrieve namespace metadata", esl.Error(err))
		return err
	}

	return z.Namespace.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(meta.SharedFolderId, meta)
	})
}

func (z *Summary) scanMember(member *mo_member.Member, info *mo_team.Info, c app_control.Control) error {
	l := c.Log().With(esl.String("member", member.Email))
	svs := sv_sharedfolder.New(z.Peer.Client().AsMemberId(member.TeamMemberId))
	namespaces, err := svs.List()
	if err != nil {
		l.Debug("Unable to retrieve namespaces", esl.Error(err))
		return err
	}

	mounts, err := sv_sharedfolder_mount.New(z.Peer.Client().AsMemberId(member.TeamMemberId)).List()
	if err != nil {
		l.Debug("Unable to retrieve mount info", esl.Error(err))
		return err
	}

	var numAppFolders int64
	err = z.AppFolderByMember.View(func(kvs kv_kvs.Kvs) error {
		v, err := kvs.GetString(member.TeamMemberId)
		switch err {
		case kv_kvs.ErrorNotFound:
			return nil
		case nil:
			if numAppFolders, err = strconv.ParseInt(v, 10, 32); err != nil {
				l.Debug("Unable to parse count", esl.Error(err), esl.String("value", v))
				return err
			}
			return nil
		default:
			return err
		}
	})
	if err != nil {
		l.Debug("Unable to retrieve app folder number")
		return err
	}

	summary := MemberNamespaceSummary{
		Email:             member.Email,
		TotalNamespaces:   len(namespaces),
		MountedNamespaces: len(mounts),
		AppFolders:        int(numAppFolders),
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

func (z *Summary) scanTeamFolders(dummy string, info *mo_team.Info, c app_control.Control) error {
	l := c.Log()
	teamfolders, err := sv_teamfolder.New(z.Peer.Client()).List()
	if err != nil {
		l.Debug("Unable to retrieve team folders", esl.Error(err))
		return err
	}

	// Count inside team folder
	countInsideTeamFolder := 0
	isExternalNamespace := make(map[string]bool)
	err = z.Namespace.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&mo_sharedfolder.SharedFolder{}, func(key string, m interface{}) error {
			meta := m.(*mo_sharedfolder.SharedFolder)
			if meta.IsInsideTeamFolder && meta.OwnerTeamId == info.TeamId {
				countInsideTeamFolder++
			}
			isExternalNamespace[meta.SharedFolderId] = meta.OwnerTeamId != info.TeamId
			return nil
		})
	})
	if err != nil {
		l.Debug("Unable to calc inside team folders", esl.Error(err))
		return err
	}

	// namespace_id -> parent namespace_id
	parents := make(map[string]string)
	err = z.Namespace.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&mo_sharedfolder.SharedFolder{}, func(key string, m interface{}) error {
			meta := m.(*mo_sharedfolder.SharedFolder)
			parents[meta.SharedFolderId] = meta.ParentSharedFolderId
			return nil
		})
	})
	if err != nil {
		l.Debug("Unable to calc parent namespaces", esl.Error(err))
		return err
	}

	// namespace_id -> top namespace_id
	tops := make(map[string]string)
	for ns := range parents {
		tops[ns] = ""
		chain := make([]string, 0)
		parent := parents[ns]
		for {
			chain = append(chain, parent)
			current, ok := parents[parent]
			if current == "" || !ok {
				break
			}
			parent = current
		}
		tops[ns] = parent
		l.Debug("Top folder", esl.String("nsid", ns), esl.String("parent", parent), esl.Strings("chain", chain))
	}

	// top namespace_id -> count descendants
	children := make(map[string]int)
	for _, t := range tops {
		if v, ok := children[t]; ok {
			children[t] = v + 1
		} else {
			children[t] = 1
		}
	}

	z.Team.Row(&TeamNamespaceSummary{
		NamespaceType:  "team_folder",
		NamespaceCount: len(teamfolders),
	})
	z.Team.Row(&TeamNamespaceSummary{
		NamespaceType:  "team_folder (inside team folder)",
		NamespaceCount: countInsideTeamFolder,
	})

	// Report team folder descendants
	for _, teamfolder := range teamfolders {
		l.Debug("team folder", esl.Any("teamfolder", teamfolder))
		if v, ok := children[teamfolder.TeamFolderId]; ok {
			z.TeamFolder.Row(&TeamFolderSummary{
				Name:                teamfolder.Name,
				NumNamespacesInside: v,
			})
		} else {
			z.TeamFolder.Row(&TeamFolderSummary{
				Name:                teamfolder.Name,
				NumNamespacesInside: 0,
			})
		}
	}

	// Report top folders not in children list
	activeTeamFolder := make(map[string]bool)
	for _, tf := range teamfolders {
		activeTeamFolder[tf.TeamFolderId] = true
	}
	for c := range children {
		if c == "" {
			continue
		}
		// skip external namespaces
		if v, ok := isExternalNamespace[c]; v && ok {
			continue
		}
		orphaned := true
		for at := range activeTeamFolder {
			if at == c {
				orphaned = false
				break
			}
		}
		if orphaned {
			_ = z.Namespace.View(func(kvs kv_kvs.Kvs) error {
				var meta mo_sharedfolder.SharedFolder
				kvErr := kvs.GetJsonModel(c, &meta)
				switch kvErr {
				case nil:
					l.Debug("Orphaned folder", esl.Any("meta", meta))
					z.FolderWithoutParent.Row(meta)
				case kv_kvs.ErrorNotFound:
					l.Debug("Orphaned (no metadata)", esl.String("namespace", c))
					z.FolderWithoutParent.Row(&mo_sharedfolder.SharedFolder{SharedFolderId: c})
				default:
					l.Error("Unable to retrieve orphaned folder info", esl.Error(err))
					return kvErr
				}
				return nil
			})
		}
	}

	return nil
}

func (z *Summary) scanSharedFolders(dummy string, c app_control.Control) error {
	countSharedFolders := 0
	err := z.Namespace.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&mo_sharedfolder.SharedFolder{}, func(key string, m interface{}) error {
			meta := m.(*mo_sharedfolder.SharedFolder)
			switch {
			case meta.IsTeamFolder, meta.IsInsideTeamFolder:
				return nil
			default:
				countSharedFolders++
				return nil
			}
		})
	})
	if err != nil {
		return err
	}

	z.Team.Row(&TeamNamespaceSummary{
		NamespaceType:  "shared_folder",
		NamespaceCount: countSharedFolders,
	})
	return nil
}

func (z *Summary) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Member.Open(rp_model.ShowReportTitle()); err != nil {
		return err
	}
	if err := z.Team.Open(rp_model.ShowReportTitle()); err != nil {
		return err
	}
	if err := z.TeamFolder.Open(rp_model.ShowReportTitle()); err != nil {
		return err
	}
	if err := z.FolderWithoutParent.Open(rp_model.ShowReportTitle()); err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(z.Peer.Client()).Admin()
	if err != nil {
		l.Debug("Unable to retrieve admin info", esl.Error(err))
		return err
	}

	members, err := sv_member.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	info, err := sv_team.New(z.Peer.Client()).Info()
	if err != nil {
		return err
	}

	c.Sequence().DoThen(func(s eq_sequence.Stage) {
		s.Define("scan_namespace", z.scanNamespace, admin, c)
		s.Define("scan_team", z.scanTeam, s, c)

		qt := s.Get("scan_team")
		qt.Enqueue("")
	}).Do(func(s eq_sequence.Stage) {
		s.Define("scan_member", z.scanMember, info, c)
		s.Define("scan_teamfolders", z.scanTeamFolders, info, c)
		s.Define("scan_sharedfolders", z.scanSharedFolders, c)

		qs := s.Get("scan_sharedfolders")
		qs.Enqueue("")

		qt := s.Get("scan_teamfolders")
		qt.Enqueue("")

		if !z.SkipMemberSummary {
			qm := s.Get("scan_member")
			for _, m := range members {
				qm.Enqueue(m)
			}
		}
	})

	return nil
}

func (z *Summary) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Summary{}, rc_recipe.NoCustomValues)
}
