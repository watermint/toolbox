package cmd_group_member

import "github.com/watermint/toolbox/cmd"

func NewCmdGroupMember() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "member",
		CommandDesc: "cmd.group.member.desc",
		SubCommands: []cmd.Commandlet{
			&CmdGroupMemberList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
