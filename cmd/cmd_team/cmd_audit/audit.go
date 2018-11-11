package cmd_audit

import "github.com/watermint/toolbox/cmd"

func NewCmdTeamAudit() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "audit",
		CommandDesc: "Dropbox Business team audit",
		SubCommands: []cmd.Commandlet{
			&CmdTeamAuditSharing{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
