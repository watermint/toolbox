package uc_member_mirror

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/usecase/uc_file_mirror"
	"go.uber.org/zap"
)

type Mirror interface {
	Mirror(srcEmail, dstEmail string) error
}

func New(ctxFileSrc, ctxFileDst api_context.Context) Mirror {
	return &mirrorImpl{
		ctxFileSrc: ctxFileSrc,
		ctxFileDst: ctxFileDst,
	}
}

type mirrorImpl struct {
	ctxFileSrc api_context.Context
	ctxFileDst api_context.Context
}

func (z *mirrorImpl) log() *zap.Logger {
	return z.ctxFileSrc.Log()
}

func (z *mirrorImpl) Mirror(srcEmail, dstEmail string) error {
	l := z.log().With(zap.String("srcEmail", srcEmail), zap.String("dstEmail", dstEmail))
	l.Debug("Start mirroring process")

	l.Debug("Lookup member profiles")
	srcProfile, err := sv_member.New(z.ctxFileSrc).ResolveByEmail(srcEmail)
	if err != nil {
		l.Error("Unable to lookup member", zap.String("lookupEmail", srcEmail), zap.Error(err))
		return err
	}
	dstProfile, err := sv_member.New(z.ctxFileDst).ResolveByEmail(dstEmail)
	if err != nil {
		l.Error("Unable to lookup member", zap.String("lookupEmail", dstEmail), zap.Error(err))
		return err
	}

	ctxFileSrcAsMember := z.ctxFileSrc.AsMemberId(srcProfile.TeamMemberId).WithPath(api_context.Namespace(srcProfile.MemberFolderId))
	ctxFileDstAsMember := z.ctxFileDst.AsMemberId(dstProfile.TeamMemberId).WithPath(api_context.Namespace(dstProfile.MemberFolderId))

	l.Info("Start mirroring files")
	ucm := uc_file_mirror.New(ctxFileSrcAsMember, ctxFileDstAsMember)
	return ucm.Mirror(mo_path.NewPath("/"), mo_path.NewPath("/"))
}
