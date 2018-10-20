package namespace

import "github.com/watermint/toolbox/cmdlet"

func NewCmdTeamNamespace() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "namespace",
		CommandDesc: "Team namespace management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamNamespaceList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
