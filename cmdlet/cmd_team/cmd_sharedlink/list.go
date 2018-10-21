package cmd_sharedlink

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_task/member"
	"github.com/watermint/toolbox/dbx_task/sharedlink"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
)

type CmdTeamSharedLinkList struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     cmdlet.Report
	filter     cmdlet.SharedLinkFilter
}

func (CmdTeamSharedLinkList) Name() string {
	return "list"
}

func (CmdTeamSharedLinkList) Desc() string {
	return "List all shared links of the team members' accounts"
}

func (CmdTeamSharedLinkList) Usage() string {
	return ""
}

func (c *CmdTeamSharedLinkList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
	c.filter.FlagConfig(f)
}

func (c *CmdTeamSharedLinkList) Exec(ec *infra.ExecContext, args []string) {
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
	ft, fs, err := c.filter.FilterStages(rt)
	if err != nil {
		return
	}
	wrapUpTask := rt
	if ft != "" {
		wrapUpTask = ft
	}

	wkSharedLinkList := &sharedlink.WorkerSharedLinkList{
		Api:      apiMgmt,
		NextTask: wrapUpTask,
	}
	wkAsMemberIdDispatch := &workflow.WorkerAsMemberIdDispatch{
		NextTask: wkSharedLinkList.Prefix(),
	}
	wkTeamMemberList := &member.WorkerTeamMemberList{
		Api:      apiMgmt,
		NextTask: wkAsMemberIdDispatch.Prefix(),
	}

	stages := []workflow.Worker{
		wkTeamMemberList,
		wkAsMemberIdDispatch,
		wkSharedLinkList,
	}

	stages = append(stages, fs...)
	stages = append(stages, rs...)

	p := workflow.Pipeline{
		Infra:  ec,
		Stages: stages,
	}

	p.Init()
	defer p.Close()

	p.Enqueue(
		workflow.MarshalTask(
			wkTeamMemberList.Prefix(),
			wkTeamMemberList.Prefix(),
			nil,
		),
	)
	p.Loop()
}
