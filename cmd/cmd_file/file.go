package cmd_file

import "github.com/watermint/toolbox/cmd"

func NewCmdFile() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "file",
		CommandDesc: "cmd.file.desc",
		SubCommands: []cmd.Commandlet{
			&CmdFileCompare{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdFileMetadata{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdFileMirror{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdFileList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdFileCopy{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdFileMove{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
