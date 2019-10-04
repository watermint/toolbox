package cmd_file

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdFile() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "file",
		CommandDesc: "cmd.file.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdFileMirror{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdFileCopy{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdFileMove{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdFileSave{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
