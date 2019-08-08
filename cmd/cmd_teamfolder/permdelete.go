package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"go.uber.org/zap"
	"strings"
)

type CmdTeamFolderPermDelete struct {
	*cmd.SimpleCommandlet
}

func (z *CmdTeamFolderPermDelete) Name() string {
	return "permdelete"
}

func (z *CmdTeamFolderPermDelete) Desc() string {
	return "cmd.teamfolder.permdelete.desc"
}

func (z *CmdTeamFolderPermDelete) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamFolderPermDelete) FlagConfig(f *flag.FlagSet) {
}

func (z *CmdTeamFolderPermDelete) Exec(args []string) {
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

	for _, name := range args {
		nameLower := strings.ToLower(name)
		for _, folder := range folders {
			if strings.ToLower(folder.Name) == nameLower {
				err := svt.PermDelete(folder)
				if err != nil {
					z.Log().Warn("Unable to delete team folder", zap.String("name", folder.Name))
					api_util.UIMsgFromError(err).TellError()
				}
			}
		}
	}
}
