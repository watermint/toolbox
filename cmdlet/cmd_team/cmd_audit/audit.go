package cmd_audit

import "github.com/watermint/toolbox/cmdlet"

func NewCmdTeamAudit() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "audit",
		CommandDesc: "Dropbox Business team audit",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamAuditSharing{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
