package cmd_dev_log

import "github.com/watermint/toolbox/cmd"

func NewCmdDevLog() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "log",
		CommandDesc: "cmd.dev.log.desc",
		SubCommands: []cmd.Commandlet{
			&CmdDevLogLoad{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
