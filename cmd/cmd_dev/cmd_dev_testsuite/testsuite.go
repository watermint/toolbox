package cmd_dev_testsuite

import "github.com/watermint/toolbox/cmd"

func NewCmdDevTestSuite() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "testsuite",
		CommandDesc: "cmd.dev.testsuite.desc",
		SubCommands: []cmd.Commandlet{
			&CmdDevTestSuiteMigration{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
