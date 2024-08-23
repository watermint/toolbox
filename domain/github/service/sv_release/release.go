package sv_release

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_client"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
)

var (
	ErrorNotFound = errors.New("not found")
)

type Release interface {
	Get(tagName string) (release *mo_release.Release, err error)
	List() (releases []*mo_release.Release, err error)
	ListEach(f func(release *mo_release.Release) error) error
	Latest() (release *mo_release.Release, err error)
	CreateDraft(tagName, name, body, branch string) (release *mo_release.Release, err error)
	Publish(releaseId string) (release *mo_release.Release, err error)
}

func New(ctx gh_client.Client, owner, repo string) Release {
	return &releaseImpl{
		ctx:   ctx,
		owner: owner,
		repo:  repo,
	}
}

type releaseImpl struct {
	ctx   gh_client.Client
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

func (z *releaseImpl) ListEach(f func(release *mo_release.Release) error) error {
	l := z.ctx.Log()
	paging := struct {
		PageSize int `url:"per_page"`
		Page     int `url:"page"`
	}{
		PageSize: 100,
		Page:     1,
	}

	for {
		l.Debug("List releases", esl.Int("page", paging.Page))
		endpoint := "repos/" + z.owner + "/" + z.repo + "/releases"
		res := z.ctx.Get(endpoint, api_request.Query(&paging))
		if err, fail := res.Failure(); fail {
			l.Debug("Unable to list releases", esl.Error(err))
			if res.Code() == 404 && paging.Page > 1 {
				l.Debug("No more releases")
				return nil
			}
			return err
		}

		entries, ok := res.Success().Json().Array()
		if !ok {
			l.Debug("Unable to parse response as array")
			return errors.New("invalid response")
		}
		if len(entries) == 0 {
			l.Debug("No more releases")
			return nil
		}
		for _, e := range entries {
			release := &mo_release.Release{}
			if err := e.Model(release); err != nil {
				return err
			}
			if err := f(release); err != nil {
				return err
			}
		}

		paging.Page++
	}
}
