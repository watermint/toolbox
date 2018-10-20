package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_task/task/team"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
)

type CmdTeamNamespaceMemberList struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.ApiContext
	report     cmdlet.Report
}

func (CmdTeamNamespaceMemberList) Name() string {
	return "list"
}

func (CmdTeamNamespaceMemberList) Desc() string {
	return "List all namespaces of the team"
}

func (CmdTeamNamespaceMemberList) Usage() string {
	return ""
}

func (c *CmdTeamNamespaceMemberList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdTeamNamespaceMemberList) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()

	apiMgmt, err := ec.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	c.report.DataHeaders = []string{
		"team_member_id",
		"app_id",
	}

	rt, rs, err := c.report.ReportStages()
	if err != nil {
		return
	}

	stages := []workflow.Worker{
		&team.WorkerTeamNamespaceList{
			Api:      apiMgmt,
			NextTask: rt,
		},
	}

	stages = append(stages, rs...)

	p := workflow.Pipeline{
		Infra:  ec,
		Stages: stages,
	}

	p.Init()
	defer p.Close()

	p.Enqueue(
		workflow.MarshalTask(
			team.WORKER_TEAM_NAMESPACE_LIST,
			team.WORKER_TEAM_NAMESPACE_LIST,
			team.ContextTeamNamespaceList{},
		),
	)
	p.Loop()
}
