package sv_reference

import (
	"github.com/watermint/toolbox/domain/github/api/gh_client"
	"github.com/watermint/toolbox/domain/github/model/mo_reference"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Reference interface {
	Create(ref, sha string) (created *mo_reference.Reference, err error)
}

func New(ctx gh_client.Client, owner, repository string) Reference {
	return &referenceImpl{
		ctx:        ctx,
		owner:      owner,
		repository: repository,
	}
}

type referenceImpl struct {
	ctx        gh_client.Client
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
	res := z.ctx.Post(endpoint, api_request.Param(&p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	created = &mo_reference.Reference{}
	err = res.Success().Json().Model(created)
	return
}
