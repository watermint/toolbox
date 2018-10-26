package cmd_sharedlink

import (
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_team/cmd_sharedlink/cmd_update"
)

func NewCmdTeamSharedLink() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "sharedlink",
		CommandDesc: "Member shared link management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamSharedLinkList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			cmd_update.NewCmdMemberSharedLinkUpdate(),
		},
	}
}
