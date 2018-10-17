package cmd_member

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx/task/member"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/workflow"
	"github.com/watermint/toolbox/workflow/report"
)

type CmdMemberList struct {
	optIncludeRemoved bool
	optReportPath     string
	optReportFormat   string
	apiContext        *api.ApiContext
	infraContext      *infra.InfraContext
}

func NewCmdMemberList() *CmdMemberList {
	c := CmdMemberList{
		infraContext: &infra.InfraContext{},
	}
	return &c
}

func (c *CmdMemberList) Name() string {
	return "list"
}

func (c *CmdMemberList) Desc() string {
	return "List members"
}

func (c *CmdMemberList) UsageTmpl() string {
	return `
Usage: {{.Command}}
`
}

func (c *CmdMemberList) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	descCsv := "Include removed members"
	f.BoolVar(&c.optIncludeRemoved, "include-removed", false, descCsv)

	descReportPath := "Output file path of the report (default: STDOUT)"
	f.StringVar(&c.optReportPath, "report-path", "", descReportPath)

	descReportFormat := "Output file format (csv|jsonl) (default: jsonl)"
	f.StringVar(&c.optReportFormat, "report-format", "jsonl", descReportFormat)

	c.infraContext.PrepareFlags(f)
	return f
}

func (c *CmdMemberList) Exec(cc cmdlet.CommandletContext) error {
	_, err := cmdlet.ParseFlags(cc, c)
	if err != nil {
		return err
	}
	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	seelog.Debugf("invite:%s", util.MarshalObjectToString(c))

	apiMgmt, err := c.infraContext.LoadOrAuthBusinessManagement()
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
			&member.WorkerTeamMemberList{
				ApiManagement:  apiMgmt,
				IncludeRemoved: c.optIncludeRemoved,
				NextTask:       reportTask,
			},
			&report.WorkerReportJsonl{
				ReportPath: c.optReportPath,
			},
			&report.WorkerReportCsv{
				ReportPath: c.optReportPath,
				DataHeaders: []string{
					"member.profile.team_member_id",
					"member.profile.email",
					"member.profile.email_verified",
					"member.profile.status.\\.tag",
					"member.profile.name.given_name",
					"member.profile.name.surname",
					"member.profile.name.familiar_name",
					"member.profile.name.display_name",
					"member.profile.name.abbreviated_name",
					"member.profile.external_id",
					"member.profile.account_id",
					"member.profile.joined_on",
					"member.role.\\.tag",
				},
			},
		},
	}

	p.Init()
	defer p.Close()

	p.Enqueue(
		workflow.MarshalTask(
			member.WORKER_TEAM_MEMBER_LIST,
			member.WORKER_TEAM_MEMBER_LIST,
			member.ContextTeamMemberList{},
		),
	)
	p.Loop()

	return nil
}
