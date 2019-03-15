package cmd_sharedfolder

import "github.com/watermint/toolbox/cmd"

func NewSharedFolder() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "sharedfolder",
		CommandDesc: "cmd.sharedfolder.desc",
		SubCommands: []cmd.Commandlet{
			&CmdSharedFolderList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
