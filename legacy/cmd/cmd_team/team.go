package cmd_team

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_audit2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_audit"
	cmd_bulk2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_bulk"
	cmd_device2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_device"
	cmd_linkedapp2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_linkedapp"
	cmd_migrate2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_migrate"
	cmd_namespace2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_namespace"
	cmd_sharedlink2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_sharedlink"
)

func NewCmdTeam() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "team",
		CommandDesc: "cmd.team.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamInfo{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamFeature{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			cmd_audit2.NewCmdTeamAudit(),
			cmd_device2.NewCmdTeamDevice(),
			cmd_linkedapp2.NewCmdMemberLinkedApp(),
			cmd_sharedlink2.NewCmdTeamSharedLink(),
			cmd_namespace2.NewCmdTeamNamespace(),
			cmd_migrate2.NewCmdTeamMigrate(),
			cmd_bulk2.NewCmdTeamBulk(),
		},
	}
}
