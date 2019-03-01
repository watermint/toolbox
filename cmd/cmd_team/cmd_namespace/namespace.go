package cmd_namespace

import (
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_team/cmd_namespace/cmd_namespace_file"
	"github.com/watermint/toolbox/cmd/cmd_team/cmd_namespace/cmd_namespace_member"
)

func NewCmdTeamNamespace() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "namespace",
		CommandDesc: "cmd.team.namespace.desc",
		SubCommands: []cmd.Commandlet{
			&CmdTeamNamespaceList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			cmd_namespace_member.NewCmdTeamNamespaceMember(),
			cmd_namespace_file.NewCmdTeamNamespaceFile(),
		},
	}
}
