package cmd_teamfolder

import (
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_teamfolder/cmd_teamfolder_file"
)

func NewCmdTeamFolder() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "teamfolder",
		CommandDesc: "cmd.teamfolder.desc",
		SubCommands: []cmd.Commandlet{
			&CmdTeamFolderList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamFolderSize{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamFolderMirror{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamFolderArchive{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamFolderPermDelete{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			cmd_teamfolder_file.NewCmdTeamFolderFile(),
		},
	}
}
