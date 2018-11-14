package cmd_sharedlink

import (
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_team/cmd_sharedlink/cmd_update"
)

func NewCmdTeamSharedLink() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "sharedlink",
		CommandDesc: "Member shared link management",
		SubCommands: []cmd.Commandlet{
			&CmdTeamSharedLinkList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			cmd_update.NewCmdMemberSharedLinkUpdate(),
		},
	}
}
