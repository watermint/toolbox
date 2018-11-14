package cmd_teamfolder_file

import "github.com/watermint/toolbox/cmd"

func NewCmdTeamTeamFolderFile() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "file",
		CommandDesc: "Team team folder file management",
		SubCommands: []cmd.Commandlet{
			&CmdTeamTeamFolderFileList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
