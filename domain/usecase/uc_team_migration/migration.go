package uc_team_migration

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_group_member"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"go.uber.org/zap"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Migration interface {
	// Define scope
	Scope(opts ...ScopeOpt) (ctx Context, err error)

	// Resume from preserved state
	Resume(opts ...ResumeOpt) (ctx Context, err error)

	// Preflight check (inspect, preserve)
	Preflight(ctx Context) (err error)

	// Mirror team folders
	Content(ctx Context) (err error)

	// Migration process (inspect, preserve, bridge, transfer, permission, clean up)
	Migrate(ctx Context) (err error)

	// Inspect team status.
	// Ensure both team allow externally sharing shared folders.
	Inspect(ctx Context) (err error)

	// Preserve members, groups, and sharing status.
	Preserve(ctx Context) (err error)

	// Bridge shared folders.
	// Share all shared folders to destination admin.
	Bridge(ctx Context) (err error)

	// Transfer members.
	// Convert accounts into Basic, and invite from destination team.
	Transfer(ctx Context) (err error)

	// Mirror permissions.
	// Create groups, invite members to shared folders or nested folders,
	// leave destination admin from bridged shared folders.
	Permissions(ctx Context) (err error)

	// Cleanup
	Cleanup(ctx Context) (err error)
}

type ResumeOpt func(opt *resumeOpts) *resumeOpts
type resumeOpts struct {
	storagePath string
	ec          *app.ExecContext
}

func ResumeFromPath(path string) ResumeOpt {
	return func(opt *resumeOpts) *resumeOpts {
		opt.storagePath = path
		return opt
	}
}
func ResumeExecContext(ec *app.ExecContext) ResumeOpt {
	return func(opt *resumeOpts) *resumeOpts {
		opt.ec = ec
		return opt
	}
}

func New(ctxExec *app.ExecContext, ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst api_context.Context) Migration {
	return &migrationImpl{
		ctxExec:          ctxExec,
		ctxFileSrc:       ctxFileSrc,
		ctxMgtSrc:        ctxMgtSrc,
		ctxFileDst:       ctxFileDst,
		ctxMgtDst:        ctxMgtDst,
		teamFolderMirror: uc_teamfolder_mirror.New(ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst),
	}
}

type migrationImpl struct {
	ctxExec          *app.ExecContext
	ctxFileSrc       api_context.Context
	ctxFileDst       api_context.Context
	ctxMgtSrc        api_context.Context
	ctxMgtDst        api_context.Context
	teamFolderMirror uc_teamfolder_mirror.TeamFolder
}

func (z *migrationImpl) Resume(opts ...ResumeOpt) (ctx Context, err error) {
	ro := &resumeOpts{}
	for _, o := range opts {
		o(ro)
	}
	ctxImpl := &contextImpl{}

	{
		b, err := ioutil.ReadFile(filepath.Join(ro.storagePath, "context.json"))
		if err != nil {
			z.ctxExec.Log().Error("unable to read stored context", zap.Error(err))
			return nil, err
		}
		err = json.Unmarshal(b, ctxImpl)
		if err != nil {
			z.ctxExec.Log().Error("unable to unmarshal context", zap.Error(err))
			return nil, err
		}
	}

	{
		b, err := ioutil.ReadFile(filepath.Join(ro.storagePath, "namespace_member.json"))
		if err != nil {
			z.ctxExec.Log().Error("unable to read stored context", zap.Error(err))
			return nil, err
		}
		j := gjson.ParseBytes(b)
		ctxImpl.MapNamespaceMember = make(map[string][]mo_sharedfolder_member.Member)
		if j.Exists() && j.IsObject() {
			for k, ja := range j.Map() {
				if ja.IsArray() {
					var members []mo_sharedfolder_member.Member
					members = make([]mo_sharedfolder_member.Member, 0)
					for _, je := range ja.Array() {
						member := &mo_sharedfolder_member.Metadata{}
						if err := api_parser.ParseModel(member, je.Get("Raw")); err != nil {
							z.log().Error("Unable to parse", zap.Error(err), zap.String("entry", je.Raw))
							return nil, err
						}
						members = append(members, member)
					}
					ctxImpl.MapNamespaceMember[k] = members
				}
			}
		}
	}

	{
		tb, err := ioutil.ReadFile(filepath.Join(ro.storagePath, "teamfolder_content.json"))
		if err != nil {
			z.ctxExec.Log().Error("unable to read stored context", zap.Error(err))
			return nil, err
		}
		tmc, err := uc_teamfolder_mirror.UnmarshalContext(tb)
		if err != nil {
			z.ctxExec.Log().Error("unable to read stored context", zap.Error(err))
			return nil, err
		}
		ctxImpl.ctxTeamFolder = tmc
	}

	ctxImpl.init(ro.ec)
	z.ctxExec.Log().Info("Context restored", zap.String("path", ro.storagePath))
	return ctxImpl, nil
}

