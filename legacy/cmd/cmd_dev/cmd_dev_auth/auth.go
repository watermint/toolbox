package cmd_dev_auth

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdDevAuth() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "auth",
		CommandDesc: "cmd.dev.auth.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdDevAuthAppKey{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
