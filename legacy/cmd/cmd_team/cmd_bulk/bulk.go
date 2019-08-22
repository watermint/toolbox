package cmd_bulk

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdTeamBulk() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "bulk",
		CommandDesc: "cmd.team.bulk.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamBulkVerify{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamBulkExec{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
