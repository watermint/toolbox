package cmd_dev

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_dev_auth2 "github.com/watermint/toolbox/legacy/cmd/cmd_dev/cmd_dev_auth"
	cmd_dev_testsuite2 "github.com/watermint/toolbox/legacy/cmd/cmd_dev/cmd_dev_testsuite"
)

func NewCmdDev() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "dev",
		CommandDesc: "cmd.dev.desc",
		Hidden:      true,
		SubCommands: []cmd2.Commandlet{
			&CmdDevDoc{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			cmd_dev_testsuite2.NewCmdDevTestSuite(),
			cmd_dev_auth2.NewCmdDevAuth(),
		},
	}
}
