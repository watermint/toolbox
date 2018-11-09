package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api/dbx_member"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/report"
	"go.uber.org/zap"
)

type CmdMemberSync struct {
	*cmdlet.SimpleCommandlet
	optCsv    string
	optRemove string
	optWipe   bool
	optSilent bool
	report    report.Factory
}

func (z *CmdMemberSync) Name() string {
	return "sync"
}

func (z *CmdMemberSync) Desc() string {
	return "Sync member information with provided csv"
}

func (z *CmdMemberSync) Usage() string {
	return `{{.Command}} -csv MEMBER_FILENAME`
}

func (z *CmdMemberSync) FlagConfig(f *flag.FlagSet) {
	z.report.FlagConfig(f)

	descCsv := "CSV file name"
	f.StringVar(&z.optCsv, "csv", "", descCsv)

	descSilent := "Silent provisioning"
	f.BoolVar(&z.optSilent, "silent", false, descSilent)

	// first release includes only invite/update
	//descRemove := "Action for missing member (none|remove|detach)"
	//f.StringVar(&z.optRemove, "remove-action", "none", descRemove)
	//
	//descWipe := "Wipe data on remove user"
	//f.BoolVar(&z.optWipe, "wipe", false, descWipe)
}

func (z *CmdMemberSync) Exec(args []string) {
	if z.optCsv == "" {
		z.Log().Error("Please specify input csv")
		z.PrintUsage(z)
		return
	}
	apiMgmt, err := z.ExecContext.LoadOrAuthBusinessManagement()
	if err != nil {
		return
	}

	mp := MembersProvision{
		Logger: z.Log(),
	}
	err = mp.LoadCsv(z.optCsv)
	if err != nil {
		return
	}

	z.report.Init(z.Log())
	defer z.report.Close()

	memberReport := MemberReport{
		Report: &z.report,
	}

	ml := dbx_member.MembersList{
		OnError: z.DefaultErrorHandler,
	}
	mu := dbx_member.MemberUpdate{
		OnError: z.DefaultErrorHandler,
		OnSuccess: func(m *dbx_profile.Member) bool {
			z.report.Report(m)
			return true
		},
	}
	mi := dbx_member.MembersInvite{
		OnError:   z.DefaultErrorHandler,
		OnFailure: memberReport.HandleFailure,
		OnSuccess: memberReport.HandleSuccess,
	}
	members := ml.ListAsMap(apiMgmt, false)
	invites := make([]*dbx_member.InviteMember, 0)

	for _, m := range mp.Members {
		if em, ok := members[m.Email]; ok {
			z.Log().Info(
				"Updating member",
				zap.String("team_member_id", em.Profile.TeamMemberId),
				zap.String("current_email", em.Profile.Email),
				zap.String("new_email", m.Email),
				zap.String("new_given_name", m.GivenName),
				zap.String("new_surname", m.Surname),
			)
			email, um := m.UpdateMember()
			mu.Update(apiMgmt, email, um)
		} else {
			invites = append(
				invites,
				m.InviteMember(z.optSilent),
			)
		}
	}

	mi.Invite(apiMgmt, invites)
}
