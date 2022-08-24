package sv_team

import (
	"github.com/watermint/toolbox/domain/asana/api/as_client"
	"github.com/watermint/toolbox/domain/asana/api/as_pagination"
	"github.com/watermint/toolbox/domain/asana/model/mo_team"
	"github.com/watermint/toolbox/domain/asana/model/mo_workspace"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Team interface {
	Resolve(gid string) (team *mo_team.Team, err error)
	List(opt ...Opt) (teams []*mo_team.Team, err error)
}

type Opts struct {
	Workspace string `url:"workspace,omitempty"`
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

func New(ctx as_client.Client) Team {
	return &teamImpl{
		ctx: ctx,
	}
}

type teamImpl struct {
	ctx as_client.Client
	ws  *mo_workspace.Workspace
}

func (z teamImpl) Resolve(gid string) (team *mo_team.Team, err error) {
	team = &mo_team.Team{}
	res := z.ctx.Get("teams/" + gid)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	err = res.Success().Json().FindModel("data", team)
	return
}

func (z teamImpl) List(opt ...Opt) (teams []*mo_team.Team, err error) {
	opts := Opts{}.Apply(opt...)
	pg := as_pagination.New(z.ctx).WithEndpoint("organizations/" + opts.Workspace + "/teams")
	teams = make([]*mo_team.Team, 0)
	err = pg.OnData(func(entry es_json.Json) error {
		t := &mo_team.Team{}
		if errM := entry.Model(t); errM != nil {
			return errM
		} else {
			teams = append(teams, t)
			return nil
		}
	})
	return
}
