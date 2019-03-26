package uc_teamfolder_mirror

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_profile"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/domain/usecase/uc_file_compare"
	"github.com/watermint/toolbox/domain/usecase/uc_file_mirror"
	"go.uber.org/zap"
	"strings"
	"time"
)

type TeamFolder interface {
	// All team folder scope
	AllFolderScope() (ctx Context, err error)

	// Specific team folders.
	PartialScope(names []string) (ctx Context, err error)

	// Mirror
	Mirror(ctx Context) (err error)

	// Inspect team folder.
	Inspect(ctx Context) (err error)

	// Create group to bridge permissions
	Bridge(ctx Context) (err error)

	// Mount, or create dest team folder if required.
	Mount(ctx Context, scope Scope) (err error)

	// Mirror contents
	Content(ctx Context, scope Scope) (err error)

	// Verify contents
	Verify(ctx Context, scope Scope) (err error)

	// Unmount
	Unmount(ctx Context, scope Scope) (err error)

	// Clean up permissions which used for mirroring
	Cleanup(ctx Context) (err error)
}

type MirrorPair struct {
	Src *mo_teamfolder.TeamFolder
	Dst *mo_teamfolder.TeamFolder
}

type Scope interface {
	Pair() (pair *MirrorPair)
}

func NewScope(pair *MirrorPair) Scope {
	return &scopeImpl{
		pair: pair,
	}
}

type scopeImpl struct {
	pair *MirrorPair
}

func (z *scopeImpl) Pair() (pair *MirrorPair) {
	return z.pair
}

// Mutable state of mirroring.
type Context interface {
	Pairs() (pairs []*MirrorPair)
	SetGroups(src, dst *mo_group.Group)
	SetAdmins(src, dst *mo_profile.Profile)
	GroupSrc() *mo_group.Group
	GroupDst() *mo_group.Group
	AdminSrc() *mo_profile.Profile
	AdminDst() *mo_profile.Profile
}

type mirrorContext struct {
	pairs    []*MirrorPair
	groupSrc *mo_group.Group
	groupDst *mo_group.Group
	adminSrc *mo_profile.Profile
	adminDst *mo_profile.Profile
}

func (z *mirrorContext) SetGroups(src, dst *mo_group.Group) {
	z.groupSrc = src
	z.groupDst = dst
}

func (z *mirrorContext) SetAdmins(src, dst *mo_profile.Profile) {
	z.adminSrc = src
	z.adminDst = dst
}

func (z *mirrorContext) Pairs() (pairs []*MirrorPair) {
	return z.pairs
}

func (z *mirrorContext) GroupSrc() *mo_group.Group {
	return z.groupSrc
}

func (z *mirrorContext) GroupDst() *mo_group.Group {
	return z.groupDst
}

func (z *mirrorContext) AdminSrc() *mo_profile.Profile {
	return z.adminSrc
}

func (z *mirrorContext) AdminDst() *mo_profile.Profile {
	return z.adminDst
}

func New(ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst api_context.Context) TeamFolder {
	return &teamFolderImpl{
		ctxFileSrc: ctxFileSrc,
		ctxMgtSrc:  ctxMgtSrc,
		ctxFileDst: ctxFileDst,
		ctxMgtDst:  ctxMgtDst,
	}
}

type teamFolderImpl struct {
	ctxFileSrc api_context.Context
	ctxFileDst api_context.Context
	ctxMgtSrc  api_context.Context
	ctxMgtDst  api_context.Context
}

func (z *teamFolderImpl) log() *zap.Logger {
	return z.ctxFileSrc.Log()
}

func (z *teamFolderImpl) AllFolderScope() (ctx Context, err error) {
	mc := &mirrorContext{
		pairs: make([]*MirrorPair, 0),
	}
	ctx = mc
	svt := sv_teamfolder.New(z.ctxFileSrc)
	folders, err := svt.List()
	if err != nil {
		return nil, err
	}
	for _, folder := range folders {
		mc.pairs = append(mc.pairs, &MirrorPair{
			Src: folder,
			Dst: nil,
		})
	}
	return
}

