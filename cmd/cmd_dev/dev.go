package cmd_dev

import (
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_dev/cmd_dev_auth"
	"github.com/watermint/toolbox/cmd/cmd_dev/cmd_dev_log"
	"github.com/watermint/toolbox/cmd/cmd_dev/cmd_dev_testsuite"
)

func NewCmdDev() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "dev",
		CommandDesc: "cmd.dev.desc",
		Hidden:      true,
		SubCommands: []cmd.Commandlet{
			&CmdDevDoc{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			cmd_dev_testsuite.NewCmdDevTestSuite(),
			cmd_dev_auth.NewCmdDevAuth(),
			cmd_dev_log.NewCmdDevLog(),
		},
	}
}
