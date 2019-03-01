package cmd_team

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_team"
	"github.com/watermint/toolbox/report"
)

type CmdTeamFeature struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.Context
	report     report.Factory
}

func (z *CmdTeamFeature) Name() string {
	return "feature"
}

func (z *CmdTeamFeature) Desc() string {
	return "cmd.team.feature.desc"
}

func (CmdTeamFeature) Usage() string {
	return ""
}

func (z *CmdTeamFeature) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamFeature) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiInfo, err := au.Auth(dbx_auth.DropboxTokenBusinessInfo)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	l := dbx_team.FeatureList{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(feature *dbx_team.Feature) bool {
			z.report.Report(feature)
			return true
		},
	}
	l.List(apiInfo)
}
