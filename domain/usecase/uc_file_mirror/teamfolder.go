package uc_file_mirror

import "github.com/watermint/toolbox/domain/infra/api_context"

type TeamFolder interface {
	Mirror(teamFolderName string) (err error)
}

func NewTeamFolder(ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst api_context.Context) TeamFolder {
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

func (*teamFolderImpl) Mirror(teamFolderName string) (err error) {
	panic("implement me")
}
