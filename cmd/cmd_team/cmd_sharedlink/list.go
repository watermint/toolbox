package cmd_sharedlink

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_member"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/model/dbx_sharing"
	"github.com/watermint/toolbox/report"
)

type CmdTeamSharedLinkList struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.Context
	report     report.Factory
	filter     cmd.SharedLinkFilter
}

func (CmdTeamSharedLinkList) Name() string {
	return "list"
}

func (CmdTeamSharedLinkList) Desc() string {
	return "cmd.team.sharedlink.list.desc"
}

func (CmdTeamSharedLinkList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamSharedLinkList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
	z.filter.FlagConfig(f)
}

func (z *CmdTeamSharedLinkList) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}
	z.report.Init(z.ExecContext)
	defer z.report.Close()

	ml := dbx_member.MembersList{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(member *dbx_profile.Member) bool {
			sl := dbx_sharing.SharedLinkList{
				AsMemberId:    member.Profile.TeamMemberId,
				AsMemberEmail: member.Profile.Email,
				OnError:       z.DefaultErrorHandler,
				OnEntry: func(link *dbx_sharing.SharedLink) bool {
					if z.filter.IsAcceptable(link) {
						z.report.Report(link)
					}
					return true
				},
			}
			sl.List(apiFile)
			return true
		},
	}
	ml.List(apiFile, false)
}
