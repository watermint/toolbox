package sv_issue

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_issue"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

var (
	ErrorUnexpectedResponse = errors.New("unexpected response")
)

type Issue interface {
	List() (issues []*mo_issue.Issue, err error)
}

func New(ctx gh_context.Context, owner, repo string) Issue {
	return &repoIssueImpl{
		ctx:   ctx,
		owner: owner,
		repo:  repo,
	}
}

type repoIssueImpl struct {
	ctx   gh_context.Context
	owner string
	repo  string
}

func (z *repoIssueImpl) List() (issues []*mo_issue.Issue, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repo + "/issues"
	res := z.ctx.Get(endpoint)
	if err, f := res.Failure(); f {
		return nil, err
	}
	issues = make([]*mo_issue.Issue, 0)
	err = res.Success().Json().ArrayEach(func(e es_json.Json) error {
		issue := &mo_issue.Issue{}
		if err := e.Model(issue); err != nil {
			return err
		}
		issues = append(issues, issue)
		return nil
	})
	return issues, err
}
