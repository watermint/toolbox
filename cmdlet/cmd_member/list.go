package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_task/task/member"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
)

type CmdMemberList struct {
	*cmdlet.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.ApiContext
	report            cmdlet.Report
}

func (c *CmdMemberList) Name() string {
	return "list"
}

func (c *CmdMemberList) Desc() string {
	return "List members"
}

func (CmdMemberList) Usage() string {
	return ""
}

func (c *CmdMemberList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)

	descCsv := "Include removed members"
	f.BoolVar(&c.optIncludeRemoved, "include-removed", false, descCsv)
}

func (c *CmdMemberList) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()

	apiMgmt, err := ec.LoadOrAuthBusinessInfo()
	if err != nil {
		return
	}

	c.report.DataHeaders = []string{
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
	}

	rt, rs, err := c.report.ReportStages()
	if err != nil {
		return
	}

	stages := []workflow.Worker{
		&member.WorkerTeamMemberList{
			ApiManagement:  apiMgmt,
			IncludeRemoved: c.optIncludeRemoved,
			NextTask:       rt,
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
			member.WORKER_TEAM_MEMBER_LIST,
			member.WORKER_TEAM_MEMBER_LIST,
			member.ContextTeamMemberList{},
		),
	)
	p.Loop()
}
