package cmd_group

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/infra"
)

type CmdGrouplist struct {
	*cmdlet.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            cmdlet.Report
}

func (c *CmdGrouplist) Name() string {
	return "list"
}

func (c *CmdGrouplist) Desc() string {
	return "List groups"
}

func (c *CmdGrouplist) Usage() string {
	return ""
}

func (c *CmdGrouplist) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)

	descCsv := "Include removed members"
	f.BoolVar(&c.optIncludeRemoved, "include-removed", false, descCsv)
}

func (c *CmdGrouplist) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()
	//
	//apiMgmt, err := ec.LoadOrAuthBusinessInfo()
	//if err != nil {
	//	return
	//}
	//
	//c.report.DataHeaders = []string{
	//	"group_id",
	//	"group_name",
	//	"group_management_type",
	//	"member_count",
	//}
	//
	//rt, rs, err := c.report.ReportStages()
	//if err != nil {
	//	return
	//}
	//
	//wkGroupList := &group.WorkerTeamGroupList{
	//	Api:      apiMgmt,
	//	NextTask: rt,
	//}
	//
	//stages := []workflow.Worker{
	//	wkGroupList,
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
	//		wkGroupList.Prefix(),
	//		wkGroupList.Prefix(),
	//		nil,
	//	),
	//)
	//p.Loop()
}
