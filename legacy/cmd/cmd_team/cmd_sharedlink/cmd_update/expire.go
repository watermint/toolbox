package cmd_update

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"time"
)

type CmdTeamSharedLinkUpdateExpire struct {
	*cmd2.SimpleCommandlet

	report  app_report.Factory
	optDays int
}

func (CmdTeamSharedLinkUpdateExpire) Name() string {
	return "expire"
}

func (CmdTeamSharedLinkUpdateExpire) Desc() string {
	return "cmd.team.sharedlink.update.expire.desc"
}

func (CmdTeamSharedLinkUpdateExpire) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdTeamSharedLinkUpdateExpire) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descDays := z.ExecContext.Msg("cmd.team.sharedlink.update.expire.flag.days").T()
	f.IntVar(&z.optDays, "days", 0, descDays)
}

func (z *CmdTeamSharedLinkUpdateExpire) Exec(args []string) {
	if z.optDays < 1 {
		z.ExecContext.Msg("cmd.team.sharedlink.update.expire.err.days_required").TellError()
		z.Log().Error("Please specify expiration date")
		return
	}

	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	svm := sv_member.New(ctx)
	members, err := svm.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	z.report.Init(z.ExecContext)
	defer z.report.Close()

	newExpire := api_util.RebaseTime(time.Now().Add(time.Duration(z.optDays*24) * time.Hour))
	for _, member := range members {
		mctx := ctx.AsMemberId(member.TeamMemberId)
		svs := sv_sharedlink.New(mctx)
		links, err := svs.List()
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}
		for _, link := range links {
			updated, err := svs.Update(link, sv_sharedlink.Expires(newExpire))
			if err != nil {
				api_util.UIMsgFromError(err).TellError()
				return
			}
			slm := mo_sharedlink.NewSharedLinkMember(updated, member)
			z.report.Report(slm)
		}
	}
}
