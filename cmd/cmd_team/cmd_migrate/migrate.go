package cmd_migrate

import "github.com/watermint/toolbox/cmd"

func NewCmdTeamMigrate() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "migrate",
		CommandDesc: "cmd.team.migrate.desc",
		SubCommands: []cmd.Commandlet{
			&CmdTeamMigratePreflight{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamMigrateContent{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamMigrateTransfer{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
