package cmd_audit

import "github.com/watermint/toolbox/cmd"

func NewCmdTeamAudit() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "audit",
		CommandDesc: "cmd.team.audit.desc",
		SubCommands: []cmd.Commandlet{
			&CmdTeamAuditSharing{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
