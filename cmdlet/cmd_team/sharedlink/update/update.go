package update

import "github.com/watermint/toolbox/cmdlet"

func NewCmdMemberSharedLinkUpdate() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "update",
		CommandDesc: "Update shared link",
		SubCommands: []cmdlet.Commandlet{
			&CmdTeamSharedLinkUpdateExpire{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
