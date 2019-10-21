package cmd_sharedlink

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdSharedLink() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "sharedlink",
		CommandDesc: "cmd.sharedlink.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdSharedLinkRemove{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
