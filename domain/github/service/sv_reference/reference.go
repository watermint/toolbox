package sv_reference

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_reference"
)

type Reference interface {
	Create(ref, sha string) (created *mo_reference.Reference, err error)
}

func New(ctx gh_context.Context, owner, repository string) Reference {
	return &referenceImpl{
		ctx:        ctx,
		owner:      owner,
		repository: repository,
	}
}

type referenceImpl struct {
	ctx        gh_context.Context
	owner      string
	repository string
}

func (z *referenceImpl) Create(ref, sha string) (created *mo_reference.Reference, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repository + "/git/refs"

	p := struct {
		Ref string `json:"ref"`
		Sha string `json:"sha"`
	}{
		Ref: ref,
		Sha: sha,
	}
	res, err := z.ctx.Post(endpoint).Param(&p).Call()
	if err != nil {
		return nil, err
	}
	created = &mo_reference.Reference{}
	if _, err := res.Body().Json().Model(created); err != nil {
		return nil, err
	}
	return created, nil
}
