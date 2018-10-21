package cmd_linkedapp

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_task/member"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
)

type CmdMemberLinkedAppList struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     cmdlet.Report
}

func (CmdMemberLinkedAppList) Name() string {
	return "list"
}

func (CmdMemberLinkedAppList) Desc() string {
	return "List all applications linked to the team members' accounts"
}

func (CmdMemberLinkedAppList) Usage() string {
	return ""
}

func (c *CmdMemberLinkedAppList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdMemberLinkedAppList) Exec(ec *infra.ExecContext, args []string) {
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
		&member.WorkerTeamMemberLinkedApps{
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

	//p.Enqueue(
	//	workflow.MarshalTask(
	//		member.WORKER_TEAM_MEMBER_LINKEDAPPS,
	//		member.WORKER_TEAM_MEMBER_LINKEDAPPS,
	//		member.ContextTeamMemberLinkedApps{},
	//	),
	//)
	p.Loop()
}
