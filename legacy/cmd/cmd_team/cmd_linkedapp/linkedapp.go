package cmd_linkedapp

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdMemberLinkedApp() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "app",
		CommandDesc: "cmd.team.linkedapp.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdMemberLinkedAppList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
