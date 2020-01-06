package teamfolder

import (
	"encoding/json"
	"errors"
	"fmt"
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
	"github.com/watermint/toolbox/domain/usecase/uc_compare_paths"
	"github.com/watermint/toolbox/domain/usecase/uc_file_mirror"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"go.uber.org/zap"
	"strings"
	"time"
)

const (
	MirrorGroupNamePrefix = "toolbox-teamfolder-mirror"
)

type Replicator interface {
	// All team folder scope
	AllFolderScope() (ctx Context, err error)

	// Specific team folders.
	PartialScope(names []string) (ctx Context, err error)

	// Mirror
	Mirror(ctx Context, opts ...MirrorOpt) (err error)

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

	// Archive
	Archive(ctx Context, scope Scope) (err error)

	// Clean up permissions which used for mirroring
	Cleanup(ctx Context) (err error)

	// Verify scope
	VerifyScope(ctx Context) (err error)
}

type MirrorOpt func(opt *mirrorOpts) *mirrorOpts
type mirrorOpts struct {
	archiveOnSuccess bool
	skipVerify       bool
}

func ArchiveOnSuccess() MirrorOpt {
	return func(opt *mirrorOpts) *mirrorOpts {
		opt.archiveOnSuccess = true
		return opt
	}
}
func SkipVerify() MirrorOpt {
	return func(opt *mirrorOpts) *mirrorOpts {
		opt.skipVerify = true
		return opt
	}
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

func MarshalContext(c Context) (b []byte, err error) {
	b, err = json.Marshal(c)
	return
}
func UnmarshalContext(b []byte) (c Context, err error) {
	mc := &mirrorContext{}
	err = json.Unmarshal(b, mc)
	if err != nil {
		return nil, err
	}
	return mc, nil
}

type mirrorContext struct {
	MirrorPairs    []*MirrorPair       `json:"pairs"`
	MirrorGroupSrc *mo_group.Group     `json:"group_src"`
	MirrorGroupDst *mo_group.Group     `json:"group_dst"`
	MirrorAdminSrc *mo_profile.Profile `json:"admin_src"`
	MirrorAdminDst *mo_profile.Profile `json:"admin_dst"`
}

func (z *mirrorContext) SetGroups(src, dst *mo_group.Group) {
	z.MirrorGroupSrc = src
	z.MirrorGroupDst = dst
}

func (z *mirrorContext) SetAdmins(src, dst *mo_profile.Profile) {
	z.MirrorAdminSrc = src
	z.MirrorAdminDst = dst
}

func (z *mirrorContext) Pairs() (pairs []*MirrorPair) {
	return z.MirrorPairs
}

func (z *mirrorContext) GroupSrc() *mo_group.Group {
	return z.MirrorGroupSrc
}

func (z *mirrorContext) GroupDst() *mo_group.Group {
	return z.MirrorGroupDst
}

func (z *mirrorContext) AdminSrc() *mo_profile.Profile {
	return z.MirrorAdminSrc
}

func (z *mirrorContext) AdminDst() *mo_profile.Profile {
	return z.MirrorAdminDst
}

type Replication struct {
	TargetNames  []string
	TargetAll    bool
	Verification rp_model.RowReport
	SrcFile      rc_conn.ConnBusinessFile
	SrcMgmt      rc_conn.ConnBusinessMgmt
	DstFile      rc_conn.ConnBusinessFile
	DstMgmt      rc_conn.ConnBusinessMgmt
}

func (z *Replication) Exec(c app_control.Control) (err error) {
	var ctx Context
	if z.TargetAll {
		ctx, err = z.AllFolderScope()
		if err != nil {
			return err
		}
	} else {
		ctx, err = z.PartialScope(z.TargetNames)
		if err != nil {
			return err
		}
	}
	return z.Mirror(c, ctx)
}

func (z *Replication) Test(c app_control.Control) error {
	return qt_endtoend.HumanInteractionRequired()
}

func (z *Replication) Preset() {
	z.Verification.SetModel(&mo_file_diff.Diff{})
	z.SrcFile.SetPeerName("src")
	z.SrcMgmt.SetPeerName("src")
	z.DstFile.SetPeerName("dst")
	z.DstMgmt.SetPeerName("dst")
}

func (z *Replication) AllFolderScope() (ctx Context, err error) {
	mc := &mirrorContext{
		MirrorPairs: make([]*MirrorPair, 0),
	}
	ctx = mc
	svt := sv_teamfolder.New(z.SrcFile.Context())
	folders, err := svt.List()
	if err != nil {
		return nil, err
	}
	for _, folder := range folders {
		mc.MirrorPairs = append(mc.MirrorPairs, &MirrorPair{
			Src: folder,
			Dst: nil,
		})
	}
	return
}

func (z *Replication) PartialScope(names []string) (ctx Context, err error) {
	mc := &mirrorContext{
		MirrorPairs: make([]*MirrorPair, 0),
	}
	ctx = mc
	svt := sv_teamfolder.New(z.SrcFile.Context())
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
			mc.MirrorPairs = append(mc.MirrorPairs, &MirrorPair{
				Src: folder,
				Dst: nil,
			})
		}
	}
	return
}