func (z *migrationImpl) log() *zap.Logger {
	return z.ctxExec.Log()
}

func (z *migrationImpl) Scope(opts ...ScopeOpt) (ctx Context, err error) {
	so := &scopeOpts{
		membersSpecifiedEmail:    make([]string, 0),
		teamFoldersSpecifiedName: make([]string, 0),
	}
	for _, o := range opts {
		o(so)
	}

	z.log().Info("Define scope")

	// Prepare migration context
	ctx = newContext(z.ctxExec)
	ctx.SetGroupsOnlyRelated(so.groupsOnlyRelated)

	// validation
	if so.membersAllExceptAdmin && len(so.membersSpecifiedEmail) > 0 {
		z.log().Warn("Conflicted option `membersAllExceptAdmin` and `membersSpecifiedEmail`")
		return nil, errors.New("conflicted option")
	}
	if !so.membersAllExceptAdmin && len(so.membersSpecifiedEmail) < 1 {
		z.log().Warn("Please specify `memberAllExceptAdmin` or `membersSpecifiedEmail`")
		return nil, errors.New("not enough options")
	}
	if so.teamFoldersAll && len(so.teamFoldersSpecifiedName) > 0 {
		z.log().Warn("Conflicted option `teamFoldersAll` and `teamFoldersSpecifiedName`")
		return nil, errors.New("conflicted option")
	}
	if !so.teamFoldersAll && len(so.teamFoldersSpecifiedName) < 1 {
		z.log().Warn("Please specify `teamFoldersAll` or `teamFoldersSpecifiedName`")
		return nil, errors.New("not enough options")
	}

	// Identify admins
	identifyAdmins := func() error {
		adminSrc, err := sv_profile.NewTeam(z.ctxMgtSrc).Admin()
		if err != nil {
			return err
		}
		adminDst, err := sv_profile.NewTeam(z.ctxMgtDst).Admin()
		if err != nil {
			return err
		}
		z.log().Debug("Admins identified",
			zap.String("srcId", adminSrc.TeamMemberId),
			zap.String("srcEmail", adminSrc.Email),
			zap.String("dstId", adminDst.TeamMemberId),
			zap.String("dstEmail", adminDst.Email),
		)
		ctx.SetAdmins(adminSrc, adminDst)
		return nil
	}
	if err = identifyAdmins(); err != nil {
		return nil, err
	}

	// Define scope of members
	z.log().Info("Define scope: members")
	allMembers, err := sv_member.New(z.ctxMgtSrc).List()
	if err != nil {
		return nil, err
	}
	if so.membersAllExceptAdmin {
		for _, member := range allMembers {
			ctx.AddMember(member.Profile())
		}
	} else if len(so.membersSpecifiedEmail) > 0 {
		err = nil
		for _, email := range so.membersSpecifiedEmail {
			found := false
			emailLower := strings.ToLower(email)
			for _, member := range allMembers {
				if strings.ToLower(member.Email) == emailLower {
					ctx.AddMember(member.Profile())
					found = true
					break
				}
			}
			if !found {
				z.log().Warn("Member not found for email address", zap.String("email", email))
				err = errors.New("member not found")
			}
		}
		if err != nil {
			return nil, err
		}
	}
	if len(ctx.Members()) < 1 {
		z.log().Warn("No members found")
		return nil, errors.New("no member to migrate")
	}
	z.log().Debug("Members to migrate", zap.Int("count", len(ctx.Members())))

	// Define scope of team folders
	z.log().Info("Define scope: team folders")
	allFolders, err := sv_teamfolder.New(z.ctxFileSrc).List()
	if err != nil {
		return nil, err
	}
	if so.teamFoldersAll {
		for _, folder := range allFolders {
			ctx.AddTeamFolder(folder)
		}
	} else if len(so.teamFoldersSpecifiedName) > 0 {
		err = nil
		for _, name := range so.teamFoldersSpecifiedName {
			found := false
			nameLower := strings.ToLower(name)
			for _, folder := range allFolders {
				if strings.ToLower(folder.Name) == nameLower {
					ctx.AddTeamFolder(folder)
					found = true
					break
				}
			}
			if !found {
				z.log().Warn("Team folder not found for name", zap.String("name", name))
				err = errors.New("team folder not found")
			}
		}
		if err != nil {
			return nil, err
		}
	}
	z.log().Debug("Team folders to migrate", zap.Int("count", len(ctx.TeamFolders())))

	// Team folder mirror
	z.log().Info("Define scope: mirroring content of team folders")
	prepTeamFolderMirror := func() error {
		names := make([]string, 0)
		for _, f := range ctx.TeamFolders() {
			names = append(names, f.Name)
		}
		ctxTf, err := z.teamFolderMirror.PartialScope(names)
		if err != nil {
			return err
		}
		ctx.SetContextTeamFolder(ctxTf)
		return nil
	}
	if err = prepTeamFolderMirror(); err != nil {
		return nil, err
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return nil, err
	}

	return ctx, nil
}

