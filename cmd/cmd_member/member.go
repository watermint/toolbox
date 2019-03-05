package cmd_member

import (
	"github.com/watermint/toolbox/cmd"
)

func NewCmdMember() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "member",
		CommandDesc: "cmd.member.desc",
		SubCommands: []cmd.Commandlet{
			&CmdMemberList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdMemberInvite{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdMemberRemove{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdMemberDetach{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdMemberSync{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
