package cmd_auth_business

import "github.com/watermint/toolbox/cmd"

func NewCmdAuthBusiness() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "business",
		CommandDesc: "cmd.auth.business.desc",
		SubCommands: []cmd.Commandlet{
			&CmdAuthBusinessInfo{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdAuthBusinessFile{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdAuthBusinessAudit{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdAuthBusinessManagement{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