func (z *migrationImpl) Preflight(ctx Context) (err error) {
	if err = z.Inspect(ctx); err != nil {
		return err
	}

	if err = z.Preserve(ctx); err != nil {
		return err
	}

	return nil
}

func (z *migrationImpl) Migrate(ctx Context) (err error) {
	if err = z.Bridge(ctx); err != nil {
		return err
	}

	if err = z.Transfer(ctx); err != nil {
		return err
	}

	if err = z.Permissions(ctx); err != nil {
		return err
	}

	if err = z.Cleanup(ctx); err != nil {
		return err
	}

	return nil
}

func (z *migrationImpl) Inspect(ctx Context) (err error) {
	// Inspect team information.
	z.log().Info("Inspect: team information")
	inspectTeams := func() error {
		var inspectErr error
		inspectErr = nil
		infoSrc, err := sv_team.New(z.ctxMgtSrc).Info()
		if err != nil {
			return err
		}
		infoDst, err := sv_team.New(z.ctxMgtDst).Info()
		if err != nil {
			return err
		}
		z.log().Debug("Team info",
			zap.String("srcId", infoSrc.TeamId),
			zap.String("srcName", infoSrc.Name),
			zap.Int("srcLicenses", infoSrc.NumLicensedUsers),
			zap.Int("srcProvisioned", infoSrc.NumProvisionedUsers),
			zap.String("srcPolicySharedFolderMember", infoSrc.PolicySharedFolderMember),
			zap.String("srcPolicySharedFolderJoin", infoSrc.PolicySharedFolderJoin),
			zap.String("dstId", infoDst.TeamId),
			zap.String("dstName", infoDst.Name),
			zap.Int("dstLicenses", infoDst.NumLicensedUsers),
			zap.Int("dstProvisioned", infoDst.NumProvisionedUsers),
			zap.String("dstPolicySharedFolderMember", infoDst.PolicySharedFolderMember),
			zap.String("dstPolicySharedFolderJoin", infoDst.PolicySharedFolderJoin),
		)
		if infoSrc.TeamId == infoDst.TeamId {
			z.log().Warn("Source and destination team are the same team.")
			inspectErr = errors.New("source and destination teams are the same team")
		}
		if infoSrc.PolicySharedFolderMember != "anyone" {
			z.log().Warn("Source team: Shared folder member policy must be `anyone` during migration")
			inspectErr = errors.New("invalid sharing policy")
		}
		if infoSrc.PolicySharedFolderJoin != "from_anyone" {
			z.log().Warn("Source team: Shared folder join policy must be `from_anyone` during migration")
			inspectErr = errors.New("invalid sharing policy")
		}
		if infoDst.PolicySharedFolderMember != "anyone" {
			z.log().Warn("Dest team: Shared folder member policy must be `anyone` during migration")
			inspectErr = errors.New("invalid sharing policy")
		}
		if infoDst.PolicySharedFolderJoin != "from_anyone" {
			z.log().Warn("Dest team: Shared folder join policy must be `from_anyone` during migration")
			inspectErr = errors.New("invalid sharing policy")
		}

		return inspectErr
	}
	if err = inspectTeams(); err != nil {
		return err
	}

	// Inspect team folder mirror
	z.log().Info("Inspect: team folder mirroring")
	if err = z.teamFolderMirror.Inspect(ctx.ContextTeamFolder()); err != nil {
		return err
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return err
	}

	return nil
}

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

	// Store context
	if err = ctx.StoreState(); err != nil {
		return err
	}

	return nil
}

