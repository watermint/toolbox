package cmd_teamfolder

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdTeamFolder() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "teamfolder",
		CommandDesc: "cmd.teamfolder.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamFolderSize{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamFolderMirror{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamFolderArchive{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamFolderPermDelete{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
