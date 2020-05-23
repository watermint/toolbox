package sv_content

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_commit"
	"github.com/watermint/toolbox/domain/github/model/mo_content"
	"github.com/watermint/toolbox/infra/api/api_request"
	"strings"
)

type Content interface {
	Get(path string, opts ...ContentOpt) (c mo_content.Contents, err error)
	CreateOrUpdate(message, content string, pts ...ContentOpt) (cts mo_content.Content, commit mo_commit.Commit, err error)
}

type ContentOpt func(o contentOpts) contentOpts
type contentOpts struct {
	sha    string
	branch string
	ref    string
}

func (z contentOpts) Apply(opts ...ContentOpt) contentOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		x, y := opts[0], opts[1:]
		return x(z).Apply(y...)
	}
}

func Ref(ref string) ContentOpt {
	return func(o contentOpts) contentOpts {
		o.ref = ref
		return o
	}
}

func Sha(sha string) ContentOpt {
	return func(o contentOpts) contentOpts {
		o.sha = sha
		return o
	}
}
func Branch(branch string) ContentOpt {
	return func(o contentOpts) contentOpts {
		o.branch = branch
		return o
	}
}

type ParamRef struct {
	Ref string `json:"ref" url:"ref"`
}

func New(ctx gh_context.Context, owner, repo string) Content {
	return &ctsImpl{
		ctx:   ctx,
		owner: owner,
		repo:  repo,
	}
}

type ctsImpl struct {
	ctx   gh_context.Context
	owner string
	repo  string
}

func (z ctsImpl) makePath(path string) string {
	base := strings.Join([]string{
		"repos",
		z.owner,
		z.repo,
		"contents",
		path,
	}, "/")
	return strings.ReplaceAll(base, "//", "/")
}

func (z ctsImpl) Get(path string, opts ...ContentOpt) (c mo_content.Contents, err error) {
	co := contentOpts{}.Apply(opts...)
	endpoint := z.makePath(path)

	d := make([]api_request.RequestDatum, 0)
	if co.ref != "" {
		d = append(d, api_request.Param(&ParamRef{Ref: co.ref}))
	}

	res := z.ctx.Get(endpoint, d...)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	if c, err := mo_content.NewContents(res.Success().Json()); err != nil {
		return nil, err
	} else {
		return c, nil
	}
}

func (z ctsImpl) CreateOrUpdate(message, content string, pts ...ContentOpt) (cts mo_content.Content, commit mo_commit.Commit, err error) {
	panic("implement me")
}
