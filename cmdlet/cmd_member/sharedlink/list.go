package sharedlink

import (
	"flag"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx/task/member"
	"github.com/watermint/toolbox/dbx/task/sharedlink"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
)

type CmdMemberSharedLinkList struct {
	*cmdlet.SimpleCommandlet

	apiContext *api.ApiContext
	report     cmdlet.Report
}

func (CmdMemberSharedLinkList) Name() string {
	return "list"
}

func (CmdMemberSharedLinkList) Desc() string {
	return "List all shared links of the team members' accounts"
}

func (CmdMemberSharedLinkList) Usage() string {
	return ""
}

func (c *CmdMemberSharedLinkList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdMemberSharedLinkList) Exec(ec *infra.ExecContext, args []string) {
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
		&member.WorkerTeamMemberList{
			ApiManagement: apiMgmt,
			NextTask:      workflow.WORKER_WORKFLOW_AS_MEMBER_ID,
		},
		&workflow.WorkerAsMemberIdDispatch{
			NextTask: sharedlink.WORKER_SHAREDLINK_LIST,
		},
		&sharedlink.WorkerSharedLinkList{
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
			member.WORKER_TEAM_MEMBER_LIST,
			member.WORKER_TEAM_MEMBER_LIST,
			member.ContextTeamMemberList{},
		),
	)
	p.Loop()
}
