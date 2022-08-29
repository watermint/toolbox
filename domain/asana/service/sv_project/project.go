package sv_project

import (
	"github.com/watermint/toolbox/domain/asana/api/as_client"
	"github.com/watermint/toolbox/domain/asana/api/as_pagination"
	"github.com/watermint/toolbox/domain/asana/model/mo_project"
	"github.com/watermint/toolbox/domain/asana/model/mo_team"
	"github.com/watermint/toolbox/domain/asana/model/mo_workspace"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Project interface {
	Resolve(gid string) (prj *mo_project.Project, err error)
	List(opts ...Opt) (prj []*mo_project.Project, err error)
}

type Opts struct {
	Workspace string `url:"workspace,omitempty"`
	Team      string `url:"team,omitempty"`
}

type Opt func(o Opts) Opts

func (z Opts) Apply(opts ...Opt) Opts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:]...)
	}
}

func Workspace(ws *mo_workspace.Workspace) Opt {
	return func(o Opts) Opts {
		o.Workspace = ws.Gid
		return o
	}
}

func Team(team *mo_team.Team) Opt {
	return func(o Opts) Opts {
		o.Team = team.Gid
		return o
	}
}

func New(ctx as_client.Client) Project {
	return &prjImpl{
		ctx: ctx,
	}
}

type prjImpl struct {
	ctx as_client.Client
}

func (z prjImpl) Resolve(gid string) (prj *mo_project.Project, err error) {
	prj = &mo_project.Project{}
	res := z.ctx.Get("projects/" + gid)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	err = res.Success().Json().FindModel("data", prj)
	return
}

func (z prjImpl) List(opts ...Opt) (prjs []*mo_project.Project, err error) {
	o := Opts{}.Apply(opts...)
	pg := as_pagination.New(z.ctx).WithEndpoint("projects").WithData(api_request.Query(o))
	prjs = make([]*mo_project.Project, 0)
	err = pg.OnData(func(entry es_json.Json) error {
		prj := &mo_project.Project{}
		if errM := entry.Model(prj); errM != nil {
			return errM
		} else {
			prjs = append(prjs, prj)
			return nil
		}
	})
	return
}
