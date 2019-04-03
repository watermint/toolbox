package cmd_auth

import (
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_auth/cmd_auth_business"
)

func NewCmdAuth() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "auth",
		CommandDesc: "cmd.auth.desc",
		SubCommands: []cmd.Commandlet{
			&CmdAuthUser{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			cmd_auth_business.NewCmdAuthBusiness(),
		},
	}
}
