package cmd_group

import (
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_group/cmd_group_member"
)

func NewCmdGroup() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "group",
		CommandDesc: "cmd.group.desc",
		SubCommands: []cmd.Commandlet{
			cmd_group_member.NewCmdGroupMember(),
			&CmdGroupList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
