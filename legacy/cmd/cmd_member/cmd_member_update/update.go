package cmd_member_update

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdMemberUpdate() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "update",
		CommandDesc: "cmd.member.update.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdMemberUpdateEmail{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
