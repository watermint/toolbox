package cmd_update

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_sharing"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/infra"
	"time"
)

type CmdTeamSharedLinkUpdateExpire struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
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

	c.report.Open()
	defer c.report.Close()

	type UpdateReport struct {
		MemberId     string `json:"member_id"`
		MemberEmail  string `json:"member_email"`
		SharedLinkId string `json:"shared_link_id"`
		OldExpires   string `json:"old_expires"`
		NewExpires   string `json:"new_expires"`
	}

	newExpire := dbx_api.RebaseTimeForAPI(time.Now().Add(time.Duration(c.optDays*24) * time.Hour))
	ml := dbx_team.MembersList{
		OnError: cmdlet.DefaultErrorHandler,
		OnEntry: func(member *dbx_profile.Member) bool {

			sl := dbx_sharing.SharedLinkList{
				AsMemberId:    member.Profile.TeamMemberId,
				AsMemberEmail: member.Profile.Email,
				OnError:       cmdlet.DefaultErrorHandler,
				OnEntry: func(link *dbx_sharing.SharedLink) bool {
					if c.filter.IsAcceptable(link) {
						newLink, ea, _ := link.UpdateExpire(apiMgmt, newExpire)
						if ea.IsFailure() {
							cmdlet.DefaultErrorHandlerIgnoreError(ea)
							return true
						}
						if newLink != nil {
							ur := UpdateReport{
								MemberId:     member.Profile.TeamMemberId,
								MemberEmail:  member.Profile.Email,
								SharedLinkId: link.SharedLinkId,
								OldExpires:   gjson.Get(string(link.Link), "expires").String(),
								NewExpires:   gjson.Get(string(newLink.Link), "expires").String(),
							}
							c.report.Report(ur)
						}
					}
					return true
				},
			}
			sl.List(apiMgmt)
			return true
		},
	}
	ml.List(apiMgmt, false)
}
