package cmd_sharedfolder

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_sharedfolder_member2 "github.com/watermint/toolbox/legacy/cmd/cmd_sharedfolder/cmd_sharedfolder_member"
)

func NewSharedFolder() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "sharedfolder",
		CommandDesc: "cmd.sharedfolder.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdSharedFolderList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			cmd_sharedfolder_member2.NewCmdSharedFolderMember(),
		},
	}
}
