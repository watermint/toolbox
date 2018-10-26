package cmd_teamfolder

import "github.com/watermint/toolbox/cmdlet"

func NewCmdTeamTeamFolder() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "teamfolder",
		CommandDesc: "Team folder management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamTeamFolderList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
