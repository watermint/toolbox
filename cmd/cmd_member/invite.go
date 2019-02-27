package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_member"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/report"
)

type CmdMemberInvite struct {
	*cmd.SimpleCommandlet
	optSilent bool
	report    report.Factory
	provision MembersProvision
}

func (z *CmdMemberInvite) Name() string {
	return "invite"
}

func (z *CmdMemberInvite) Desc() string {
	return "Invite members"
}

func (z *CmdMemberInvite) Usage() string {
	return z.provision.Usage()
}

func (z *CmdMemberInvite) FlagConfig(f *flag.FlagSet) {
	descSilent := "Silent provisioning"
	f.BoolVar(&z.optSilent, "silent", false, descSilent)

	z.provision.FlagConfig(f)
	z.report.FlagConfig(f)
}

func (z *CmdMemberInvite) Exec(args []string) {
	z.provision.Logger = z.Log()
	err := z.provision.Load(args)
	if err != nil {
		z.PrintUsage(z.ExecContext, z)
		return
	}

	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiMgmt, err := au.Auth(dbx_auth.DropboxTokenBusinessManagement)
	if err != nil {
		return
	}

	z.report.Init(z.Log())
	defer z.report.Close()

	memberReport := MemberReport{
		Report: &z.report,
	}

	members := z.provision.Members
	invites := make([]*dbx_member.InviteMember, len(members))

	for i, m := range members {
		invites[i] = m.InviteMember(z.optSilent)
	}

	mi := dbx_member.MembersInvite{
		OnError:   z.DefaultErrorHandler,
		OnFailure: memberReport.HandleFailure,
		OnSuccess: func(m *dbx_profile.Member) bool {
			z.report.Report(m)
			return true
		},
	}
	if !mi.Invite(apiMgmt, invites) {
		z.Log().Warn("terminate operation due to error")
		// quit, in case of the error
		return
	}

}
