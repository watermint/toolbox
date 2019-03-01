package cmd_namespace_file

import "github.com/watermint/toolbox/cmd"

func NewCmdTeamNamespaceFile() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "file",
		CommandDesc: "cmd.team.namespace.file.desc",
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
