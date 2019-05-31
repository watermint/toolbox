package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"go.uber.org/zap"
	"strings"
)

type CmdTeamFolderArchive struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdTeamFolderArchive) Name() string {
	return "archive"
}

func (z *CmdTeamFolderArchive) Desc() string {
	return "cmd.teamfolder.archive.desc"
}

func (z *CmdTeamFolderArchive) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamFolderArchive) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamFolderArchive) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	svt := sv_teamfolder.New(ctx)
	folders, err := svt.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	z.report.Init(z.ExecContext)

	for _, name := range args {
		nameLower := strings.ToLower(name)
		for _, folder := range folders {
			if strings.ToLower(folder.Name) == nameLower {
				af, err := svt.Archive(folder)
				if err != nil {
					z.Log().Warn("Unable to archive team folder", zap.String("name", folder.Name))
					api_util.UIMsgFromError(err).TellError()
				} else {
					z.report.Report(af)
				}
			}
		}
	}
}
