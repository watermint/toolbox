package cmd_team

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/infra"
)

type CmdTeamInfo struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     cmdlet.Report
}

func (CmdTeamInfo) Name() string {
	return "info"
}

func (CmdTeamInfo) Desc() string {
	return "Team info"
}

func (CmdTeamInfo) Usage() string {
	return ""
}

func (c *CmdTeamInfo) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdTeamInfo) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()

	//apiMgmt, err := ec.LoadOrAuthBusinessInfo()
	//if err != nil {
	//	return
	//}

	//c.report.DataHeaders = []string{
	//	"team_id",
	//	"info.name",
	//	"info.num_licensed_users",
	//	"info.num_provisioned_users",
	//}
	//
	//rt, rs, err := c.report.ReportStages()
	//if err != nil {
	//	return
	//}
	//
	//stages := []workflow.Worker{
	//	&team.WorkerTeamInfo{
	//		ApiManagement: apiMgmt,
	//		NextTask:      rt,
	//	},
	//}
	//
	//stages = append(stages, rs...)
	//
	//p := workflow.Pipeline{
	//	Infra:  ec,
	//	Stages: stages,
	//}
	//
	//p.Init()
	//defer p.Close()
	//
	//p.Enqueue(
	//	workflow.MarshalTask(
	//		team.WORKER_TEAM_INFO,
	//		team.WORKER_TEAM_INFO,
	//		team.ContextTeamInfo{},
	//	),
	//)
	//p.Loop()
}
