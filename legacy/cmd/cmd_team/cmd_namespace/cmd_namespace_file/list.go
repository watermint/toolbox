package cmd_namespace_file

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/usecase/uc_file_traverse"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdTeamNamespaceFileList struct {
	*cmd2.SimpleCommandlet
	optIncludeMediaInfo bool
	optIncludeDeleted   bool
	report              app_report.Factory
}

func (CmdTeamNamespaceFileList) Name() string {
	return "list"
}

func (CmdTeamNamespaceFileList) Desc() string {
	return "cmd.team.namespace.file.list.desc"
}

func (CmdTeamNamespaceFileList) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdTeamNamespaceFileList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	//descIncludeDeleted := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_deleted").T()
	//f.BoolVar(&z.namespaceFile.OptIncludeDeleted, "include-deleted", false, descIncludeDeleted)
	//
	//descIncludeMediaInfo := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_media_info").T()
	//f.BoolVar(&z.namespaceFile.OptIncludeMediaInfo, "include-media-info", false, descIncludeMediaInfo)
	//
	//descIncludeTeamFolder := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_team_folder").T()
	//f.BoolVar(&z.namespaceFile.OptIncludeTeamFolder, "include-team-folder", true, descIncludeTeamFolder)
	//
	//descIncludeSharedFolder := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_shared_folder").T()
	//f.BoolVar(&z.namespaceFile.OptIncludeSharedFolder, "include-shared-folder", true, descIncludeSharedFolder)
	//
	//descIncludeAppFolder := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_app_folder").T()
	//f.BoolVar(&z.namespaceFile.OptIncludeAppFolder, "include-app-folder", false, descIncludeAppFolder)
	//
	//descIncludeMemberFolder := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_member_folder").T()
	//f.BoolVar(&z.namespaceFile.OptIncludeMemberFolder, "include-member-folder", false, descIncludeMemberFolder)
}

func (z *CmdTeamNamespaceFileList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	svt := sv_profile.NewTeam(ctx)
	admin, err := svt.Admin()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	cta := ctx.AsAdminId(admin.TeamMemberId)

	svn := sv_namespace.New(ctx)
	namespaces, err := svn.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	if err := z.report.Init(z.ExecContext); err != nil {
		return
	}

	for _, namespace := range namespaces {
		ctn := cta.WithPath(api_context.Namespace(namespace.NamespaceId))
		uct := uc_file_traverse.New(ctn)
		err := uct.Traverse(mo_path.NewPath(""), func(entry mo_file.Entry) error {
			z.report.Report(entry)
			return nil
		})
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}
	}
}
