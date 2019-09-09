package cmd_namespace

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_namespace_file2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_namespace/cmd_namespace_file"
	cmd_namespace_member2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_namespace/cmd_namespace_member"
)

func NewCmdTeamNamespace() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "namespace",
		CommandDesc: "cmd.team.namespace.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamNamespaceList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			cmd_namespace_member2.NewCmdTeamNamespaceMember(),
			cmd_namespace_file2.NewCmdTeamNamespaceFile(),
		},
	}
}
