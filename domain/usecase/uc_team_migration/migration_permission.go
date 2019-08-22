package uc_team_migration

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/api/api_util"
	"go.uber.org/zap"
	"strings"
)

func (z *migrationImpl) Permissions(ctx Context, opts ...PermOpt) (err error) {
	// Permission opt
	po := &permOpts{}
	for _, o := range opts {
		o(po)
	}

	// Resolve dest member by source email
	resolveDstMember := func(srcEmail string) (m *mo_member.Member, err error) {
		l := z.log().With(zap.String("srcEmail", srcEmail))
		targetEmail := ""
		if po.emailMappings == nil {
			l.Debug("Use srcEmail as target email")
			targetEmail = srcEmail
		} else {
			if dstEmail, e := po.emailMappings[srcEmail]; e {
				l.Debug("Use mapped email as target email", zap.String("mappedEmail", dstEmail))
				targetEmail = dstEmail
			} else {
				l.Debug("Not found in mapped email. Use srcEmail as target email")
				targetEmail = srcEmail
			}
		}

		// lookup dst member
		l = l.With(zap.String("targetEmail", targetEmail))
		l.Debug("Lookup member")
		member, err := sv_member.New(z.ctxMgtDst).ResolveByEmail(targetEmail)
		if err != nil {
			l.Debug("Unable to lookup member", zap.Error(err))
			return nil, err
		}

		return member, nil
	}

	// group name (lower) to group mapping
	groupNameToSrcGroup := make(map[string]*mo_group.Group)
	groupNameToDstGroup := make(map[string]*mo_group.Group)
	srcGroupIdToDstGroup := make(map[string]*mo_group.Group)

	// create map name to source group
	z.log().Info("Permissions: creating groups")
	sortSourceGroup := func() error {
		for _, group := range ctx.Groups() {
			groupNameToSrcGroup[strings.ToLower(group.GroupName)] = group
		}
		return nil
	}
	if err = sortSourceGroup(); err != nil {
		return err
	}

	// fetch destination groups
	z.log().Info("Permissions: retrieve destination team groups")
	fetchDestGroups := func() error {
		groups, err := sv_group.New(z.ctxMgtDst).List()
		if err != nil {
			return err
		}
		for _, group := range groups {
			ctx.AddDestGroup(group)
			groupNameToDstGroup[strings.ToLower(group.GroupName)] = group
		}
		return nil
	}
	if err = fetchDestGroups(); err != nil {
		return err
	}

	// create group if not exist
	z.log().Info("Permissions: create groups if not exist")
	createDestGroups := func() error {
		for n, src := range groupNameToSrcGroup {
			l := z.log().With(zap.String("groupName", src.GroupName))
			if _, e := groupNameToDstGroup[n]; !e {
				if src.GroupManagementType == "system_managed" {
					l.Debug("Skip system managed group")
					continue
				}
				l.Debug("Creating group in the dest team")
				group, err := sv_group.New(z.ctxMgtDst).Create(src.GroupName, sv_group.ManagementType(src.GroupManagementType))
				if err != nil {
					l.Warn("Unable to create group in the dest team", zap.Error(err))
					return err
				}
				groupNameToDstGroup[n] = group
				ctx.AddDestGroup(group)
			}
		}
		return nil
	}
	if err = createDestGroups(); err != nil {
		return err
	}

	// add members to groups
	z.log().Info("Permissions: add members to groups")
	addMembersToGroups := func() error {
		for gn, srcGrp := range groupNameToSrcGroup {
			l := z.log().With(zap.String("groupName", gn), zap.String("srcGroupId", srcGrp.GroupId))
			if srcGrp.GroupManagementType == "system_managed" {
				l.Debug("Skip: system managed group")
				continue
			}
			dstGrp, e := groupNameToDstGroup[gn]
			if !e {
				l.Error("Unable to find dest group")
				return errors.New("unable to find dest group")
			}

			l = l.With(zap.String("dstGroupId", dstGrp.GroupId))

			members := ctx.GroupMembers(srcGrp) // lookup by src group
			sgm := sv_group_member.New(z.ctxMgtDst, dstGrp)
			for _, member := range members {
				if member.TeamMemberId == ctx.AdminSrc().TeamMemberId {
					l.Debug("Skip: Admin should not be added", zap.String("member", member.Email))
					continue
				}

				// lookup member in dst
				dstMember, err := resolveDstMember(member.Email)
				if err != nil {
					l.Warn("The member not found in dst team", zap.Error(err))
					continue
				}

				l.Info("Adding member to group", zap.Any("member", dstMember))
				_, err = sgm.Add(sv_group_member.ByEmail(dstMember.Email))
				if err != nil {
					if strings.HasPrefix(api_util.ErrorSummary(err), "duplicate_user") {
						l.Debug("The member already added")
					} else {
						l.Error("Unable to add member to group", zap.Error(err))
						//return err
					}
				}
			}
		}
		return nil
	}
	if err = addMembersToGroups(); err != nil {
		return err
	}

	// create map src group id to dst group
	z.log().Info("Permissions: mapping source to destination groups")
	createSrcGroupIdToDstGroupMap := func() error {
		for n, src := range groupNameToSrcGroup {
			if src.GroupManagementType == "system_managed" {
				z.log().Debug("Skip: system managed group", zap.String("groupName", src.GroupName))
				continue
			}
			if dst, e := groupNameToDstGroup[n]; e {
				srcGroupIdToDstGroup[src.GroupId] = dst
			} else {
				// should not happen
				z.log().Warn("Unable to find dst group", zap.String("groupName", src.GroupName))
				return errors.New("unable to find dst group")
			}
		}
		return nil
	}
	if err = createSrcGroupIdToDstGroupMap(); err != nil {
		return err
	}

	// create map of name to dest team folders
	nameToDestTeamFolders := make(map[string]*mo_teamfolder.TeamFolder)
	createDestTeamFolderMap := func() error {
		folders, err := sv_teamfolder.New(z.ctxFileDst).List()
		if err != nil {
			z.log().Error("Unable to resolve dest team folders", zap.Error(err))
			return err
		}
		for _, folder := range folders {
			nameToDestTeamFolders[strings.ToLower(folder.Name)] = folder
		}
		return nil
	}
	if err = createDestTeamFolderMap(); err != nil {
		return err
	}

	// create team folder if it's not exist
	z.log().Info("Permission: create team folder(s) if not exist in dest team")
	createTeamFolderIfNotExist := func() error {
		for _, stf := range ctx.TeamFolders() {
			l := z.log().With(zap.String("teamFolderName", stf.Name), zap.String("srcTeamFolderId", stf.TeamFolderId))
			if _, e := nameToDestTeamFolders[strings.ToLower(stf.Name)]; !e {
				svt := sv_teamfolder.New(z.ctxFileDst)
				l.Info("Creating team folder")
				dtf, err := svt.Create(stf.Name, sv_teamfolder.SyncNoSync())
				if err != nil {
					l.Error("Unable to create team folder", zap.Error(err))
					return err
				}
				nameToDestTeamFolders[strings.ToLower(dtf.Name)] = dtf
			}
		}
		return nil
	}
	if err = createTeamFolderIfNotExist(); err != nil {
		return err
	}

	// restore permission for team folders
	z.log().Info("Permissions: restore permission of team folders")
	restorePermissionTeamFolder := func() error {
		for _, stf := range ctx.TeamFolders() {
			l := z.log().With(zap.String("teamFolderName", stf.Name), zap.String("srcTeamFolderId", stf.TeamFolderId))
			// resolve team folder in dst
			if dtf, e := nameToDestTeamFolders[strings.ToLower(stf.Name)]; !e {
				l.Error("Unable to find dest team folder")
			} else {
				l.Info("Permissions: restore permission of team folder")
				members := ctx.NamespaceMembers(stf.TeamFolderId)
				ctf := z.ctxFileDst.AsAdminId(ctx.AdminDst().TeamMemberId)
				for _, member := range members {
					if srcGrp, e := member.Group(); e {
						svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, dtf.TeamFolderId)
						if dstGrp, e := srcGroupIdToDstGroup[srcGrp.GroupId]; !e {
							l.Error("Unable to find mapping of src-dst group", zap.String("srcGroup", srcGrp.GroupId), zap.String("srcGroupName", srcGrp.GroupName))
							return err
						} else {
							if err = svm.Add(sv_sharedfolder_member.AddByGroup(dstGrp, member.AccessType())); err != nil {
								l.Error("Unable to add group to team folder", zap.String("destTeamFolderId", dtf.TeamFolderId), zap.String("dstGroup", dstGrp.GroupId), zap.String("dstGroupName", dstGrp.GroupName), zap.Error(err))
							}
						}
					}
					if u, e := member.User(); e {
						l.Error("Team folder should not have individual sharing member", zap.String("destTeamFolderId", dtf.TeamFolderId), zap.String("member", u.Email))
					}
				}
			}
		}
		return nil
	}
	if err = restorePermissionTeamFolder(); err != nil {
		return err
	}

	// restore permissions for nested folders
	z.log().Info("Permissions: restore permission of nested folders")
	restorePermissionNestedFolder := func() error {
		dstFolders, err := sv_teamfolder.New(z.ctxFileDst).List()
		if err != nil {
			z.log().Error("Unable to retrieve team folders of dest team", zap.Error(err))
			return err
		}
		nameToDestTeamFolder := make(map[string]*mo_teamfolder.TeamFolder)
		for _, f := range dstFolders {
			nameToDestTeamFolder[strings.ToLower(f.Name)] = f
		}

		for _, srcTeamFolder := range ctx.TeamFolders() {
			l := z.log().With(zap.String("teamFolderName", srcTeamFolder.Name))
			dstTeamFolder, e := nameToDestTeamFolder[strings.ToLower(srcTeamFolder.Name)]
			if !e {
				l.Warn("Unable to find src-dst team folder map")
				return errors.New("unable to find src-dst map")
			}

			nestedFolders, e := ctx.NestedFolderPath(srcTeamFolder.Name)
			if !e {
				l.Info("Nested folders not found for team folder")
				continue
			}

			for relPath, srcNestedFolderId := range nestedFolders {
				ll := l.With(zap.String("dstTeamFolderId", dstTeamFolder.TeamFolderId), zap.String("relPath", relPath), zap.String("srcNestedFolderId", srcNestedFolderId))
				ll.Info("Restore permissions of nested folder")

				// create nested folder
				svs := sv_sharedfolder.New(z.ctxFileDst.AsAdminId(ctx.AdminDst().TeamMemberId).WithPath(api_context.Namespace(dstTeamFolder.TeamFolderId)))
				nf, err := svs.Create(mo_path.NewPath(relPath))
				if err != nil {
					if strings.Contains(api_util.ErrorSummary(err), "bad_path/already_shared") {
						ll.Debug("Skip: already shared")
						eb := api_util.ErrorBody(err)
						if eb == nil {
							ll.Error("Unable to verify nested folder", zap.Error(err))
							continue
						}
						j := gjson.ParseBytes(eb)
						badPath := j.Get("bad_path")
						nf = &mo_sharedfolder.SharedFolder{}
						if err = api_parser.ParseModel(nf, badPath); err != nil {
							ll.Error("Unable to verify nested folder", zap.Error(err))
							continue
						}
						ll.Debug("Nested folder", zap.String("namespaceId", nf.SharedFolderId))

					} else {
						ll.Error("Unable to create nested folder", zap.Error(err))
						continue
					}
				}

				ll.Info("Permissions: restore permission of nested folder")
				members := ctx.NamespaceMembers(srcNestedFolderId) // get members from src namespace id
				ctf := z.ctxFileDst.AsAdminId(ctx.AdminDst().TeamMemberId)
				for _, member := range members {
					if srcGrp, e := member.Group(); e {
						svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, nf.SharedFolderId) // add member by new nested folder namespace id
						if dstGrp, e := srcGroupIdToDstGroup[srcGrp.GroupId]; !e {
							ll.Error("Unable to find mapping of src-dst group", zap.String("srcGroup", srcGrp.GroupId), zap.String("srcGroupName", srcGrp.GroupName))
							return err
						} else {
							if err = svm.Add(sv_sharedfolder_member.AddByGroup(dstGrp, member.AccessType())); err != nil {
								ll.Error("Unable to add group to nested folder", zap.String("folderId", nf.SharedFolderId), zap.String("dstGroup", dstGrp.GroupId), zap.String("dstGroupName", dstGrp.GroupName), zap.Error(err))
							}
						}
					}
					if u, e := member.User(); e {
						// lookup dst member
						dstMember, err := resolveDstMember(u.Email)
						if err != nil {
							ll.Debug("The member not found. Skip", zap.Error(err))
							continue
						}

						svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, nf.SharedFolderId)
						if err = svm.Add(sv_sharedfolder_member.AddByEmail(dstMember.Email, member.AccessType())); err != nil {
							ll.Error("Unable to add member to nested folder", zap.Any("folder", nf), zap.Any("member", dstMember), zap.Error(err))
						}
					}
					if inv, e := member.Invitee(); e {
						// lookup dst member
						dstMember, err := resolveDstMember(inv.InviteeEmail)
						if err != nil {
							ll.Debug("The member not found. Invite via inviteeEmail", zap.Any("invitee", inv), zap.Error(err))

							svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, nf.SharedFolderId)
							if err = svm.Add(sv_sharedfolder_member.AddByEmail(inv.InviteeEmail, member.AccessType())); err != nil {
								ll.Error("Unable to add member to nested folder", zap.Any("folder", nf), zap.Any("invitee", inv), zap.Error(err))
							}
						} else {
							ll.Debug("Invite via mapped user", zap.Any("invitee", inv), zap.Any("dstMember", dstMember))

							svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, nf.SharedFolderId)
							if err = svm.Add(sv_sharedfolder_member.AddByEmail(dstMember.Email, member.AccessType())); err != nil {
								ll.Error("Unable to add member to nested folder", zap.Any("folder", nf), zap.Any("member", dstMember), zap.Error(err))
							}
						}
					}
				}
			}
		}
		return nil
	}
	if err = restorePermissionNestedFolder(); err != nil {
		return err
	}

	// restore permissions for shared folders
	z.log().Info("Permissions: restore permission of shared folders")
	restorePermissionSharedFolder := func() error {
		for _, folder := range ctx.SharedFolders() {
			l := z.log().With(zap.String("name", folder.Name))
			if folder.IsTeamFolder || folder.IsInsideTeamFolder {
				l.Debug("Skip team folder & nested folder")
				continue
			}
			owner, sameTeam := z.isTeamOwnedSharedFolder(ctx, folder.SharedFolderId)
			if !sameTeam {
				l.Debug("Skip non team owned folder")
				continue
			}
			if owner.TeamMemberId == ctx.AdminSrc().TeamMemberId ||
				owner.Email == ctx.AdminSrc().Email {
				l.Debug("Skip shared folder which owned by src admin")
				continue
			}
			var ctf api_context.Context
			ownerMember, err := sv_member.New(z.ctxMgtDst).ResolveByEmail(owner.Email)
			if err != nil {
				l.Debug("Unable to resolve folder owner user in dest team", zap.Any("owner", owner), zap.Error(err))

				// Then try from src team (assuming the user exists in src team because of emailMapping enabled)
				ownerMember, err = sv_member.New(z.ctxMgtSrc).ResolveByEmail(owner.Email)
				if err != nil {
					l.Error("Unable to resolve folder owner in both src, dest team", zap.Any("owner", owner), zap.Error(err))
					return err
				}
				l.Debug("Owner found in src team", zap.Any("ownerMember", ownerMember))
				ctf = z.ctxFileSrc.AsMemberId(ownerMember.TeamMemberId)
			} else {
				ctf = z.ctxFileDst.AsMemberId(ctx.AdminDst().TeamMemberId)
			}

			l.Info("Permissions: restore permission of shared folder")

			members := ctx.NamespaceMembers(folder.SharedFolderId)
			svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, folder.SharedFolderId)
			dstMembers, err := svm.List()
			if err != nil {
				l.Error("Unable to list dst shared folder members", zap.Error(err))
				return err
			}

			for _, member := range members {
				if srcGrp, e := member.Group(); e {
					if dstGrp, e := srcGroupIdToDstGroup[srcGrp.GroupId]; !e {
						l.Error("Unable to find mapping of src-dst group", zap.String("srcGroup", srcGrp.GroupId), zap.String("srcGroupName", srcGrp.GroupName))
						return err
					} else {
						found := false
						for _, dstMember := range dstMembers {
							if dg, e := dstMember.Group(); e {
								if dstGrp.GroupId == dg.GroupId && dstMember.AccessType() == dg.EntryAccessType {
									l.Debug("Skip: dst group already added to shared folder", zap.String("srcGroup", srcGrp.GroupId), zap.String("dstGroup", dstGrp.GroupId), zap.String("groupName", dstGrp.GroupName), zap.String("accessType", dg.AccessType()))
									found = true
									break
								}
							}
						}
						if found {
							continue
						}

						if err = svm.Add(sv_sharedfolder_member.AddByGroup(dstGrp, member.AccessType())); err != nil {
							l.Error("Unable to add group to shared folder", zap.String("folderId", folder.SharedFolderId), zap.String("dstGroup", dstGrp.GroupId), zap.String("dstGroupName", dstGrp.GroupName), zap.Error(err))
							return err
						}
					}
				}
				if u, e := member.User(); e {
					accessType := member.AccessType()
					if accessType == sv_sharedfolder_member.LevelOwner {
						accessType = sv_sharedfolder_member.LevelEditor
					}

					// lookup dst member
					dstMember, err := resolveDstMember(u.Email)
					if err != nil {
						l.Debug("Skip not found member", zap.Error(err))
						continue
					}

					found := false
					for _, dm := range dstMembers {
						if du, e := dm.User(); e {
							if du.Email == dstMember.Email && du.AccessType() == u.AccessType() {
								l.Debug("Skip: user already added to shared folder", zap.Any("user", u), zap.Any("dstMember", dstMember), zap.String("accessType", u.AccessType()))
								found = true
								break
							}
						}
					}
					if found {
						continue
					}

					if err = svm.Add(sv_sharedfolder_member.AddByEmail(dstMember.Email, accessType)); err != nil {
						l.Error("Unable to add member to shared folder", zap.String("folderId", folder.SharedFolderId), zap.Any("user", u), zap.Any("dstMember", dstMember), zap.Error(err))
						return err
					}
				}
				if inv, e := member.Invitee(); e {
					// lookup dst member
					dstMember, err := resolveDstMember(inv.InviteeEmail)
					if err != nil {
						l.Debug("Skip not found member", zap.Error(err))
						continue
					}

					found := false
					for _, dm := range dstMembers {
						if di, e := dm.Invitee(); e {
							if di.InviteeEmail == dstMember.Email && di.AccessType() == inv.AccessType() {
								l.Debug("Skip: invitee already added to shared folder", zap.String("invitee", inv.InviteeEmail), zap.String("accessType", inv.AccessType()))
								found = true
								break
							}
						}
					}
					if found {
						continue
					}

					if err = svm.Add(sv_sharedfolder_member.AddByEmail(dstMember.Email, inv.AccessType())); err != nil {
						l.Error("Unable to add invitee to shared folder", zap.String("folderId", folder.SharedFolderId), zap.String("member", inv.InviteeEmail), zap.Error(err))
						return err
					}
				}
			}

			// transfer ownership
			if ctx.DontTransferFolderOwnership() {
				l.Debug("Skip transfer ownership")
			} else {
				err = sv_sharedfolder.New(ctf).Transfer(folder, sv_sharedfolder.ToTeamMemberId(ownerMember.TeamMemberId))
				if err != nil {
					if strings.HasPrefix(api_util.ErrorSummary(err), "new_owner_not_a_member") {
						l.Debug("Unable to restore due to original owner not yet activated", zap.String("originalOwner", ownerMember.Email), zap.Error(err))
					} else {
						l.Warn("Unable to restore ownership", zap.String("originalOwner", ownerMember.Email), zap.Error(err))
					}
				}
			}
		}
		return nil
	}
	if err = restorePermissionSharedFolder(); err != nil {
		return err
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return err
	}

	return nil
}
