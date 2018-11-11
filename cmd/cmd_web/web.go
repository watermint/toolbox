package cmd_web

import (
	"github.com/watermint/toolbox/cmd"
)

func NewCmdWeb() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "web",
		CommandDesc: "Web Interface",
		SubCommands: []cmd.Commandlet{
			&CmdWebStart{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
