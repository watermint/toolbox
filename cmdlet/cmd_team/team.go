package cmd_team

import (
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_team/cmd_audit"
	"github.com/watermint/toolbox/cmdlet/cmd_team/cmd_linkedapp"
	"github.com/watermint/toolbox/cmdlet/cmd_team/cmd_namespace"
	"github.com/watermint/toolbox/cmdlet/cmd_team/cmd_sharedlink"
	"github.com/watermint/toolbox/cmdlet/cmd_team/cmd_teamfolder"
)

func NewCmdTeam() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "team",
		CommandDesc: "Dropbox Business Team management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamInfo{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			&CmdTeamFeature{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			cmd_audit.NewCmdTeamAudit(),
			cmd_linkedapp.NewCmdMemberLinkedApp(),
			cmd_sharedlink.NewCmdTeamSharedLink(),
			cmd_teamfolder.NewCmdTeamTeamFolder(),
			cmd_namespace.NewCmdTeamNamespace(),
		},
	}
}
