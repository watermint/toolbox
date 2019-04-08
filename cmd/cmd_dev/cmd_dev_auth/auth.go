package cmd_dev_auth

import "github.com/watermint/toolbox/cmd"

func NewCmdDevAuth() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "auth",
		CommandDesc: "cmd.dev.auth.desc",
		SubCommands: []cmd.Commandlet{
			&CmdDevAuthAppKey{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
