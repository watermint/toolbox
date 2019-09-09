package cmd_teamfolder_file

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdTeamFolderFile() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "file",
		CommandDesc: "cmd.teamfolder.file.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamFolderFileList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
