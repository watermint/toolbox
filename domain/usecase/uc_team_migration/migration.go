package uc_team_migration

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"go.uber.org/zap"
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

	// Verify
	Verify(ctx Context) (err error)
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

func New(ctxExec *app.ExecContext, ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst api_context.Context, report app_report.Report) Migration {
	return &migrationImpl{
		ctxExec:          ctxExec,
		ctxFileSrc:       ctxFileSrc,
		ctxMgtSrc:        ctxMgtSrc,
		ctxFileDst:       ctxFileDst,
		ctxMgtDst:        ctxMgtDst,
		teamFolderMirror: uc_teamfolder_mirror.New(ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst, report),
		report:           report,
	}
}

type migrationImpl struct {
	ctxExec          *app.ExecContext
	ctxFileSrc       api_context.Context
	ctxFileDst       api_context.Context
	ctxMgtSrc        api_context.Context
	ctxMgtDst        api_context.Context
	teamFolderMirror uc_teamfolder_mirror.TeamFolder
	report           app_report.Report
}

func (z *migrationImpl) log() *zap.Logger {
	return z.ctxExec.Log()
}

func (z *migrationImpl) isTeamOwnedSharedFolder(ctx Context, namespaceId string) (user *mo_sharedfolder_member.User, exist bool) {
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
