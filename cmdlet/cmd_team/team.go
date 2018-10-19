package cmd_team

import (
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_team/linkedapp"
	"github.com/watermint/toolbox/cmdlet/cmd_team/sharedlink"
)

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
			linkedapp.NewCmdMemberLinkedApp(),
			sharedlink.NewCmdMemberSharedLink(),
		},
	}
}