func (z *teamFolderImpl) PartialScope(names []string) (ctx Context, err error) {
	mc := &mirrorContext{
		pairs: make([]*MirrorPair, 0),
	}
	ctx = mc
	svt := sv_teamfolder.New(z.ctxFileSrc)
	folders, err := svt.List()
	if err != nil {
		return nil, err
	}
	matches := func(fnl string) bool {
		for _, name := range names {
			if strings.ToLower(name) == fnl {
				return true
			}
		}
		return false
	}
	for _, folder := range folders {
		fnl := strings.ToLower(folder.Name)
		if matches(fnl) {
			mc.pairs = append(mc.pairs, &MirrorPair{
				Src: folder,
				Dst: nil,
			})
		}
	}
	return
}

func (z *teamFolderImpl) Mirror(ctx Context) (err error) {
	if err = z.Inspect(ctx); err != nil {
		return err
	}
	var lastErr error
	if err = z.Bridge(ctx); err != nil {
		lastErr = err
	}
	for _, pair := range ctx.Pairs() {
		scope := NewScope(pair)

		if err = z.Mount(ctx, scope); err != nil {
			lastErr = err
			continue
		}
		if err = z.Content(ctx, scope); err != nil {
			lastErr = err
		} else {
			if err = z.Verify(ctx, scope); err != nil {
				lastErr = err
			}
		}
		if err = z.Unmount(ctx, scope); err != nil {
			lastErr = err
		}
	}
	if err = z.Cleanup(ctx); err != nil {
		lastErr = err
	}
	return lastErr
}

func (z *teamFolderImpl) Inspect(ctx Context) (err error) {
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
		return err
	}

	// Inspect team information.
	inspectTeams := func() error {
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
			zap.String("dstId", infoDst.TeamId),
			zap.String("dstName", infoDst.Name),
			zap.Int("dstLicenses", infoDst.NumLicensedUsers),
			zap.Int("dstProvisioned", infoDst.NumProvisionedUsers),
		)
		if infoSrc.TeamId == infoDst.TeamId {
			z.log().Warn("Source and destination team are the same team.")
			return errors.New("source and destination teams are the same team")
		}
		return nil
	}
	if err = inspectTeams(); err != nil {
		return err
	}

	// Inspect src folders
	inspectSrcFolders := func() error {
		var inspectErr error
		for _, pair := range ctx.Pairs() {
			z.log().Info("SRC: Team folder status",
				zap.String("id", pair.Src.TeamFolderId),
				zap.String("name", pair.Src.Name),
				zap.String("status", pair.Src.Status),
			)
			if pair.Src.Status != "active" {
				z.log().Info("SRC: Non active folder found",
					zap.String("srcId", pair.Src.TeamFolderId),
					zap.String("srcName", pair.Src.Name),
					zap.String("srcStatus", pair.Src.Status),
				)
				inspectErr = errors.New("one or more team folders are not active")
			}
		}
		return inspectErr
	}
	if err := inspectSrcFolders(); err != nil {
		return err
	}

	// retrieve destination folders
	svt := sv_teamfolder.New(z.ctxFileDst)
	folders, err := svt.List()
	if err != nil {
		return err
	}

	// find dst folder
	reduceFolder := func(name string) *mo_teamfolder.TeamFolder {
		nameLower := strings.ToLower(name)
		for _, folder := range folders {
			if strings.ToLower(folder.Name) == nameLower {
				return folder
			}
		}
		return nil
	}

	// Match pair
	for _, pair := range ctx.Pairs() {
		if dstFolder := reduceFolder(pair.Src.Name); dstFolder != nil {
			pair.Dst = dstFolder
		}
	}

	// Inspect dest folders
	inspectDstFolders := func() error {
		var inspectErr error
		for _, pair := range ctx.Pairs() {
			if folder := pair.Dst; folder != nil {
				z.log().Info("DST: Team folder status",
					zap.String("srcId", pair.Src.TeamFolderId),
					zap.String("srcName", pair.Src.Name),
					zap.String("srcStatus", pair.Src.Status),
					zap.String("dstId", folder.TeamFolderId),
					zap.String("dstName", folder.Name),
					zap.String("dstStatus", folder.Status),
				)
				if pair.Dst.Status != "active" {
					z.log().Info("DST: Non active folder found",
						zap.String("dstId", folder.TeamFolderId),
						zap.String("dstName", folder.Name),
						zap.String("dstStatus", folder.Status),
					)
					inspectErr = errors.New("one or more team folders are not active")
				}
			}
		}
		return inspectErr
	}
	if err = inspectDstFolders(); err != nil {
		return err
	}

	return nil
}

