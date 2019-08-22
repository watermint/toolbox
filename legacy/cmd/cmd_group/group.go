package cmd_group

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_group_member2 "github.com/watermint/toolbox/legacy/cmd/cmd_group/cmd_group_member"
)

func NewCmdGroup() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "group",
		CommandDesc: "cmd.group.desc",
		SubCommands: []cmd2.Commandlet{
			cmd_group_member2.NewCmdGroupMember(),
			&CmdGroupList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdGroupRemove{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
