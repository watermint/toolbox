package cmd_member

import (
	"github.com/watermint/toolbox/cmdlet"
)

func NewCmdMember() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "member",
		CommandDesc: "Dropbox Business member management",
		SubCommands: []cmdlet.Commandlet{
			&CmdMemberList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			&CmdMemberInvite{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			&CmdMemberRemove{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			&CmdMemberSync{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
