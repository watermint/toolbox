package cmd_namespace

import (
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_team/cmd_namespace/cmd_namespace_file"
	"github.com/watermint/toolbox/cmdlet/cmd_team/cmd_namespace/cmd_namespace_member"
)

func NewCmdTeamNamespace() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "namespace",
		CommandDesc: "Team namespace management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamNamespaceList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			cmd_namespace_member.NewCmdTeamNamespaceMember(),
			cmd_namespace_file.NewCmdTeamNamespaceFile(),
		},
	}
}
