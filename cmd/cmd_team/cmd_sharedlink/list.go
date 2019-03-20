package cmd_sharedlink

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
)

type CmdTeamSharedLinkList struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.DbxContext
	report     app_report.Factory
	//	filter     cmd.SharedLinkFilter
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
	//	z.filter.FlagConfig(f) // TODO filtering
}

func (z *CmdTeamSharedLinkList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	svm := sv_member.New(ctx)
	members, err := svm.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}
	z.report.Init(z.ExecContext)
	defer z.report.Close()

	for _, member := range members {
		mctx := ctx.AsMemberId(member.TeamMemberId)
		svs := sv_sharedlink.New(mctx)
		links, err := svs.List()
		if err != nil {
			mctx.ErrorMsg(err).TellError()
			return
		}
		for _, link := range links {
			slm := mo_sharedlink.NewSharedLinkMember(link, member)
			z.report.Report(slm)
		}
	}
}
