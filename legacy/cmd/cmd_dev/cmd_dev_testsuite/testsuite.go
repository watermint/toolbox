package cmd_dev_testsuite

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdDevTestSuite() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "testsuite",
		CommandDesc: "cmd.dev.testsuite.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdDevTestSuiteClean{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdDevTestSuiteCreate{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
