package task

import (
	"github.com/watermint/toolbox/domain/asana/api/as_conn"
	"github.com/watermint/toolbox/domain/asana/model/mo_project"
	"github.com/watermint/toolbox/domain/asana/model/mo_task"
	"github.com/watermint/toolbox/domain/asana/model/mo_team"
	"github.com/watermint/toolbox/domain/asana/model/mo_workspace"
	"github.com/watermint/toolbox/domain/asana/service/sv_project"
	"github.com/watermint/toolbox/domain/asana/service/sv_task"
	"github.com/watermint/toolbox/domain/asana/service/sv_team"
	"github.com/watermint/toolbox/domain/asana/service/sv_workspace"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
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
	Project           mo_filter.Filter
	Tasks             rp_model.RowReport
	ProgressWorkspace app_msg.Message
	ProgressTeam      app_msg.Message
	ProgressProject   app_msg.Message
}

func (z *List) Preset() {
	z.Tasks.SetModel(&mo_task.Task{})
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
	z.Project.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
}

func (z *List) listTask(taskCompact *mo_task.Task, s eq_sequence.Stage) error {
	svt := sv_task.New(z.Peer.Context())
	task, err := svt.Resolve(taskCompact.Gid)
	if err != nil {
		return err
	}
	z.Tasks.Row(task)
	return nil
}

func (z *List) listTasks(prj *mo_project.Project, s eq_sequence.Stage) error {
	//c.UI().Progress(z.ProgressProject.With("Name", prj.Name))
	svt := sv_task.New(z.Peer.Context())
	tasks, err := svt.List(sv_task.Project(prj))
	if err != nil {
		return err
	}

	queue := s.Get("task").Batch(prj.Gid)
	for _, taskCompact := range tasks {
		queue.Enqueue(taskCompact)
	}
	return nil
}

func (z *List) listProjects(team *mo_team.Team, s eq_sequence.Stage) error {
	//c.UI().Progress(z.ProgressTeam.With("Name", team.Name))
	prjs, err := sv_project.New(z.Peer.Context()).List(sv_project.Team(team))
	if err != nil {
		return err
	}

	queue := s.Get("tasks").Batch(team.Gid)
	for _, prj := range prjs {
		if z.Project.Accept(prj.Name) || z.Project.Accept(prj.Gid) {
			queue.Enqueue(prj)
		}
	}
	return nil
}

func (z *List) listTeam(ws *mo_workspace.Workspace, s eq_sequence.Stage) error {
	//c.UI().Progress(z.ProgressWorkspace.With("Name", ws.Name))
	teams, err := sv_team.New(z.Peer.Context()).List(sv_team.Workspace(ws))
	if err != nil {
		return err
	}

	queue := s.Get("projects").Batch(ws.Gid)

	for _, team := range teams {
		if z.Team.Accept(team.Name) || z.Team.Accept(team.Gid) {
			queue.Enqueue(team)
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

	if err := z.Tasks.Open(); err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("team", z.listTeam, s)
		s.Define("projects", z.listProjects, s)
		s.Define("tasks", z.listTasks, s)
		s.Define("task", z.listTask, s)

		q := s.Get("team")

		for _, wsCompact := range workspaces {
			ws, err := sv_workspace.New(z.Peer.Context()).Resolve(wsCompact.Gid)
			if err != nil {
				l.Warn("Unable to retrieve workspace", esl.Error(err))
				continue
			}
			if !ws.IsOrganization {
				l.Debug("Skip non organization workspace", esl.Any("workspace", ws))
				continue
			}
			if z.Workspace.Accept(ws.Name) || z.Workspace.Accept(ws.Gid) {
				q.Enqueue(ws)
			}
		}
	})

	return nil
}

func (z *List) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &List{}, "recipe-services-asana-team-task-list.json.gz", rc_recipe.NoCustomValues)
	if err != nil {
		return err
	}
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
