package sharedlink

import "github.com/watermint/toolbox/cmdlet"

func NewCmdMemberSharedLink() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "sharedlink",
		CommandDesc: "Member shared link management",
		SubCommands: []cmdlet.Commandlet{
			&CmdMemberSharedLinkList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
