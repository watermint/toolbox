package project

import (
	"github.com/watermint/toolbox/domain/asana/api/as_conn"
	"github.com/watermint/toolbox/domain/asana/model/mo_project"
	"github.com/watermint/toolbox/domain/asana/model/mo_workspace"
	"github.com/watermint/toolbox/domain/asana/service/sv_project"
	"github.com/watermint/toolbox/domain/asana/service/sv_workspace"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type List struct {
	Peer              as_conn.ConnAsanaApi
	Workspace         mo_filter.Filter
	Projects          rp_model.RowReport
	ProgressWorkspace app_msg.Message
}

func (z *List) Preset() {
	z.Projects.SetModel(&mo_project.Project{})
	z.Workspace.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
}

func (z *List) listProject(prjCompact *mo_project.Project, c app_control.Control) error {
	svp := sv_project.New(z.Peer.Context())
	prj, err := svp.Resolve(prjCompact.Gid)
	if err != nil {
		return err
	}
	z.Projects.Row(prj)
	return nil
}

func (z *List) listProjects(ws *mo_workspace.Workspace, c app_control.Control) error {
	c.UI().Progress(z.ProgressWorkspace.With("Name", ws.Name))
	svp := sv_project.New(z.Peer.Context())
	prjs, err := svp.List(sv_project.Workspace(ws))
	if err != nil {
		return err
	}

	for _, prjCompact := range prjs {
		if err := z.listProject(prjCompact, c); err != nil {
			return err
		}
	}
	return nil
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()
	workspaces, err := sv_workspace.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	if err := z.Projects.Open(); err != nil {
		return err
	}

	for _, wsCompact := range workspaces {
		ws, err := sv_workspace.New(z.Peer.Context()).Resolve(wsCompact.Gid)
		if err != nil {
			return err
		}
		if !ws.IsOrganization {
			l.Debug("Skip non organization workspace", esl.Any("workspace", ws))
			continue
		}
		if z.Workspace.Accept(ws.Name) || z.Workspace.Accept(ws.Gid) {
			if err := z.listProjects(ws, c); err != nil {
				return err
			}
		}
	}

	return nil
}

func (z *List) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &List{}, "recipe-services-asana-workspace-project-list.json.gz", rc_recipe.NoCustomValues)
	if err != nil {
		return err
	}
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
