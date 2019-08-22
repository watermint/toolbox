package cmd_auth_business

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdAuthBusiness() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "business",
		CommandDesc: "cmd.auth.business.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdAuthBusinessInfo{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdAuthBusinessFile{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdAuthBusinessAudit{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdAuthBusinessManagement{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
