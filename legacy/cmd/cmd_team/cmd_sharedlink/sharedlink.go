package cmd_sharedlink

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_update2 "github.com/watermint/toolbox/legacy/cmd/cmd_team/cmd_sharedlink/cmd_update"
)

func NewCmdTeamSharedLink() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "sharedlink",
		CommandDesc: "cmd.team.sharedlink.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamSharedLinkList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			cmd_update2.NewCmdMemberSharedLinkUpdate(),
		},
	}
}
