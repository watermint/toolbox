package cmd_team

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_audit2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_audit"
	cmd_device2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_device"
)

func NewCmdTeam() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "team",
		CommandDesc: "cmd.team.desc",
		SubCommands: []cmd2.Commandlet{
			cmd_audit2.NewCmdTeamAudit(),
			cmd_device2.NewCmdTeamDevice(),
		},
	}
}
