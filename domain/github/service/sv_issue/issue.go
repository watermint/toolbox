package sv_issue

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_issue"
	"github.com/watermint/toolbox/infra/api/api_parser"
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
	j, err := res.Json()
	if err != nil {
		return nil, err
	}
	if !j.IsArray() {
		return nil, ErrorUnexpectedResponse
	}
	issues = make([]*mo_issue.Issue, 0)
	for _, entry := range j.Array() {
		issue := &mo_issue.Issue{}
		if err := api_parser.ParseModel(issue, entry); err != nil {
			return nil, err
		}
		issues = append(issues, issue)
	}
	return issues, nil
}
