package sv_content

import (
	"encoding/base64"
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_commit"
	"github.com/watermint/toolbox/domain/github/model/mo_content"
	"github.com/watermint/toolbox/infra/api/api_request"
	"strings"
)

var (
	ErrorUnexpectedFormat = errors.New("unexpected format")
)

type Content interface {
	Get(path string, opts ...ContentOpt) (c mo_content.Contents, err error)
	Put(path, message, content string, opts ...ContentOpt) (cts mo_content.Content, commit mo_commit.Commit, err error)
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
		d = append(d, api_request.Query(&ParamRef{Ref: co.ref}))
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

func (z ctsImpl) Put(path, message, content string, opts ...ContentOpt) (cts mo_content.Content, commit mo_commit.Commit, err error) {
	co := contentOpts{}.Apply(opts...)
	endpoint := z.makePath(path)

	req := struct {
		Message string `json:"message"`
		Content string `json:"content"`
		Sha     string `json:"sha,omitempty"`
		Branch  string `json:"branch,omitempty"`
	}{
		Message: message,
		Content: base64.StdEncoding.EncodeToString([]byte(content)),
		Sha:     co.sha,
		Branch:  co.branch,
	}

	res := z.ctx.Put(endpoint, api_request.Param(&req))
	if err, fail := res.Failure(); fail {
		return cts, commit, err
	}

	j := res.Success().Json()
	if err := j.FindModel("content", &cts); err != nil {
		return cts, commit, err
	}
	if err := j.FindModel("commit", &commit); err != nil {
		return cts, commit, err
	}
	return cts, commit, nil
}