func (z *teamFolderImpl) Bridge(ctx Context) (err error) {
	groupName := fmt.Sprintf("toolbox-teamfolder-mirror-%x", time.Now().Unix())

	// Create groups
	groupSrc, err := sv_group.New(z.ctxMgtSrc).CreateCompanyManaged(groupName)
	if err != nil {
		return err
	}
	ctx.SetGroups(groupSrc, nil)

	groupDst, err := sv_group.New(z.ctxMgtDst).CreateCompanyManaged(groupName)
	if err != nil {
		return err
	}
	ctx.SetGroups(groupSrc, groupDst)
	z.log().Debug("Groups created", zap.String("srcGroupId", groupSrc.GroupId), zap.String("dstGroupId", groupDst.GroupId), zap.String("groupName", groupName))

	// Add admins to groups
	_, err = sv_group_member.New(z.ctxMgtSrc, groupSrc).Add([]string{ctx.AdminSrc().TeamMemberId})
	if err != nil {
		return err
	}
	_, err = sv_group_member.New(z.ctxMgtDst, groupDst).Add([]string{ctx.AdminDst().TeamMemberId})
	if err != nil {
		return err
	}
	z.log().Debug("Admins added to groups", zap.String("srcGroupId", groupSrc.GroupId), zap.String("dstGroupId", groupDst.GroupId), zap.String("groupName", groupName))

	return nil
}

func (z *teamFolderImpl) Mount(ctx Context, scope Scope) (err error) {
	// Create team folder if required
	createIfRequired := func() error {
		var createErr error
		svt := sv_teamfolder.New(z.ctxMgtDst)
		pair := scope.Pair()
		if pair.Dst == nil {
			folder, err := svt.Create(pair.Src.Name)
			if err != nil {
				z.log().Warn("DST: Unable to create team folder", zap.String("name", pair.Src.Name), zap.Error(err))
				createErr = errors.New("could not create one or more team folders in the destination team")
			}
			z.log().Debug("DST: Team folder created", zap.String("id", folder.TeamFolderId), zap.String("name", folder.Name))
			pair.Dst = folder
		}
		return createErr
	}
	if err := createIfRequired(); err != nil {
		return err
	}

	// Attach group to the team folder
	attachGroupToTeamFolders := func() error {
		var attachErr error
		ctxFileSrcAsAdmin := z.ctxFileSrc.AsAdminId(ctx.AdminSrc().TeamMemberId)
		ctxFileDstAsAdmin := z.ctxFileDst.AsAdminId(ctx.AdminDst().TeamMemberId)
		svmSrc := sv_sharedfolder_member.NewBySharedFolderId(ctxFileSrcAsAdmin, scope.Pair().Src.TeamFolderId)
		svmDst := sv_sharedfolder_member.NewBySharedFolderId(ctxFileDstAsAdmin, scope.Pair().Dst.TeamFolderId)
		if attachErr = svmSrc.Add(sv_sharedfolder_member.AddByGroup(ctx.GroupSrc(), sv_sharedfolder_member.LevelEditor)); err != nil {
			return attachErr
		}
		if attachErr = svmDst.Add(sv_sharedfolder_member.AddByGroup(ctx.GroupDst(), sv_sharedfolder_member.LevelEditor)); err != nil {
			return attachErr
		}
		return nil
	}
	if err := attachGroupToTeamFolders(); err != nil {
		return err
	}

	// Ensure access
	ensureAccess := func(admin *mo_profile.Profile, ctx api_context.Context, folder *mo_teamfolder.TeamFolder) error {
		mount, ensureErr := sv_sharedfolder.New(ctx.AsMemberId(admin.TeamMemberId)).Resolve(folder.TeamFolderId)
		if ensureErr != nil {
			return ensureErr
		}
		if mount.PathLower == "" {
			return errors.New("the folder is not mounted on the account")
		}
		_, ensureErr = sv_file.NewFiles(ctx.AsMemberId(admin.TeamMemberId)).List(mo_path.NewPath(mount.PathLower))
		if ensureErr != nil {
			return ensureErr
		}
		return nil
	}
	if err := ensureAccess(ctx.AdminSrc(), z.ctxFileSrc, scope.Pair().Src); err != nil {
		z.log().Warn("Could not access to src team folder", zap.String("srcName", scope.Pair().Src.Name))
		return err
	}
	if err := ensureAccess(ctx.AdminDst(), z.ctxFileDst, scope.Pair().Dst); err != nil {
		z.log().Warn("Could not access to src team folder", zap.String("dstName", scope.Pair().Dst.Name))
		return err
	}
	return nil
}

