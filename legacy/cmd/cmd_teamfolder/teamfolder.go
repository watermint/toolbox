package cmd_teamfolder

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_teamfolder_file2 "github.com/watermint/toolbox/legacy/cmd/cmd_teamfolder/cmd_teamfolder_file"
)

func NewCmdTeamFolder() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "teamfolder",
		CommandDesc: "cmd.teamfolder.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamFolderList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
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
			cmd_teamfolder_file2.NewCmdTeamFolderFile(),
		},
	}
}
