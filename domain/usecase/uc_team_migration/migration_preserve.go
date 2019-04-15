package uc_team_migration

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_group_member"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"go.uber.org/zap"
	"strings"
)

func (z *migrationImpl) Preserve(ctx Context) (err error) {
	// Group to member mapping
	groupToMembers := make(map[string][]*mo_group_member.Member)

	// All groups
	allGroups := make(map[string]*mo_group.Group)

	// Preserve group
	z.log().Info("Preserve: group", zap.Bool("onlyRelated", ctx.GroupsOnlyRelated()))
	preserveGroups := func() error {
		groups, err := sv_group.New(z.ctxMgtSrc).List()
		if err != nil {
			return err
		}

		// fetch group members
		for _, group := range groups {
			members, err := sv_group_member.New(z.ctxMgtSrc, group).List()
			if err != nil {
				return err
			}
			groupToMembers[group.GroupId] = members
			allGroups[group.GroupId] = group

			// filter into members who migrate to dest team
			membersDst := make(map[string]*mo_group_member.Member)
			for _, m := range ctx.Members() {
				for _, gm := range members {
					if gm.TeamMemberId == m.TeamMemberId {
						membersDst[gm.TeamMemberId] = gm
						break
					}
				}
			}

			migrate := false
			if ctx.GroupsOnlyRelated() {
				// ensure any of group member is in migration target members
				if len(membersDst) > 0 && group.GroupManagementType != "system_managed" {
					z.log().Info("Group added because of at least one member associated with the group", zap.String("groupId", group.GroupId), zap.String("groupName", group.GroupName), zap.Int("numberOfMembers", len(membersDst)))
					migrate = true
				}
			} else {
				migrate = true
			}

			if migrate {
				// preserve group
				ctx.AddGroup(group)

				// preserve group members
				for _, m := range membersDst {
					ctx.AddGroupMember(group, m)
				}
			}
		}

		return nil
	}
	if err = preserveGroups(); err != nil {
		return err
	}

	// Preserve shared folders & members
	z.log().Info("Preserve: shared folders & members")
	preserveSharedFolders := func() error {
		// sharedFolderId to teamMemberId
		folderToMember := make(map[string]string)

		// fetch all shared folders of migrating members
		for _, member := range ctx.Members() {
			ctxFileOfMember := z.ctxFileSrc.AsMemberId(member.TeamMemberId)
			folders, err := sv_sharedfolder.New(ctxFileOfMember).List()
			if err != nil {
				return err
			}
			for _, folder := range folders {
				// Access type must above `editor`
				if folder.AccessType != sv_sharedfolder_member.LevelViewer &&
					folder.AccessType != sv_sharedfolder_member.LevelViewerNoComment {
					z.log().Debug("Preserve shared folder", zap.String("Id", folder.SharedFolderId), zap.String("Name", folder.Name), zap.String("Access", folder.AccessType))
					ctx.AddSharedFolder(folder)
					folderToMember[folder.SharedFolderId] = member.TeamMemberId
				}
			}
		}
		return nil
	}
	if err = preserveSharedFolders(); err != nil {
		return err
	}

	// Preserve namespaces
	z.log().Info("Preserve: namespaces")
	preserveNamespaces := func() error {
		namespaces, err := sv_namespace.New(z.ctxFileSrc).List()
		if err != nil {
			return err
		}

		for _, namespace := range namespaces {
			ctx.AddNamespace(namespace)
		}

		return nil
	}
	if err = preserveNamespaces(); err != nil {
		return nil
	}

	// Preserve namespace members
	z.log().Info("Preserve: namespace members")
	preserveNamespaceMembers := func() error {
		ctxFileSrcAdmin := z.ctxFileSrc.AsAdminId(ctx.AdminSrc().TeamMemberId)
		for _, namespace := range ctx.Namespaces() {
			// Skip personal folders
			if namespace.NamespaceType == "app_folder" ||
				namespace.NamespaceType == "team_member_folder" {
				continue
			}

			members, err := sv_sharedfolder_member.NewBySharedFolderId(ctxFileSrcAdmin, namespace.NamespaceId).List()
			if err != nil {
				return err
			}

			for _, member := range members {
				ctx.AddNamespaceMember(namespace, member)

			}
		}
		return nil
	}
	if err = preserveNamespaceMembers(); err != nil {
		return err
	}

	// Preserve group which appear from folder permission
	z.log().Info("Preserve: group from folder permission")
	preserveGroupFromPermissionNamespace := func() error {
		for _, sf := range ctx.NamespaceDetails() {
			if sf.IsTeamFolder || sf.IsInsideTeamFolder {
				members := ctx.NamespaceMembers(sf.SharedFolderId)
				for _, member := range members {
					// Preserve group
					if g, e := member.Group(); e {
						if gg, ee := allGroups[g.GroupId]; ee {
							z.log().Info("Group added because of at least one folder associated with the group", zap.String("groupId", gg.GroupId), zap.String("groupName", gg.GroupName))
							ctx.AddGroup(gg)
						}
					}
				}
			}
		}
		return nil
	}
	if err = preserveGroupFromPermissionNamespace(); err != nil {
		return err
	}

	// Preserve group which appear from folder permission
	z.log().Info("Preserve: group from shared folder permission")
	preserveGroupFromPermissionSharedFolder := func() error {
		for _, sf := range ctx.SharedFolders() {
			if sf.IsTeamFolder || sf.IsInsideTeamFolder {
				members := ctx.NamespaceMembers(sf.SharedFolderId)
				for _, member := range members {
					// Preserve group
					if g, e := member.Group(); e {
						if gg, ee := allGroups[g.GroupId]; ee {
							z.log().Info("Group added because of at least one folder associated with the group", zap.String("groupId", gg.GroupId), zap.String("groupName", gg.GroupName))
							ctx.AddGroup(gg)
						}
					}
				}
			}
		}
		return nil
	}
	if err = preserveGroupFromPermissionSharedFolder(); err != nil {
		return err
	}

	// Preserve group members (2nd scan, group appear from namespace members)
	z.log().Info("Preserve: group members")
	preserveGroupMembers2ndScan := func() error {
		targets := ctx.Members()
		for _, group := range ctx.Groups() {
			// fetch group members
			var members []*mo_group_member.Member
			var e bool
			if members, e = groupToMembers[group.GroupId]; !e {
				members, err = sv_group_member.New(z.ctxMgtSrc, group).List()
				if err != nil {
					return err
				}
			}

			// filter & add group members
			for _, member := range members {
				if _, e := targets[member.TeamMemberId]; e {
					ctx.AddGroupMember(group, member)
				}
			}
		}
		return nil
	}
	if err = preserveGroupMembers2ndScan(); err != nil {
		return err
	}

	// Verify any of group already created in dest team or not
	// But do not block operation
	z.log().Info("Preserve: verify groups")
	verifyDestGroups := func() error {
		groupsDst, err := sv_group.New(z.ctxMgtDst).List()
		if err != nil {
			return err
		}

		for _, groupSrc := range ctx.Groups() {
			nameSrcLower := strings.ToLower(groupSrc.GroupName)
			for _, groupDst := range groupsDst {
				if strings.ToLower(groupDst.GroupName) == nameSrcLower {
					z.log().Warn("Group already exists in dest team", zap.String("GroupName", groupSrc.GroupName))
				}
			}
		}
		return nil
	}
	if err = verifyDestGroups(); err != nil {
		return err
	}

	// Nested folder and relative path mapping
	z.log().Info("Preserve: nested folder and relative path")
	nestedFolderToRelativePathMap := func(tf *mo_teamfolder.TeamFolder) error {
		l := z.log().With(zap.String("teamFolder", tf.Name))
		l.Info("Scanning team folder")

		var scanPath func(ctf api_context.Context, path mo_path.Path) (err error)
		scanPath = func(ctf api_context.Context, path mo_path.Path) (err error) {
			l.Debug("Scanning path", zap.String("path", path.Path()))
			err = sv_file.NewFiles(ctf).ListChunked(path, func(entry mo_file.Entry) {
				if f, e := entry.Folder(); e {
					j := gjson.ParseBytes(f.Raw)
					sfId := j.Get("sharing_info.shared_folder_id")
					childPath := path.ChildPath(f.Name())
					if sfId.Exists() {
						ctx.AddNestedFolderPath(tf.Name, childPath.Path(), sfId.String())
					}
				}
			}, sv_file.Recursive())
			if err != nil {
				l.Error("Unable to retrieve file list at path", zap.String("path", path.Path()), zap.Error(err))
				return err
			}
			return nil
		}

		ctf := z.ctxFileSrc.AsAdminId(ctx.AdminSrc().TeamMemberId).WithPath(api_context.Namespace(tf.TeamFolderId))

		return scanPath(ctf, mo_path.NewPath("/"))
	}
	for _, tf := range ctx.TeamFolders() {
		if err = nestedFolderToRelativePathMap(tf); err != nil {
			return err
		}
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return err
	}

	return nil
}
