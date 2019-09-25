package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_file_size"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/usecase/uc_file_size"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report_legacy"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdTeamFolderSize struct {
	*cmd2.SimpleCommandlet
	optDepth int
	report   app_report_legacy.Factory
}

func (CmdTeamFolderSize) Name() string {
	return "size"
}

func (CmdTeamFolderSize) Desc() string {
	return "cmd.teamfolder.size.desc"
}

func (CmdTeamFolderSize) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdTeamFolderSize) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descOptDepth := z.ExecContext.Msg("cmd.teamfolder.size.flag.depth").T()
	f.IntVar(&z.optDepth, "depth", 2, descOptDepth)
}

func (z *CmdTeamFolderSize) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	admin, err := sv_profile.NewTeam(ctx).Admin()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	svn := sv_namespace.New(ctx)
	namespaces, err := svn.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	z.report.Init(z.ExecContext)
	defer z.report.Close()

	cta := ctx.AsAdminId(admin.TeamMemberId)

	for _, namespace := range namespaces {
		if namespace.NamespaceType != "team_folder" {
			continue
		}

		ucs := uc_file_size.New(cta.WithPath(api_context.Namespace(namespace.NamespaceId)))
		sizes, err := ucs.Size(mo_path.NewPath("/"), z.optDepth)
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}

		for _, size := range sizes {
			ns := mo_file_size.NewNamespaceSize(namespace, size)
			z.report.Report(ns)
		}
	}
}
