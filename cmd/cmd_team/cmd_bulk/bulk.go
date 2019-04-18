package cmd_bulk

import "github.com/watermint/toolbox/cmd"

func NewCmdTeamBulk() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "bulk",
		CommandDesc: "cmd.team.bulk.desc",
		SubCommands: []cmd.Commandlet{
			&CmdTeamBulkVerify{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamBulkExec{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
