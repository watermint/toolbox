package uc_file_merge

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
)

type Merge interface {
	Merge(from, to mo_path.Path) error
}

func New(ctx api_context.Context, k app_kitchen.Kitchen) Merge {
	return &mergeImpl{
		ctx: ctx,
		k:   k,
	}
}

type mergeImpl struct {
	ctx api_context.Context
	k   app_kitchen.Kitchen
}

func (z *mergeImpl) Merge(from, to mo_path.Path) error {
	panic("implement me")
}
