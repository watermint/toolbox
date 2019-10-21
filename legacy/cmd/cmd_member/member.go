package cmd_member

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_member_mirror2 "github.com/watermint/toolbox/legacy/cmd/cmd_member/cmd_member_mirror"
)

func NewCmdMember() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "member",
		CommandDesc: "cmd.member.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdMemberSync{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			cmd_member_mirror2.NewCmdMemberMirror(),
		},
	}
}
