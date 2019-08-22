package cmd_auth

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_auth_business2 "github.com/watermint/toolbox/legacy/cmd/cmd_auth/cmd_auth_business"
)

func NewCmdAuth() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "auth",
		CommandDesc: "cmd.auth.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdAuthUser{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			cmd_auth_business2.NewCmdAuthBusiness(),
		},
	}
}
