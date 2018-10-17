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
)

type CmdMemberInvite struct {
	optCsv       string
	optSilent    bool
	apiContext   *api.ApiContext
	infraContext *infra.InfraContext
}

func NewCmdMemberInvite() *CmdMemberInvite {
	c := CmdMemberInvite{
		infraContext: &infra.InfraContext{},
	}
	return &c
}

func (c *CmdMemberInvite) Name() string {
	return "invite"
}

func (c *CmdMemberInvite) Desc() string {
	return "Invite members"
}

func (c *CmdMemberInvite) UsageTmpl() string {
	return `
Usage: {{.Command}} -csv MEMBER_FILENAME
`
}

func (c *CmdMemberInvite) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	descCsv := "CSV file name"
	f.StringVar(&c.optCsv, "csv", "", descCsv)

	descSilent := "Silent provisioning"
	f.BoolVar(&c.optSilent, "silent", false, descSilent)

	c.infraContext.PrepareFlags(f)
	return f
}

func (c *CmdMemberInvite) Exec(cc cmdlet.CommandletContext) error {
	_, err := cmdlet.ParseFlags(cc, c)
	if err != nil {
		return err
	}
	if c.optCsv == "" {
		return &cmdlet.CommandShowUsageError{
			Context:     cc,
			Instruction: "missing `-csv` option.",
		}
	}
	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	seelog.Debugf("invite:%s", util.MarshalObjectToString(c))

	apiMgmt, err := c.infraContext.LoadOrAuthBusinessManagement()
	if err != nil {
		seelog.Warnf("Unable to acquire token : error[%s]", err)
		return err
	}

	p := workflow.Pipeline{
		Infra: c.infraContext,
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

	return nil
}
