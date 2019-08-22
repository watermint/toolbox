package cmd_update

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdMemberSharedLinkUpdate() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "update",
		CommandDesc: "cmd.team.sharedlink.update.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamSharedLinkUpdateExpire{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
