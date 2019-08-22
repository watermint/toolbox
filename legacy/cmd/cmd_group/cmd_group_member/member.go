package cmd_group_member

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdGroupMember() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "member",
		CommandDesc: "cmd.group.member.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdGroupMemberList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdGroupMemberAdd{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
