package cmd_file

import "github.com/watermint/toolbox/cmd"

func NewCmdFile() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "file",
		CommandDesc: "cmd.file.desc",
		SubCommands: []cmd.Commandlet{
			&CmdFileMetadata{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
