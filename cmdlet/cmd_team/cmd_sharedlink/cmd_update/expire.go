package cmd_update

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_task/task/member"
	"github.com/watermint/toolbox/dbx_task/task/sharedlink"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
)

type CmdTeamSharedLinkUpdateExpire struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.ApiContext
	report     cmdlet.Report
	filter     cmdlet.SharedLinkFilter
	optDays    int
}

func (CmdTeamSharedLinkUpdateExpire) Name() string {
	return "expire"
}

func (CmdTeamSharedLinkUpdateExpire) Desc() string {
	return "Update all shared link expire date of team members' accounts"
}

func (CmdTeamSharedLinkUpdateExpire) Usage() string {
	return ""
}

func (c *CmdTeamSharedLinkUpdateExpire) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
	c.filter.FlagConfig(f)

	descDays := "Update and overwrite expiration date"
	f.IntVar(&c.optDays, "days", 0, descDays)
}

func (c *CmdTeamSharedLinkUpdateExpire) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()
	if c.optDays < 1 {
		seelog.Warnf("Please specify expiration date")
		return
	}

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
	ft, fs, err := c.filter.FilterStages(sharedlink.WORKER_SHAREDLINK_UPDATE_EXPIRES)
	if err != nil {
		return
	}
	wrapUpTask := sharedlink.WORKER_SHAREDLINK_UPDATE_EXPIRES
	if ft != "" {
		wrapUpTask = ft
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
			NextTask: wrapUpTask,
		},
	}

	stages = append(stages, fs...)
	stages = append(
		stages,
		&sharedlink.WorkerSharedLinkUpdateExpires{
			Api:      apiMgmt,
			Days:     c.optDays,
			NextTask: rt,
		},
	)

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
