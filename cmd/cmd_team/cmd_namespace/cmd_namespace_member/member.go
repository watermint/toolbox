package cmd_namespace_member

import (
	"github.com/watermint/toolbox/cmd"
)

func NewCmdTeamNamespaceMember() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "member",
		CommandDesc: "Team namespace member management",
		SubCommands: []cmd.Commandlet{
			&CmdTeamNamespaceMemberList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
