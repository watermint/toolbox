package project

import (
	"github.com/watermint/toolbox/domain/asana/api/as_conn"
	"github.com/watermint/toolbox/domain/asana/model/mo_project"
	"github.com/watermint/toolbox/domain/asana/model/mo_team"
	"github.com/watermint/toolbox/domain/asana/model/mo_workspace"
	"github.com/watermint/toolbox/domain/asana/service/sv_project"
	"github.com/watermint/toolbox/domain/asana/service/sv_team"
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
	Team              mo_filter.Filter
	Projects          rp_model.RowReport
	ProgressWorkspace app_msg.Message
	ProgressTeam      app_msg.Message
}

func (z *List) Preset() {
	z.Projects.SetModel(&mo_project.Project{})
	z.Workspace.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
	z.Team.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
}

func (z *List) listProjects(c app_control.Control, team *mo_team.Team) error {
	c.UI().Progress(z.ProgressTeam.With("Name", team.Name))
	svp := sv_project.New(z.Peer.Context())
	prjs, err := svp.List(sv_project.Team(team))
	if err != nil {
		return err
	}

	for _, prjCompact := range prjs {
		prj, err := svp.Resolve(prjCompact.Gid)
		if err != nil {
			return err
		}
		z.Projects.Row(prj)
	}
	return nil
}

func (z *List) listTeam(c app_control.Control, ws *mo_workspace.Workspace) error {
	c.UI().Progress(z.ProgressWorkspace.With("Name", ws.Name))
	teams, err := sv_team.New(z.Peer.Context()).List(sv_team.Workspace(ws))
	if err != nil {
		return err
	}

	for _, team := range teams {
		if z.Team.Accept(team.Name) || z.Team.Accept(team.Gid) {
			if err := z.listProjects(c, team); err != nil {
				return err
			}
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
			if err := z.listTeam(c, ws); err != nil {
				return err
			}
		}
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &List{}, "recipe-services-asana-team-project-list.json.gz", rc_recipe.NoCustomValues)
	if err != nil {
		return err
	}
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