func (z *migrationImpl) Bridge(ctx Context) (err error) {
	// bridge team folders
	if err = z.teamFolderMirror.Bridge(ctx.ContextTeamFolder()); err != nil {
		return err
	}

	isTeamOwnedSharedFolder := func(namespaceId string) (user *mo_sharedfolder_member.User, exist bool) {
		members := ctx.NamespaceMembers(namespaceId)
		for _, member := range members {
			if member.AccessType() == sv_sharedfolder_member.LevelOwner {
				if u, e := member.User(); e {
					return u, u.SameTeam
				}
				if g, e := member.Group(); e {
					z.log().Error("Group should not owner of shared folder", zap.String("groupId", g.GroupId), zap.String("groupName", g.GroupName))
					return nil, false
				}
			}
		}
		return nil, false
	}

	// bridge shared folders
	z.log().Info("Bridge: shared folders")
	bridgeSharedFolders := func() error {
		folderTargets := ctx.SharedFolders()
		for _, namespace := range ctx.Namespaces() {
			// skip team folder
			if namespace.NamespaceType != "shared_folder" {
				continue
			}

			if f, e := folderTargets[namespace.NamespaceId]; e {
				owner, sameTeam := isTeamOwnedSharedFolder(namespace.NamespaceId)
				if !sameTeam {
					z.log().Debug("Skip non team owned shared folder", zap.String("namespaceId", namespace.NamespaceId), zap.String("name", namespace.Name))
					continue
				}

				l := z.log().With(zap.String("SharedFolderId", f.SharedFolderId), zap.String("SharedFolderName", f.Name), zap.String("dstAdminId", ctx.AdminDst().TeamMemberId))

				l.Debug("Bridge shared folder")
				var ctxFileAsMember api_context.Context
				ctxFileAsMember = z.ctxFileSrc.AsMemberId(owner.TeamMemberId)

				// add
				svc := sv_sharedfolder_member.NewBySharedFolderId(ctxFileAsMember, namespace.NamespaceId)
				err = svc.Add(sv_sharedfolder_member.AddByEmail(ctx.AdminDst().Email, sv_sharedfolder_member.LevelEditor), sv_sharedfolder_member.AddCustomMessage(z.ctxExec.Msg("usecase.team.migration.msg.add_shared_folder").T()))

				if err != nil {
					_, err2 := sv_member.New(z.ctxMgtSrc).ResolveByEmail(owner.Email)
					if err2 != nil {
						l.Debug("Skip bridge: assuming the owner already transferred to dest team", zap.String("namespaceId", namespace.NamespaceId), zap.String("name", namespace.Name))
						continue
					}
					l.Warn("Unable to bridge shared folder permission", zap.Error(err))
					return err
				}
			}
		}
		return nil
	}
	if err = bridgeSharedFolders(); err != nil {
		return err
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return err
	}

	return nil
}

func (z *migrationImpl) Content(ctx Context) (err error) {
	// Mirror team folders
	z.log().Info("Content: mirroring team folder contents")
	if err = z.teamFolderMirror.Mirror(ctx.ContextTeamFolder()); err != nil {
		return err
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return err
	}

	return nil
}

func (z *migrationImpl) Transfer(ctx Context) (err error) {
	// Convert accounts into Basic, and invite from new team
	z.log().Info("Transfer: transfer accounts")
	transferAccounts := func() error {
		svmSrc := sv_member.New(z.ctxMgtSrc)
		svmDst := sv_member.New(z.ctxMgtDst)
		for _, member := range ctx.Members() {
			z.log().Info("Transfer: transferring member", zap.String("email", member.Email))
			l := z.log().With(zap.String("teamMemberId", member.TeamMemberId), zap.String("email", member.Email))
			l.Debug("Transferring account")

			ms, err := svmSrc.Resolve(member.TeamMemberId)
			if err != nil {
				if strings.HasPrefix(api_util.ErrorSummary(err), "id_not_found") {
					l.Debug("Skip: assuming the user already transferred")
					continue
				}
				l.Warn("Unable to resolve existing member", zap.Error(err))
				continue
			}
			err = svmSrc.Remove(ms, sv_member.Downgrade())
			if err != nil {
				l.Warn("Unable to downgrade existing member", zap.Error(err))
				continue
			}

			_, err = svmDst.Add(member.Email)
			if err != nil {
				l.Warn("Unable to downgrade existing member", zap.Error(err))
				continue
			}

			// TODO: add role if the member is an admin
		}
		return nil
	}
	if err = transferAccounts(); err != nil {
		return err
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return err
	}

	return nil
}

