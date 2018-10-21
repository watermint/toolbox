package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_task/teamfolder"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
)

type CmdTeamTeamFolderList struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     cmdlet.Report
}

func (CmdTeamTeamFolderList) Name() string {
	return "list"
}

func (CmdTeamTeamFolderList) Desc() string {
	return "List all team folder of the team"
}

func (CmdTeamTeamFolderList) Usage() string {
	return ""
}

func (c *CmdTeamTeamFolderList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdTeamTeamFolderList) Exec(ec *infra.ExecContext, args []string) {
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
		&teamfolder.WorkerTeamFolderList{
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
			teamfolder.WORKER_TEAMFOLDER_LIST,
			teamfolder.WORKER_TEAMFOLDER_LIST,
			teamfolder.ContextTeamFolderList{},
		),
	)
	p.Loop()
}
