package cmd_teamfolder_file

import "github.com/watermint/toolbox/cmdlet"

func NewCmdTeamTeamFolderFile() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "file",
		CommandDesc: "Team team folder file management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamTeamFolderFileList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
