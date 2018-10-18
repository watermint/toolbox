package cmd_team

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx/task/team"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
	"github.com/watermint/toolbox/workflow/report"
)

type CmdTeamInfo struct {
	optReportPath   string
	optReportFormat string
	apiContext      *api.ApiContext
	infraContext    *infra.InfraContext
}

func NewCmdTeamInfo() *CmdTeamInfo {
	c := CmdTeamInfo{
		infraContext: &infra.InfraContext{},
	}
	return &c
}

func (c *CmdTeamInfo) Name() string {
	return "info"
}

func (c *CmdTeamInfo) Desc() string {
	return "Team Info"
}

func (c *CmdTeamInfo) UsageTmpl() string {
	return `
Usage: {{.Command}}
`
}

func (c *CmdTeamInfo) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	descReportPath := "Output file path of the report (default: STDOUT)"
	f.StringVar(&c.optReportPath, "report-path", "", descReportPath)

	descReportFormat := "Output file format (csv|jsonl) (default: jsonl)"
	f.StringVar(&c.optReportFormat, "report-format", "jsonl", descReportFormat)

	c.infraContext.PrepareFlags(f)
	return f
}

func (c *CmdTeamInfo) Exec(cc cmdlet.CommandletContext) error {
	_, err := cmdlet.ParseFlags(cc, c)
	if err != nil {
		return err
	}
	c.infraContext.Startup()
	defer c.infraContext.Shutdown()

	apiMgmt, err := c.infraContext.LoadOrAuthBusinessInfo()
	if err != nil {
		seelog.Warnf("Unable to acquire token : error[%s]", err)
		return err
	}

	reportTask := report.WORKER_REPORT_JSONL
	switch c.optReportFormat {
	case "jsonl":
		reportTask = report.WORKER_REPORT_JSONL

	case "csv":
		reportTask = report.WORKER_REPORT_CSV

	default:
		seelog.Warnf("Unsupported report format [%s]", c.optReportFormat)
		return err
	}

	p := workflow.Pipeline{
		Infra: c.infraContext,
		Stages: []workflow.Worker{
			&team.WorkerTeamInfo{
				ApiManagement: apiMgmt,
				NextTask:      reportTask,
			},
			&report.WorkerReportJsonl{
				ReportPath: c.optReportPath,
			},
			&report.WorkerReportCsv{
				ReportPath: c.optReportPath,
				DataHeaders: []string{
					"team_id",
					"info.name",
					"info.team_id",
					"info.num_licensed_users",
					"info.num_provisioned_users",
				},
			},
		},
	}

	p.Init()
	defer p.Close()

	p.Enqueue(
		workflow.MarshalTask(
			team.WORKER_TEAM_INFO,
			team.WORKER_TEAM_INFO,
			team.ContextTeamInfo{},
		),
	)
	p.Loop()

	return nil
}
