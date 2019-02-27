package cmd_update

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_member"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/model/dbx_sharing"
	"github.com/watermint/toolbox/report"
	"time"
)

type CmdTeamSharedLinkUpdateExpire struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.Context
	report     report.Factory
	filter     cmd.SharedLinkFilter
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

func (z *CmdTeamSharedLinkUpdateExpire) FlagConfig(f *flag.FlagSet) {
	z.report.FlagConfig(f)
	z.filter.FlagConfig(f)

	descDays := "Update and overwrite expiration date"
	f.IntVar(&z.optDays, "days", 0, descDays)
}

func (z *CmdTeamSharedLinkUpdateExpire) Exec(args []string) {
	if z.optDays < 1 {
		z.Log().Error("Please specify expiration date")
		return
	}
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	z.report.Init(z.Log())
	defer z.report.Close()

	type UpdateReport struct {
		MemberId     string `json:"member_id"`
		MemberEmail  string `json:"member_email"`
		SharedLinkId string `json:"shared_link_id"`
		OldExpires   string `json:"old_expires"`
		NewExpires   string `json:"new_expires"`
	}

	newExpire := dbx_api.RebaseTimeForAPI(time.Now().Add(time.Duration(z.optDays*24) * time.Hour))
	ml := dbx_member.MembersList{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(member *dbx_profile.Member) bool {

			sl := dbx_sharing.SharedLinkList{
				AsMemberId:    member.Profile.TeamMemberId,
				AsMemberEmail: member.Profile.Email,
				OnError:       z.DefaultErrorHandler,
				OnEntry: func(link *dbx_sharing.SharedLink) bool {
					if z.filter.IsAcceptable(link) {
						newLink, ea, _ := link.UpdateExpire(apiFile, newExpire)
						if ea.IsFailure() {
							z.DefaultErrorHandlerIgnoreError(ea)
							return true
						}
						if newLink != nil {
							ur := UpdateReport{
								MemberId:     member.Profile.TeamMemberId,
								MemberEmail:  member.Profile.Email,
								SharedLinkId: link.SharedLinkId,
								OldExpires:   link.Expires,
								NewExpires:   newLink.Expires,
							}
							z.report.Report(ur)
						}
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
