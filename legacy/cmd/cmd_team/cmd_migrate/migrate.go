package cmd_migrate

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdTeamMigrate() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "migrate",
		CommandDesc: "cmd.team.migrate.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamMigrateCheck{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamMigrateContent{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamMigrateTransfer{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamMigrateVerify{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamMigrateBridge{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamMigratePermission{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
