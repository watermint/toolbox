package cmd_team

import "github.com/watermint/toolbox/cmdlet"

func NewCmdTeam() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "team",
		CommandDesc: "Dropbox Business Team management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamInfo{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			&CmdTeamFeature{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
