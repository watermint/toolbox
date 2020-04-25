package sv_issue

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_issue"
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
	res, err := z.ctx.Get(endpoint).Call()
	if err != nil {
		return nil, err
	}
	entries, found := res.Success().Json().Array()
	if !found {
		return nil, ErrorUnexpectedResponse
	}
	issues = make([]*mo_issue.Issue, 0)
	for _, entry := range entries {
		issue := &mo_issue.Issue{}
		if _, err := entry.Model(issue); err != nil {
			return nil, err
		}
		issues = append(issues, issue)
	}
	return issues, nil
}