func (z *migrationImpl) Permissions(ctx Context) (err error) {
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
			dstGrp, e := groupNameToDstGroup[gn]
			l := z.log().With(zap.String("groupName", gn), zap.String("srcGroupId", srcGrp.GroupId))
			if !e {
				l.Error("Unable to find dest group")
				return errors.New("unable to find dest group")
			}

			l = l.With(zap.String("dstGroupId", dstGrp.GroupId))

			members := ctx.GroupMembers(srcGrp) // lookup by src group
			sgm := sv_group_member.New(z.ctxMgtDst, dstGrp)
			for _, member := range members {
				l.Info("Adding member to group", zap.String("member", member.Email))
				_, err = sgm.Add(sv_group_member.ByEmail(member.Email))
				if err != nil {
					// do not abort, just log
					l.Error("Unable to add member to group")
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
		for _, folder := range ctx.NamespaceDetails() {
			if !folder.IsInsideTeamFolder || folder.IsTeamFolder {
				continue
			}

			z.log().Info("Permissions: restore permission of nested folder", zap.String("name", folder.Name))
			members := ctx.NamespaceMembers(folder.SharedFolderId)
			ctf := z.ctxFileDst.AsAdminId(ctx.AdminDst().TeamMemberId)
			for _, member := range members {
				if srcGrp, e := member.Group(); e {
					svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, folder.SharedFolderId)
					if dstGrp, e := srcGroupIdToDstGroup[srcGrp.GroupId]; !e {
						z.log().Error("Unable to find mapping of src-dst group", zap.String("srcGroup", srcGrp.GroupId), zap.String("srcGroupName", srcGrp.GroupName))
						return err
					} else {
						if err = svm.Add(sv_sharedfolder_member.AddByGroup(dstGrp, member.AccessType())); err != nil {
							z.log().Error("Unable to add group to nested folder", zap.String("folderId", folder.SharedFolderId), zap.String("folderName", folder.Name), zap.String("dstGroup", dstGrp.GroupId), zap.String("dstGroupName", dstGrp.GroupName), zap.Error(err))
						}
					}
				}
				if u, e := member.User(); e {
					svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, folder.SharedFolderId)
					if err = svm.Add(sv_sharedfolder_member.AddByEmail(u.Email, member.AccessType())); err != nil {
						z.log().Error("Unable to add member to nested folder", zap.String("folderId", folder.SharedFolderId), zap.String("folderName", folder.Name), zap.String("member", u.Email), zap.Error(err))
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
		for _, folder := range ctx.NamespaceDetails() {
			if folder.IsInsideTeamFolder || folder.IsTeamFolder {
				continue
			}

			z.log().Info("Permissions: restore permission of shared folder", zap.String("name", folder.Name))
			members := ctx.NamespaceMembers(folder.SharedFolderId)
			ctf := z.ctxFileDst.AsMemberId(ctx.AdminDst().TeamMemberId)
			for _, member := range members {
				if srcGrp, e := member.Group(); e {
					svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, folder.SharedFolderId)
					if dstGrp, e := srcGroupIdToDstGroup[srcGrp.GroupId]; !e {
						z.log().Error("Unable to find mapping of src-dst group", zap.String("srcGroup", srcGrp.GroupId), zap.String("srcGroupName", srcGrp.GroupName))
						return err
					} else {
						if err = svm.Add(sv_sharedfolder_member.AddByGroup(dstGrp, member.AccessType())); err != nil {
							z.log().Error("Unable to add group to shared folder", zap.String("folderId", folder.SharedFolderId), zap.String("folderName", folder.Name), zap.String("dstGroup", dstGrp.GroupId), zap.String("dstGroupName", dstGrp.GroupName), zap.Error(err))
						}
					}
				}
				if u, e := member.User(); e {
					svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, folder.SharedFolderId)
					if err = svm.Add(sv_sharedfolder_member.AddByEmail(u.Email, member.AccessType())); err != nil {
						z.log().Error("Unable to add member to shared folder", zap.String("folderId", folder.SharedFolderId), zap.String("folderName", folder.Name), zap.String("member", u.Email), zap.Error(err))
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

func (z *migrationImpl) Cleanup(ctx Context) (err error) {

	return nil
}
