package sharedlink

import (
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_team/sharedlink/update"
)

func NewCmdTeamSharedLink() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "sharedlink",
		CommandDesc: "Member shared link management",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamSharedLinkList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
			update.NewCmdMemberSharedLinkUpdate(),
		},
	}
}
