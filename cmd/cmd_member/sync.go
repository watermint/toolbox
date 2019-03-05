package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_member"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/report"
	"go.uber.org/zap"
)

type CmdMemberSync struct {
	*cmd.SimpleCommandlet
	optRemove string
	optWipe   bool
	optSilent bool
	report    report.Factory
	provision MembersProvision
}

func (z *CmdMemberSync) Name() string {
	return "sync"
}

func (z *CmdMemberSync) Desc() string {
	return "cmd.member.sync.desc"
}

func (z *CmdMemberSync) Usage() func(cmd.CommandUsage) {
	return z.provision.Usage()
}

func (z *CmdMemberSync) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
	z.provision.ec = z.ExecContext
	z.provision.FlagConfig(f)

	descSilent := z.ExecContext.Msg("cmd.member.sync.flag.silent").Text()
	f.BoolVar(&z.optSilent, "silent", false, descSilent)

	// first release includes only invite/update
	//descRemove := "Action for missing member (none|remove|detach)"
	//f.StringVar(&z.optRemove, "remove-action", "none", descRemove)
	//
	//descWipe := "Wipe data on remove user"
	//f.BoolVar(&z.optWipe, "wipe", false, descWipe)
}

func (z *CmdMemberSync) Exec(args []string) {
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

	z.report.Init(z.ExecContext)
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
		OnError:   z.DefaultErrorHandlerIgnoreError,
		OnFailure: memberReport.HandleFailure,
		OnSuccess: func(m *dbx_profile.Member) bool {
			z.report.Report(m)
			return true
		},
	}
	members := ml.ListAsMap(apiMgmt, false)
	invites := make([]*dbx_member.InviteMember, 0)

	for _, m := range z.provision.Members {
		if em, ok := members[m.Email]; ok {
			z.ExecContext.Msg("cmd.member.sync.flag.progress.update").WithData(struct {
				TeamMemberId string
				CurrentEmail string
				NewEmail     string
				GivenName    string
				Surname      string
			}{
				TeamMemberId: em.Profile.TeamMemberId,
				CurrentEmail: em.Profile.Email,
				NewEmail:     m.Email,
				GivenName:    m.GivenName,
				Surname:      m.Surname,
			})
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
