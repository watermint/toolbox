package sv_issue

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_issue"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

var (
	ErrorUnexpectedResponse = errors.New("unexpected response")
)

type Issue interface {
	List(opts ...ListOpt) (issues []*mo_issue.Issue, err error)
}

type listOpts struct {
	State  string `url:"state,omitempty"`
	Filter string `url:"filter,omitempty"`
	Labels string `url:"labels,omitempty"`
	Since  string `url:"since,omitempty"`
}

type ListOpt func(o listOpts) listOpts

func ListState(state string) ListOpt {
	return func(o listOpts) listOpts {
		o.State = state
		return o
	}
}
func ListFilter(filter string) ListOpt {
	return func(o listOpts) listOpts {
		o.Filter = filter
		return o
	}
}
func ListLabels(labels string) ListOpt {
	return func(o listOpts) listOpts {
		o.Labels = labels
		return o
	}
}
func ListSince(since string) ListOpt {
	return func(o listOpts) listOpts {
		o.Since = since
		return o
	}
}

func (z listOpts) Apply(opts []ListOpt) listOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
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

func (z *repoIssueImpl) List(opts ...ListOpt) (issues []*mo_issue.Issue, err error) {
	lo := listOpts{}.Apply(opts)
	endpoint := "repos/" + z.owner + "/" + z.repo + "/issues"
	res := z.ctx.Get(endpoint, api_request.Query(&lo))
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
