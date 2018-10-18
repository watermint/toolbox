package cmd_group_member

import (
	"flag"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx/task/group"
	"github.com/watermint/toolbox/dbx/task/member"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
)

type CmdGroupMemberList struct {
	*cmdlet.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *api.ApiContext
	report            cmdlet.Report
}

func (c *CmdGroupMemberList) Name() string {
	return "list"
}

func (c *CmdGroupMemberList) Desc() string {
	return "List group members"
}

func (c *CmdGroupMemberList) Usage() string {
	return ""
}

func (c *CmdGroupMemberList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdGroupMemberList) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()

	apiMgmt, err := ec.LoadOrAuthBusinessInfo()
	if err != nil {
		return
	}

	c.report.DataHeaders = []string{
		"group_id",
		"group_name",
		"member.profile.team_member_id",
		"member.profile.account_id",
		"member.profile.email",
		"member.profile.name.given_name",
		"member.profile.name.surname",
		"member.access_type.\\.tag",
	}
	rt, rs, err := c.report.ReportStages()
	if err != nil {
		return
	}

	stages := []workflow.Worker{
		&group.WorkerTeamGroupList{
			Api:      apiMgmt,
			NextTask: group.WORKER_TEAM_GROUP_MEMBER_LIST,
		},
		&group.WorkerTeamGroupMemberList{
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
			member.WORKER_TEAM_MEMBER_LIST,
			member.WORKER_TEAM_MEMBER_LIST,
			member.ContextTeamMemberList{},
		),
	)
	p.Loop()
}
