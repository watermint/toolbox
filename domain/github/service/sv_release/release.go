package sv_release

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/infra/api/api_parser"
)

var (
	ErrorNotFound           = errors.New("not found")
	ErrorUnexpectedResponse = errors.New("unexpected response")
)

type Release interface {
	Get(tagName string) (release *mo_release.Release, err error)
	List() (releases []*mo_release.Release, err error)
	CreateDraft(tagName, name, body, branch string) (release *mo_release.Release, err error)
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

func (z *releaseImpl) Get(tagName string) (release *mo_release.Release, err error) {
	//endpoint := "repos/" + z.owner + "/" + z.repo + "/releases/tags/" + tagName
	//res, err := z.ctx.Get(endpoint).Call()
	//if err != nil {
	//	return nil, err
	//}
	//release = &mo_release.Release{}
	//if err := res.Model(release); err != nil {
	//	return nil, err
	//}
	//return release, nil

	rels, err := z.List()
	if err != nil {
		return nil, err
	}
	for _, rel := range rels {
		if rel.TagName == tagName {
			return rel, nil
		}
	}
	return nil, ErrorNotFound
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
	res, err := z.ctx.Post(endpoint).Param(&p).Call()
	if err != nil {
		return nil, err
	}
	j, err := res.Json()
	if err != nil {
		return nil, err
	}
	release = &mo_release.Release{}
	if err := api_parser.ParseModel(release, j); err != nil {
		return nil, err
	}
	return release, nil
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
