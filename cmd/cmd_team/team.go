package cmd_team

import (
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_team/cmd_audit"
	"github.com/watermint/toolbox/cmd/cmd_team/cmd_linkedapp"
	"github.com/watermint/toolbox/cmd/cmd_team/cmd_namespace"
	"github.com/watermint/toolbox/cmd/cmd_team/cmd_sharedlink"
	"github.com/watermint/toolbox/cmd/cmd_team/cmd_teamfolder"
)

func NewCmdTeam() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "team",
		CommandDesc: "cmd.team.desc",
		SubCommands: []cmd.Commandlet{
			&CmdTeamInfo{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamFeature{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			cmd_audit.NewCmdTeamAudit(),
			cmd_linkedapp.NewCmdMemberLinkedApp(),
			cmd_sharedlink.NewCmdTeamSharedLink(),
			cmd_teamfolder.NewCmdTeamTeamFolder(),
			cmd_namespace.NewCmdTeamNamespace(),
		},
	}
}
