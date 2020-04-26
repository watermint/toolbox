package sv_tag

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_tag"
	"github.com/watermint/toolbox/domain/github/service/sv_reference"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/infra/api/api_request"
)

var (
	ErrorUnexpectedResponse = errors.New("unexpected response")
)

type Tag interface {
	List() (tags []*mo_tag.Tag, err error)
	Create(tagName, message, sha string) (tag *mo_tag.Tag, err error)
}

func New(ctx gh_context.Context, owner, repo string) Tag {
	return &tagImpl{
		ctx:   ctx,
		owner: owner,
		repo:  repo,
	}
}

type tagImpl struct {
	ctx   gh_context.Context
	owner string
	repo  string
}

func (z *tagImpl) List() (tags []*mo_tag.Tag, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repo + "/tags"
	res := z.ctx.Get(endpoint)
	if err, fail := res.Failure(); fail {
		return nil, err
	}

	tags = make([]*mo_tag.Tag, 0)
	err = res.Success().Json().ArrayEach(func(e tjson.Json) error {
		tag := &mo_tag.Tag{}
		if err := e.Model(tag); err != nil {
			return err
		}
		tags = append(tags, tag)
		return nil
	})
	return
}

func (z *tagImpl) Create(tagName, message, sha string) (tag *mo_tag.Tag, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repo + "/git/tags"

	p := struct {
		Tag     string `json:"tag"`
		Message string `json:"message"`
		Object  string `json:"object"`
		Type    string `json:"type"`
	}{
		Tag:     tagName,
		Message: message,
		Object:  sha,
		Type:    "commit",
	}
	res := z.ctx.Post(endpoint, api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	tag = &mo_tag.Tag{}
	err = res.Success().Json().Model(tag)
	if err != nil {
		return nil, err
	}

	_, err = sv_reference.New(z.ctx, z.owner, z.repo).Create("refs/tags/"+tagName, tag.Sha)
	if err != nil {
		return nil, err
	}

	return tag, nil
}
