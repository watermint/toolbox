package cmd_sharedlink

import "github.com/watermint/toolbox/cmd"

func NewCmdSharedLink() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "sharedlink",
		CommandDesc: "cmd.sharedlink.desc",
		SubCommands: []cmd.Commandlet{
			&CmdSharedLinkList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdSharedLinkCreate{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdSharedLinkRemove{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
