package cmd_namespace_file

import "github.com/watermint/toolbox/cmdlet"

func NewCmdTeamNamespaceFile() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "file",
		CommandDesc: "Team namespace file management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamNamespaceFileList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
