package cmd_audit

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdTeamAudit() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "audit",
		CommandDesc: "cmd.team.audit.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamAuditEvents{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
