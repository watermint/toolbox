package workspace

import (
	"github.com/watermint/toolbox/domain/asana/api/as_conn"
	"github.com/watermint/toolbox/domain/asana/model/mo_workspace"
	"github.com/watermint/toolbox/domain/asana/service/sv_workspace"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	rc_recipe.RemarkDeprecated
	Peer       as_conn.ConnAsanaApi
	Workspaces rp_model.RowReport
}

func (z *List) Preset() {
	z.Workspaces.SetModel(&mo_workspace.Workspace{})
}

func (z *List) Exec(c app_control.Control) error {
	workspaces, err := sv_workspace.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	if err := z.Workspaces.Open(); err != nil {
		return err
	}
	for _, wsCompact := range workspaces {
		ws, err := sv_workspace.New(z.Peer.Context()).Resolve(wsCompact.Gid)
		if err != nil {
			return err
		}
		z.Workspaces.Row(ws)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &List{}, "recipe-services-asana-workspace-list.json.gz", rc_recipe.NoCustomValues)
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
