package cmd_group

import (
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_group/cmd_group_member"
)

func NewCmdGroup() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "group",
		CommandDesc: "Dropbox Business group management",
		SubCommands: []cmdlet.Commandlet{
			cmd_group_member.NewCmdGroupMember(),
			&CmdGrouplist{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
