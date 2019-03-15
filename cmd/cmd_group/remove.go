package cmd_group

import (
	"flag"
	"github.com/watermint/toolbox/app/app_matcher"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_context_impl"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/report"
	"go.uber.org/zap"
)

type CmdGroupRemove struct {
	*cmd.SimpleCommandlet
	report  report.Factory
	matcher app_matcher.Matcher
}

func (z *CmdGroupRemove) Name() string {
	return "remove"
}

func (z *CmdGroupRemove) Desc() string {
	return "cmd.group.remove.desc"
}

func (z *CmdGroupRemove) Usage() func(usage cmd.CommandUsage) {
	return nil
}

func (z *CmdGroupRemove) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
	z.matcher.ExecContext = z.ExecContext
	z.matcher.FlagConfig(f)
}

func (z *CmdGroupRemove) Exec(args []string) {
	if err := z.matcher.Init(); err != nil {
		return
	}

	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	legacyCtx, err := au.Auth(dbx_auth.DropboxTokenBusinessManagement)
	if err != nil {
		return
	}
	ctx := api_context_impl.New(z.ExecContext, api_auth_impl.NewCompatible(legacyCtx.Token))

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_group.New(ctx)
	groups, err := svc.List()
	if err != nil {
		return
	}

	type ResultReport struct {
		GroupId   string `json:"group_id"`
		GroupName string `json:"group_name"`
		Result    string `json:"result"`
		Reason    string `json:"reason"`
	}

	for _, g := range groups {
		rr := ResultReport{
			GroupId:   g.GroupId,
			GroupName: g.GroupName,
		}
		gl := z.ExecContext.Log().With(zap.String("GroupId", g.GroupId), zap.String("GroupName", g.GroupName))
		if z.matcher.Match(g.GroupName) {
			if z.matcher.IsInteractive() {
				confirm := z.ExecContext.Msg("cmd.group.remove.prompt.confirm_remove").WithData(struct {
					Name  string
					Count int
				}{
					Name:  g.GroupName,
					Count: g.MemberCount,
				})
				if !confirm.AskConfirm() {
					gl.Debug("Skip: user cancelled removal")
					rr.Result = z.ExecContext.Msg("cmd.group.remove.report.skip").T()
					rr.Reason = z.ExecContext.Msg("cmd.group.remove.report.skip.detail").T()
					z.report.Report(rr)
					continue
				}
			}
			gl.Debug("Removing group")
			err := svc.Delete(g.GroupId)
			if err != nil {
				z.ExecContext.Msg("cmd.group.remove.err.unable_to_remove").WithData(struct {
					GroupName string
					Error     string
				}{
					GroupName: g.GroupName,
					Error:     ctx.ErrorMsg(err).T(),
				}).TellFailure()

				rr.Result = z.ExecContext.Msg("cmd.group.remove.report.failure").T()
				rr.Reason = ctx.ErrorMsg(err).T()
				z.report.Report(rr)
				continue
			}

			z.ExecContext.Msg("cmd.group.remove.progress.success").WithData(struct {
				GroupId   string
				GroupName string
			}{
				GroupId:   g.GroupId,
				GroupName: g.GroupName,
			}).TellSuccess()

			rr.Result = z.ExecContext.Msg("cmd.group.remove.report.success").T()
			rr.Reason = ""
			z.report.Report(rr)
		}
	}
}
