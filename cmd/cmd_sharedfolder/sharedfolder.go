package cmd_sharedfolder

import (
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_sharedfolder/cmd_sharedfolder_member"
)

func NewSharedFolder() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "sharedfolder",
		CommandDesc: "cmd.sharedfolder.desc",
		SubCommands: []cmd.Commandlet{
			&CmdSharedFolderList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			cmd_sharedfolder_member.NewCmdSharedFolderMember(),
		},
	}
}
