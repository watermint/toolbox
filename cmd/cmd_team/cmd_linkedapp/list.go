package cmd_linkedapp

import (
	"flag"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_member"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/model/dbx_team"
)

type CmdMemberLinkedAppList struct {
	*cmd.SimpleCommandlet

	apiContext         *dbx_api.DbxContext
	OptWithMemberEmail bool
	report             app_report.Factory
}

func (CmdMemberLinkedAppList) Name() string {
	return "list"
}

func (CmdMemberLinkedAppList) Desc() string {
	return "cmd.team.linkedapp.list.desc"
}

func (z *CmdMemberLinkedAppList) Usage() func(cmd.CommandUsage) {
	return func(u cmd.CommandUsage) {
		z.ExecContext.Msg("cmd.team.linkedapp.list.desc").Tell()
	}
}

func (z *CmdMemberLinkedAppList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descWithEmail := z.ExecContext.Msg("cmd.team.linkedapp.list.flag.with_email").T()
	f.BoolVar(&z.OptWithMemberEmail, "with-email", false, descWithEmail)
}

func (z *CmdMemberLinkedAppList) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	if z.OptWithMemberEmail {
		z.withMemberReport(apiFile)
	} else {
		z.plainReport(apiFile)
	}
}

func (z *CmdMemberLinkedAppList) plainReport(apiFile *dbx_api.DbxContext) {
	l := dbx_team.LinkedAppList{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(app *dbx_team.LinkedApp) bool {
			z.report.Report(app)
			return true
		},
	}
	l.List(apiFile)
}

func (z *CmdMemberLinkedAppList) withMemberReport(apiFile *dbx_api.DbxContext) {
	z.Log().Info("Prepare for expand members")
	members := make(map[string]*dbx_profile.Member)
	lm := dbx_member.MembersList{
		OnError: z.DefaultErrorHandlerIgnoreError,
		OnEntry: func(member *dbx_profile.Member) bool {
			members[member.Profile.TeamMemberId] = member
			return true
		},
	}
	if !lm.List(apiFile, false) {
		return
	}

	z.Log().Info("Listing App information")

	type App struct {
		TeamMemberId    string `json:"team_member_id"`
		TeamMemberEmail string `json:"team_member_email"`
		AppId           string `json:"app_id"`
		AppName         string `json:"app_name"`
		Publisher       string `json:"publisher"`
		PublisherUrl    string `json:"publisher_url"`
		IsAppFolder     bool   `json:"is_app_folder"`
		Linked          string `json:"linked"`
	}

	l := dbx_team.LinkedAppList{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(app *dbx_team.LinkedApp) bool {
			email := ""
			if m, ok := members[app.TeamMemberId]; ok {
				email = m.Profile.Email
			}
			json := gjson.ParseBytes(app.LinkedApiApp)
			a := App{
				TeamMemberId:    app.TeamMemberId,
				TeamMemberEmail: email,
				AppId:           app.LinkedApiAppId,
				AppName:         json.Get("app_name").String(),
				Publisher:       json.Get("publisher").String(),
				PublisherUrl:    json.Get("publisher_url").String(),
				IsAppFolder:     json.Get("is_app_folder").Bool(),
				Linked:          json.Get("linked").String(),
			}
			z.report.Report(a)
			return true
		},
	}
	l.List(apiFile)
}
