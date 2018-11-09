package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api/dbx_member"
	"github.com/watermint/toolbox/report"
)

type CmdMemberInvite struct {
	*cmdlet.SimpleCommandlet
	optCsv    string
	optSilent bool
	report    report.Factory
}

func (c *CmdMemberInvite) Name() string {
	return "invite"
}

func (c *CmdMemberInvite) Desc() string {
	return "Invite members"
}

func (c *CmdMemberInvite) Usage() string {
	return `{{.Command}} -csv MEMBER_FILENAME`
}

func (c *CmdMemberInvite) FlagConfig(f *flag.FlagSet) {
	descCsv := "CSV file name"
	f.StringVar(&c.optCsv, "csv", "", descCsv)

	descSilent := "Silent provisioning"
	f.BoolVar(&c.optSilent, "silent", false, descSilent)

	c.report.FlagConfig(f)
}

func (c *CmdMemberInvite) Exec(args []string) {
	if c.optCsv == "" {
		c.Log().Error("Please specify input csv")
		c.PrintUsage(c)
		return
	}

	apiMgmt, err := c.ExecContext.LoadOrAuthBusinessManagement()
	if err != nil {
		return
	}

	mp := MembersProvision{
		Logger: c.Log(),
	}
	err = mp.LoadCsv(c.optCsv)
	if err != nil {
		return
	}

	c.report.Init(c.Log())
	defer c.report.Close()

	memberReport := MemberReport{
		Report: &c.report,
	}

	members := mp.Members
	invites := make([]*dbx_member.InviteMember, len(members))

	for i, m := range members {
		invites[i] = m.InviteMember(c.optSilent)
	}

	mi := dbx_member.MembersInvite{
		OnFailure: memberReport.HandleFailure,
		OnSuccess: memberReport.HandleSuccess,
		OnError:   c.DefaultErrorHandler,
	}
	if !mi.Invite(apiMgmt, invites) {
		c.Log().Warn("terminate operation due to error")
		// quit, in case of the error
		return
	}

}
