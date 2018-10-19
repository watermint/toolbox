package cmd_group_member

import "github.com/watermint/toolbox/cmdlet"

func NewCmdGroupMember() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "member",
		CommandDesc: "Dropbox Business Group member management",
		SubCommands: []cmdlet.Commandlet{
			&CmdGroupMemberList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
