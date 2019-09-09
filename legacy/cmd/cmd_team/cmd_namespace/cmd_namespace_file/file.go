package cmd_namespace_file

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdTeamNamespaceFile() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "file",
		CommandDesc: "cmd.team.namespace.file.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamNamespaceFileList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamNamespaceFileSize{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
