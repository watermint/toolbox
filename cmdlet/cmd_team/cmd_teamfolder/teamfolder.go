package cmd_teamfolder

import (
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_team/cmd_teamfolder/cmd_teamfolder_file"
)

func NewCmdTeamTeamFolder() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "teamfolder",
		CommandDesc: "Team folder management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamTeamFolderList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			cmd_teamfolder_file.NewCmdTeamTeamFolderFile(),
		},
	}
}
