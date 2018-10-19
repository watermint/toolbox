package linkedapp

import "github.com/watermint/toolbox/cmdlet"

func NewCmdMemberLinkedApp() cmdlet.Commandlet {
	return &cmdlet.CommandletGroup{
		CommandName: "app",
		CommandDesc: "Dropbox Business linked API app management",
		SubCommands: []cmdlet.Commandlet{
			&CmdMemberLinkedAppList{
				SimpleCommandlet: &cmdlet.SimpleCommandlet{},
			},
		},
	}
}
