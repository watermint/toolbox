package cmd_team

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
)

type CmdTeamFeature struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.DbxContext
	report     app_report.Factory
}

func (z *CmdTeamFeature) Name() string {
	return "feature"
}

func (z *CmdTeamFeature) Desc() string {
	return "cmd.team.feature.desc"
}

func (CmdTeamFeature) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamFeature) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamFeature) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessInfo)
	if err != nil {
		return
	}
	svc := sv_team.New(ctx)
	feature, err := svc.Feature()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	z.report.Init(z.ExecContext)
	z.report.Report(feature)
	z.report.Close()
}
