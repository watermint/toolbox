package sv_release

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

var (
	ErrorNotFound = errors.New("not found")
)

type Release interface {
	Get(tagName string) (release *mo_release.Release, err error)
	List() (releases []*mo_release.Release, err error)
	Latest() (release *mo_release.Release, err error)
	CreateDraft(tagName, name, body, branch string) (release *mo_release.Release, err error)
	Publish(releaseId string) (release *mo_release.Release, err error)
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

func (z *releaseImpl) Latest() (release *mo_release.Release, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repo + "/releases/latest"
	res := z.ctx.Get(endpoint)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	release = &mo_release.Release{}
	err = res.Success().Json().Model(release)
	return
}

func (z *releaseImpl) Get(tagName string) (release *mo_release.Release, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repo + "/releases/tags/" + tagName
	res := z.ctx.Get(endpoint)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	release = &mo_release.Release{}
	err = res.Success().Json().Model(release)
	return
}

func (z *releaseImpl) CreateDraft(tagName, name, body, branch string) (release *mo_release.Release, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repo + "/releases"
	p := struct {
		TagName         string `json:"tag_name"`
		TargetCommitish string `json:"target_commitish"`
		Name            string `json:"name"`
		Body            string `json:"body"`
		Draft           bool   `json:"draft"`
	}{
		TagName:         tagName,
		TargetCommitish: branch,
		Name:            name,
		Body:            body,
		Draft:           true,
	}
	res := z.ctx.Post(endpoint, api_request.Param(&p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	release = &mo_release.Release{}
	err = res.Success().Json().Model(release)
	return
}

func (z *releaseImpl) Publish(releaseId string) (release *mo_release.Release, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repo + "/releases/" + releaseId
	p := struct {
		Draft bool `json:"draft"`
	}{
		Draft: false,
	}
	res := z.ctx.Patch(endpoint, api_request.Param(&p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	release = &mo_release.Release{}
	err = res.Success().Json().Model(release)
	return
}

func (z *releaseImpl) List() (releases []*mo_release.Release, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repo + "/releases"
	res := z.ctx.Get(endpoint)
	if err, fail := res.Failure(); fail {
		return nil, err
	}

	releases = make([]*mo_release.Release, 0)
	err = res.Success().Json().ArrayEach(func(e es_json.Json) error {
		release := &mo_release.Release{}
		if err := e.Model(release); err != nil {
			return err
		}
		releases = append(releases, release)
		return nil
	})
	return
}
