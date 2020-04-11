package sv_release

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/infra/api/api_parser"
)

var (
	ErrorUnexpectedResponse = errors.New("unexpected response")
)

type Release interface {
	List() (releases []*mo_release.Release, err error)
}

func New(ctx gh_context.Context, owner, repo string) Release {
	return &releaseImpl{
		ctx:   ctx,
		owner: owner,
		repo:  repo,
	}
}

type releaseImpl struct {
	ctx   gh_context.Context
	owner string
	repo  string
}

func (z *releaseImpl) List() (releases []*mo_release.Release, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repo + "/releases"
	res, err := z.ctx.Get(endpoint).Call()
	if err != nil {
		return nil, err
	}
	j, err := res.Json()
	if err != nil {
		return nil, err
	}
	if !j.IsArray() {
		return nil, ErrorUnexpectedResponse
	}
	releases = make([]*mo_release.Release, 0)
	for _, entry := range j.Array() {
		release := &mo_release.Release{}
		if err := api_parser.ParseModel(release, entry); err != nil {
			return nil, err
		}
		releases = append(releases, release)
	}
	return releases, nil
}
