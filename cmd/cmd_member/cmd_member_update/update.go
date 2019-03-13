package cmd_member_update

import "github.com/watermint/toolbox/cmd"

func NewCmdMemberUpdate() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "update",
		CommandDesc: "cmd.member.update.desc",
		SubCommands: []cmd.Commandlet{
			&CmdMemberUpdateEmail{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
