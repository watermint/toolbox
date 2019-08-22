package cmd_namespace_member

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdTeamNamespaceMember() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "member",
		CommandDesc: "cmd.team.namespace.member.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamNamespaceMemberList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
