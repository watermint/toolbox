package cmd_linkedapp

import "github.com/watermint/toolbox/cmd"

func NewCmdMemberLinkedApp() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "app",
		CommandDesc: "cmd.team.linkedapp.desc",
		SubCommands: []cmd.Commandlet{
			&CmdMemberLinkedAppList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
