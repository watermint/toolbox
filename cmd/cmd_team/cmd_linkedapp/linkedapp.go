package cmd_linkedapp

import "github.com/watermint/toolbox/cmd"

func NewCmdMemberLinkedApp() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "app",
		CommandDesc: "Dropbox Business linked API app management",
		SubCommands: []cmd.Commandlet{
			&CmdMemberLinkedAppList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
