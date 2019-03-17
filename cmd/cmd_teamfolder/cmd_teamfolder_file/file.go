package cmd_teamfolder_file

import "github.com/watermint/toolbox/cmd"

func NewCmdTeamFolderFile() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "file",
		CommandDesc: "cmd.teamfolder.file.desc",
		SubCommands: []cmd.Commandlet{
			&CmdTeamFolderFileList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
