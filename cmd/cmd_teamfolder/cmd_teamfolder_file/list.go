package cmd_teamfolder_file

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/usecase/uc_file_traverse"
)

type CmdTeamFolderFileList struct {
	*cmd.SimpleCommandlet
	optIncludeMediaInfo bool
	optIncludeDeleted   bool
	report              app_report.Factory
}

func (CmdTeamFolderFileList) Name() string {
	return "list"
}

func (CmdTeamFolderFileList) Desc() string {
	return "cmd.teamfolder.file.list.desc"
}

func (CmdTeamFolderFileList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamFolderFileList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	//descIncludeDeleted := z.ExecContext.Msg("cmd.teamfolder.file.list.flag.include_deleted").T()
	//f.BoolVar(&z.namespaceFile.OptIncludeDeleted, "include-deleted", false, descIncludeDeleted)
	//
	//descIncludeMediaInfo := z.ExecContext.Msg("cmd.teamfolder.file.list.flag.include_media_info").T()
	//f.BoolVar(&z.namespaceFile.OptIncludeMediaInfo, "include-media-info", false, descIncludeMediaInfo)
}

func (z *CmdTeamFolderFileList) Exec(args []string) {
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
	if err = z.report.Init(z.ExecContext); err != nil {
		return
	}

	for _, namespace := range namespaces {
		if namespace.NamespaceType != "team_folder" {
			continue
		}

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
