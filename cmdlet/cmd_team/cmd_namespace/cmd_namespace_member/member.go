package cmd_namespace_member

import (
	"github.com/watermint/toolbox/cmdlet"
)

func NewCmdTeamNamespaceMember() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "member",
		CommandDesc: "Team namespace member management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamNamespaceMemberList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
