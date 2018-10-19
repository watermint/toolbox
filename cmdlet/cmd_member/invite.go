package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx/task/member"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/workflow"
)

type CmdMemberInvite struct {
	*cmdlet.SimpleCommandlet
	optCsv     string
	optSilent  bool
	apiContext *api.ApiContext
}

func (c *CmdMemberInvite) Name() string {
	return "invite"
}

func (c *CmdMemberInvite) Desc() string {
	return "Invite members"
}

func (c *CmdMemberInvite) Usage() string {
	return `-csv MEMBER_FILENAME`
}

func (c *CmdMemberInvite) FlagConfig(f *flag.FlagSet) {

	descCsv := "CSV file name"
	f.StringVar(&c.optCsv, "csv", "", descCsv)

	descSilent := "Silent provisioning"
	f.BoolVar(&c.optSilent, "silent", false, descSilent)
}

func (c *CmdMemberInvite) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()

	apiMgmt, err := ec.LoadOrAuthBusinessManagement()
	if err != nil {
		return
	}

	p := workflow.Pipeline{
		Infra: ec,
		Stages: []workflow.Worker{
			&member.WorkerTeamMemberInviteLoaderCsv{},
			&member.WorkerTeamMemberInvite{ApiManagement: apiMgmt, Silent: c.optSilent},
			&member.WorkerTeamMemberInviteResultAsync{ApiManagement: apiMgmt},
			&member.WorkerTeamMemberInviteResultReduce{},
		},
	}
	p.Init()
	defer p.Close()

	p.Enqueue(member.NewTaskTeamMemberInviteLoaderCsv(c.optCsv))
	p.Loop()
}