func (z *teamFolderImpl) Content(ctx Context, scope Scope) (err error) {
	ctxSrc := z.ctxFileSrc.
		AsMemberId(ctx.AdminSrc().TeamMemberId).
		WithPath(api_context.Namespace(scope.Pair().Src.TeamFolderId))
	ctxDst := z.ctxFileDst.
		AsMemberId(ctx.AdminDst().TeamMemberId).
		WithPath(api_context.Namespace(scope.Pair().Dst.TeamFolderId))

	ucm := uc_file_mirror.New(ctxSrc, ctxDst)
	return ucm.Mirror(mo_path.NewPath("/"), mo_path.NewPath("/"))
}

func (z *teamFolderImpl) Verify(ctx Context, scope Scope) (err error) {
	ctxSrc := z.ctxFileSrc.
		AsMemberId(ctx.AdminSrc().TeamMemberId).
		WithPath(api_context.Namespace(scope.Pair().Src.TeamFolderId))
	ctxDst := z.ctxFileDst.
		AsMemberId(ctx.AdminDst().TeamMemberId).
		WithPath(api_context.Namespace(scope.Pair().Dst.TeamFolderId))

	ucc := uc_file_compare.New(ctxSrc, ctxDst)
	count, err := ucc.Diff(func(diff mo_file_diff.Diff) error {
		z.log().Warn("Diff", zap.Any("diff", diff))
		return nil
	})

	if count > 0 {
		return errors.New("one or more files differ between source and destination folder")
	}
	return nil
}

func (z *teamFolderImpl) Unmount(ctx Context, scope Scope) (err error) {
	// Detach admin from team folder
	detachGroupFromTeamFolders := func() error {
		var attachErr error
		ctxFileSrcAsAdmin := z.ctxFileSrc.AsAdminId(ctx.AdminSrc().TeamMemberId)
		ctxFileDstAsAdmin := z.ctxFileDst.AsAdminId(ctx.AdminDst().TeamMemberId)
		svmSrc := sv_sharedfolder_member.NewBySharedFolderId(ctxFileSrcAsAdmin, scope.Pair().Src.TeamFolderId)
		svmDst := sv_sharedfolder_member.NewBySharedFolderId(ctxFileDstAsAdmin, scope.Pair().Dst.TeamFolderId)
		if attachErr = svmSrc.Remove(sv_sharedfolder_member.RemoveByGroup(ctx.GroupSrc())); err != nil {
			return attachErr
		}
		if attachErr = svmDst.Remove(sv_sharedfolder_member.RemoveByGroup(ctx.GroupDst())); err != nil {
			return attachErr
		}
		return nil
	}
	if err := detachGroupFromTeamFolders(); err != nil {
		return err
	}

	return nil
}

func (z *teamFolderImpl) Cleanup(ctx Context) (err error) {
	// Remove groups
	errSrc := sv_group.New(z.ctxMgtSrc).Remove(ctx.GroupSrc().GroupId)
	if errSrc != nil {
		z.log().Warn("SRC: Could not remove group", zap.String("groupName", ctx.GroupSrc().GroupName), zap.Error(errSrc))
		err = errSrc
	}

	errDst := sv_group.New(z.ctxMgtDst).Remove(ctx.GroupDst().GroupId)
	if errDst != nil {
		z.log().Warn("SRC: Could not remove group", zap.String("groupName", ctx.GroupDst().GroupName), zap.Error(errDst))
		err = errDst
	}

	return err
}