package sv_installation

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_installation"
	"github.com/watermint/toolbox/infra/api/api_parser"
)

var (
	ErrorUnexpectedResponse = errors.New("unexpected response")
)

type Installation interface {
	List() (installations []*mo_installation.Installation, err error)
}

func New(ctx gh_context.Context) Installation {
	return &installationImpl{
		ctx: ctx,
	}
}

type installationImpl struct {
	ctx gh_context.Context
}

func (z *installationImpl) List() (installations []*mo_installation.Installation, err error) {
	endpoint := "user/installations"
	res, err := z.ctx.Get(endpoint).Call()
	if err != nil {
		return nil, err
	}
	j, err := res.Json()
	if err != nil {
		return nil, err
	}
	j1 := j.Get("installations")
	if !j1.IsArray() {
		return nil, ErrorUnexpectedResponse
	}
	installations = make([]*mo_installation.Installation, 0)
	for _, entry := range j.Array() {
		installation := &mo_installation.Installation{}
		if err := api_parser.ParseModel(installation, entry); err != nil {
			return nil, err
		}
		installations = append(installations, installation)
	}
	return installations, nil
}
