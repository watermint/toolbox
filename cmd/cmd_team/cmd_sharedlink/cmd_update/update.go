package cmd_update

import "github.com/watermint/toolbox/cmd"

func NewCmdMemberSharedLinkUpdate() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "update",
		CommandDesc: "Update shared link",
		SubCommands: []cmd.Commandlet{
			&CmdTeamSharedLinkUpdateExpire{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
