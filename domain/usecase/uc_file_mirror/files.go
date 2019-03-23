package uc_file_mirror

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_path"
)

type Files interface {
	Mirror(pathSrc, pathDst mo_path.Path) (err error)
}

func NewFiles(ctxSrc, ctxDst api_context.Context) Files {
	return &filesImpl{
		ctxSrc: ctxSrc,
		ctxDst: ctxDst,
	}
}

type filesImpl struct {
	ctxSrc api_context.Context
	ctxDst api_context.Context
}

func (z *filesImpl) Mirror(pathSrc, pathDst mo_path.Path) (err error) {
	panic("implement me")
}
