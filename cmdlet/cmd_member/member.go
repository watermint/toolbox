package cmd_member

import "github.com/watermint/toolbox/cmdlet"

func NewCmdMember() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "member",
		CommandDesc: "Dropbox Business member management",
		SubCommands: []cmdlet.Commandlet{
			&CmdMemberInvite{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			&CmdMemberList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
