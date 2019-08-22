package cmd_dev_log

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdDevLog() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "log",
		CommandDesc: "cmd.dev.log.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdDevLogLoad{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