// Verify scope
func (z *Replication) VerifyScope(c app_control.Control, ctx Context) (err error) {
	if err = z.Inspect(c, ctx); err != nil {
		return err
	}
	var lastErr error
	lastErr = nil
	if err = z.Bridge(c, ctx); err != nil {
		lastErr = err
	}
	for _, pair := range ctx.Pairs() {
		scope := NewScope(pair)

		if err = z.Mount(c, ctx, scope); err != nil {
			lastErr = err
			continue
		}
		if err = z.Verify(c, ctx, scope); err != nil {
			lastErr = err
		}
		if err = z.Unmount(c, ctx, scope); err != nil {
			lastErr = err
		}
	}
	if err = z.Cleanup(c, ctx); err != nil {
		lastErr = err
	}
	return lastErr
}

func (z *Replication) Mirror(c app_control.Control, ctx Context, opts ...MirrorOpt) (err error) {
	l := c.Log()
	mo := &mirrorOpts{}
	for _, o := range opts {
		o(mo)
	}

	if err = z.Inspect(c, ctx); err != nil {
		return err
	}
	var lastErr error
	lastErr = nil
	if err = z.Bridge(c, ctx); err != nil {
		lastErr = err
	}
	for _, pair := range ctx.Pairs() {
		scope := NewScope(pair)

		if err = z.Mount(c, ctx, scope); err != nil {
			lastErr = err
			continue
		}
		archive := false
		if err = z.Content(c, ctx, scope); err != nil {
			lastErr = err
		} else {
			if mo.skipVerify {
				l.Info("Skip verification step")
			} else {
				if err = z.Verify(c, ctx, scope); err != nil {
					lastErr = err
				} else if mo.archiveOnSuccess {
					archive = true
				}
			}
		}
		if err = z.Unmount(c, ctx, scope); err != nil {
			lastErr = err
		}
		if archive {
			if err = z.Archive(c, ctx, scope); err != nil {
				lastErr = err
			}
		}
	}
	if err = z.Cleanup(c, ctx); err != nil {
		lastErr = err
	}
	return lastErr
}

