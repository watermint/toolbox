package cmd_team

import (
	"flag"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx/task/team"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
)

type CmdTeamFeature struct {
	*cmdlet.SimpleCommandlet

	apiContext *api.ApiContext
	report     cmdlet.Report
}

func (c *CmdTeamFeature) Name() string {
	return "feature"
}

func (c *CmdTeamFeature) Desc() string {
	return "List team feature values"
}

func (CmdTeamFeature) Usage() string {
	return ""
}

func (c *CmdTeamFeature) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdTeamFeature) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()

	apiMgmt, err := ec.LoadOrAuthBusinessInfo()
	if err != nil {
		return
	}

	c.report.DataHeaders = []string{}

	rt, rs, err := c.report.ReportStages()
	if err != nil {
		return
	}

	stages := []workflow.Worker{
		&team.WorkerTeamFeatures{
			ApiManagement: apiMgmt,
			NextTask:      rt,
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
			team.WORKER_TEAM_FEATURES,
			team.WORKER_TEAM_FEATURES,
			team.ContextTeamFeature{},
		),
	)
	p.Loop()
}
