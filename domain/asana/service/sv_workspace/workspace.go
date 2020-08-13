package sv_workspace

import (
	"github.com/watermint/toolbox/domain/asana/api/as_context"
	"github.com/watermint/toolbox/domain/asana/api/as_pagination"
	"github.com/watermint/toolbox/domain/asana/model/mo_workspace"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Workspace interface {
	Resolve(gid string) (ws *mo_workspace.Workspace, err error)
	List() (ws []*mo_workspace.Workspace, err error)
}

func New(ctx as_context.Context) Workspace {
	return &workspaceImpl{
		ctx: ctx,
	}
}

type workspaceImpl struct {
	ctx as_context.Context
}

func (z workspaceImpl) Resolve(gid string) (ws *mo_workspace.Workspace, err error) {
	res := z.ctx.Get("workspaces/" + gid)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	ws = &mo_workspace.Workspace{}
	err = res.Success().Json().FindModel("data", ws)
	return
}

func (z workspaceImpl) List() (ws []*mo_workspace.Workspace, err error) {
	p := as_pagination.New(z.ctx).WithEndpoint("workspaces")
	ws = make([]*mo_workspace.Workspace, 0)
	err = p.OnData(func(entry es_json.Json) error {
		w := &mo_workspace.Workspace{}
		if errM := entry.Model(w); errM != nil {
			return errM
		} else {
			ws = append(ws, w)
			return nil
		}
	})
	return
}
