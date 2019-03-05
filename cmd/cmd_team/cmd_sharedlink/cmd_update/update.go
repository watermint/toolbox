package cmd_update

import "github.com/watermint/toolbox/cmd"

func NewCmdMemberSharedLinkUpdate() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "update",
		CommandDesc: "cmd.team.sharedlink.update.desc",
		SubCommands: []cmd.Commandlet{
			&CmdTeamSharedLinkUpdateExpire{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