func (z *Replication) Inspect(c app_control.Control, ctx Context) (err error) {
	l := c.Log()
	// Identify admins
	identifyAdmins := func() error {
		adminSrc, err := sv_profile.NewTeam(z.SrcMgmt.Context()).Admin()
		if err != nil {
			return err
		}
		adminDst, err := sv_profile.NewTeam(z.DstMgmt.Context()).Admin()
		if err != nil {
			return err
		}
		l.Debug("Admins identified",
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
		infoSrc, err := sv_team.New(z.SrcMgmt.Context()).Info()
		if err != nil {
			return err
		}
		infoDst, err := sv_team.New(z.DstMgmt.Context()).Info()
		if err != nil {
			return err
		}
		l.Debug("Team info",
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
			l.Warn("Source and destination team are the same team.")
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
		inspectErr = nil
		for _, pair := range ctx.Pairs() {
			l.Info("SRC: Team folder status",
				zap.String("id", pair.Src.TeamFolderId),
				zap.String("name", pair.Src.Name),
				zap.String("status", pair.Src.Status),
			)
			if pair.Src.Status != "active" {
				l.Warn("SRC: Non active folder found",
					zap.String("srcId", pair.Src.TeamFolderId),
					zap.String("srcName", pair.Src.Name),
					zap.String("srcStatus", pair.Src.Status),
				)
				//inspectErr = errors.New("one or more team folders are not active")
			}
		}
		return inspectErr
	}
	if err := inspectSrcFolders(); err != nil {
		return err
	}

	// retrieve destination folders
	svt := sv_teamfolder.New(z.DstFile.Context())
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
		inspectErr = nil
		for _, pair := range ctx.Pairs() {
			if folder := pair.Dst; folder != nil {
				l.Info("DST: Team folder status",
					zap.String("srcId", pair.Src.TeamFolderId),
					zap.String("srcName", pair.Src.Name),
					zap.String("srcStatus", pair.Src.Status),
					zap.String("dstId", folder.TeamFolderId),
					zap.String("dstName", folder.Name),
					zap.String("dstStatus", folder.Status),
				)
				if pair.Dst.Status != "active" {
					l.Info("DST: Non active folder found",
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

func (z *Replication) Bridge(c app_control.Control, ctx Context) (err error) {
	l := c.Log()
	groupName := fmt.Sprintf("%s-%x", MirrorGroupNamePrefix, time.Now().Unix())
	l.Info("Bridge", zap.String("groupName", groupName))

	// Create groups
	groupSrc, err := sv_group.New(z.SrcMgmt.Context()).Create(groupName, sv_group.CompanyManaged())
	if err != nil {
		return err
	}
	ctx.SetGroups(groupSrc, nil)

	groupDst, err := sv_group.New(z.DstMgmt.Context()).Create(groupName, sv_group.CompanyManaged())
	if err != nil {
		return err
	}
	ctx.SetGroups(groupSrc, groupDst)
	l.Debug("Groups created", zap.String("srcGroupId", groupSrc.GroupId), zap.String("dstGroupId", groupDst.GroupId), zap.String("groupName", groupName))

	// Add admins to groups
	_, err = sv_group_member.New(z.SrcMgmt.Context(), groupSrc).Add(sv_group_member.ByTeamMemberId(ctx.AdminSrc().TeamMemberId))
	if err != nil {
		return err
	}
	_, err = sv_group_member.New(z.DstMgmt.Context(), groupDst).Add(sv_group_member.ByTeamMemberId(ctx.AdminDst().TeamMemberId))
	if err != nil {
		return err
	}
	l.Debug("Admins added to groups", zap.String("srcGroupId", groupSrc.GroupId), zap.String("dstGroupId", groupDst.GroupId), zap.String("groupName", groupName))

	return nil
}

func (z *Replication) Mount(c app_control.Control, ctx Context, scope Scope) (err error) {
	l := c.Log().With(zap.Any("pair", scope.Pair()))
	l.Info("Mount")

	// Create team folder if required
	createIfRequired := func() error {
		svt := sv_teamfolder.New(z.DstFile.Context())
		pair := scope.Pair()
		if pair.Dst == nil {
			folder, err := svt.Create(pair.Src.Name, sv_teamfolder.SyncNoSync())
			if err != nil {
				if strings.HasPrefix(api_util.ErrorSummary(err), "folder_name_already_used") {
					l.Debug("Skip: Already created")
					return nil
				}
				l.Warn("DST: Unable to create team folder", zap.String("name", pair.Src.Name), zap.Error(err))
				return errors.New("could not create one or more team folders in the destination team")
			}
			l.Debug("DST: Team folder created", zap.String("id", folder.TeamFolderId), zap.String("name", folder.Name))
			pair.Dst = folder
		}
		return nil
	}
	if err := createIfRequired(); err != nil {
		return err
	}

	// Attach group to the team folder
	attachGroupToTeamFolders := func() error {
		var attachErr error
		attachErr = nil
		srcFileAsAdmin := z.SrcFile.Context().AsAdminId(ctx.AdminSrc().TeamMemberId)
		dstFileAsAdmin := z.DstFile.Context().AsAdminId(ctx.AdminDst().TeamMemberId)
		svmSrc := sv_sharedfolder_member.NewBySharedFolderId(srcFileAsAdmin, scope.Pair().Src.TeamFolderId)
		svmDst := sv_sharedfolder_member.NewBySharedFolderId(dstFileAsAdmin, scope.Pair().Dst.TeamFolderId)
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
		_, ensureErr = sv_file.NewFiles(ctx.AsMemberId(admin.TeamMemberId)).List(mo_path.NewDropboxPath(mount.PathLower))
		if ensureErr != nil {
			return ensureErr
		}
		return nil
	}
	if err := ensureAccess(ctx.AdminSrc(), z.SrcFile.Context(), scope.Pair().Src); err != nil {
		l.Warn("Could not access to src team folder", zap.String("srcName", scope.Pair().Src.Name))
		return err
	}
	if err := ensureAccess(ctx.AdminDst(), z.DstFile.Context(), scope.Pair().Dst); err != nil {
		l.Warn("Could not access to src team folder", zap.String("dstName", scope.Pair().Dst.Name))
		return err
	}
	return nil
}

func (z *Replication) Content(c app_control.Control, ctx Context, scope Scope) (err error) {
	l := c.Log().With(
		zap.String("folderSrcId", scope.Pair().Src.TeamFolderId),
		zap.String("folderSrcName", scope.Pair().Src.Name),
		zap.String("folderDstId", scope.Pair().Dst.TeamFolderId),
		zap.String("folderDstName", scope.Pair().Dst.Name),
	)
	l.Info("Mirroring content")

	ctxSrc := z.SrcFile.Context().
		AsMemberId(ctx.AdminSrc().TeamMemberId).
		WithPath(api_context.Namespace(scope.Pair().Src.TeamFolderId))
	ctxDst := z.DstFile.Context().
		AsMemberId(ctx.AdminDst().TeamMemberId).
		WithPath(api_context.Namespace(scope.Pair().Dst.TeamFolderId))

	ucm := uc_file_mirror.New(ctxSrc, ctxDst)
	return ucm.Mirror(mo_path.NewDropboxPath("/"), mo_path.NewDropboxPath("/"))
}

func (z *Replication) Verify(c app_control.Control, ctx Context, scope Scope) (err error) {
	l := c.Log().With(
		zap.String("folderSrcId", scope.Pair().Src.TeamFolderId),
		zap.String("folderSrcName", scope.Pair().Src.Name),
		zap.String("folderDstId", scope.Pair().Dst.TeamFolderId),
		zap.String("folderDstName", scope.Pair().Dst.Name),
	)
	if err := z.Verification.Open(); err != nil {
		c.UI().ErrorK("usecase.uc_teamfolder_mirror.err.unable_to_create_diff_report", app_msg.P{
			"ErrorK": err.Error(),
		})
		return err
	}

	l.Info("Verify: comparing source and destination")

	ctxSrc := z.SrcFile.Context().
		AsMemberId(ctx.AdminSrc().TeamMemberId).
		WithPath(api_context.Namespace(scope.Pair().Src.TeamFolderId))
	ctxDst := z.DstFile.Context().
		AsMemberId(ctx.AdminDst().TeamMemberId).
		WithPath(api_context.Namespace(scope.Pair().Dst.TeamFolderId))

	ucc := uc_compare_paths.New(ctxSrc, ctxDst, c.UI())
	count, err := ucc.Diff(
		mo_path.NewDropboxPath(""), mo_path.NewDropboxPath(""),
		func(diff mo_file_diff.Diff) error {
			l.Warn("Diff", zap.Any("diff", diff))
			z.Verification.Row(&diff)
			return nil
		})

	if count > 0 {
		return errors.New("one or more files differ between source and destination folder")
	}
	return nil
}

func (z *Replication) Unmount(c app_control.Control, ctx Context, scope Scope) (err error) {
	l := c.Log().With(
		zap.String("folderSrcId", scope.Pair().Src.TeamFolderId),
		zap.String("folderSrcName", scope.Pair().Src.Name),
		zap.String("folderDstId", scope.Pair().Dst.TeamFolderId),
		zap.String("folderDstName", scope.Pair().Dst.Name),
	)
	l.Info("Unmount: detach admin from team folder(s)")

	// Detach admin from team folder
	detachGroupFromTeamFolders := func() error {
		var attachErr error
		attachErr = nil
		srcFileAsAdmin := z.SrcFile.Context().AsAdminId(ctx.AdminSrc().TeamMemberId)
		dstFileAsAdmin := z.DstFile.Context().AsAdminId(ctx.AdminDst().TeamMemberId)
		svmSrc := sv_sharedfolder_member.NewBySharedFolderId(srcFileAsAdmin, scope.Pair().Src.TeamFolderId)
		svmDst := sv_sharedfolder_member.NewBySharedFolderId(dstFileAsAdmin, scope.Pair().Dst.TeamFolderId)
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

func (z *Replication) Archive(c app_control.Control, ctx Context, scope Scope) (err error) {
	l := c.Log()
	l.Info("Archive: Archiving team folder", zap.String("name", scope.Pair().Src.Name))
	svt := sv_teamfolder.New(z.SrcFile.Context())
	if _, err := svt.Archive(scope.Pair().Src); err != nil {
		return err
	}

	return nil
}

func (z *Replication) Cleanup(c app_control.Control, ctx Context) (err error) {
	l := c.Log()
	l.Info("Cleanup")
	err = nil

	// Remove groups
	l.Info("Cleanup: Remove temporary group (source)", zap.String("name", ctx.GroupSrc().GroupName))
	errSrc := sv_group.New(z.SrcMgmt.Context()).Remove(ctx.GroupSrc().GroupId)
	if errSrc != nil {
		l.Warn("SRC: Could not remove group", zap.String("groupName", ctx.GroupSrc().GroupName), zap.Error(errSrc))
		err = errSrc
	}

	l.Info("Cleanup: Remove temporary group (dest)", zap.String("name", ctx.GroupDst().GroupName))
	errDst := sv_group.New(z.DstMgmt.Context()).Remove(ctx.GroupDst().GroupId)
	if errDst != nil {
		l.Warn("SRC: Could not remove group", zap.String("groupName", ctx.GroupDst().GroupName), zap.Error(errDst))
		err = errDst
	}

	return err
}
