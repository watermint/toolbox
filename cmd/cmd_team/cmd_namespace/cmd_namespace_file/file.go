package cmd_namespace_file

import "github.com/watermint/toolbox/cmd"

func NewCmdTeamNamespaceFile() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "file",
		CommandDesc: "Team namespace file management",
		SubCommands: []cmd.Commandlet{
			&CmdTeamNamespaceFileList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamNamespaceFileSize{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
