package cmd_namespace_file

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
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdTeamNamespaceFileSize struct {
	*cmd2.SimpleCommandlet
	report   app_report.Factory
	optDepth int
}

func (CmdTeamNamespaceFileSize) Name() string {
	return "size"
}

func (CmdTeamNamespaceFileSize) Desc() string {
	return "cmd.team.namespace.file.size.desc"
}

func (CmdTeamNamespaceFileSize) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdTeamNamespaceFileSize) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	//descIncludeTeamFolder := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.include_team_folder").T()
	//f.BoolVar(&z.nsz.OptIncludeTeamFolder, "include-team-folder", true, descIncludeTeamFolder)
	//
	//descIncludeSharedFolder := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.include_shared_folder").T()
	//f.BoolVar(&z.nsz.OptIncludeSharedFolder, "include-shared-folder", true, descIncludeSharedFolder)
	//
	//descIncludeAppFolder := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.include_app_folder").T()
	//f.BoolVar(&z.nsz.OptIncludeAppFolder, "include-app-folder", false, descIncludeAppFolder)
	//
	//descIncludeMemberFolder := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.include_member_folder").T()
	//f.BoolVar(&z.nsz.OptIncludeMemberFolder, "include-member-folder", false, descIncludeMemberFolder)

	descOptDepth := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.depth").T()
	f.IntVar(&z.optDepth, "depth", 2, descOptDepth)
}

func (z *CmdTeamNamespaceFileSize) Exec(args []string) {
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

	ctf := ctx.AsAdminId(admin.TeamMemberId)

	for _, namespace := range namespaces {
		ucs := uc_file_size.New(ctf.WithPath(api_context.Namespace(namespace.NamespaceId)))
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
