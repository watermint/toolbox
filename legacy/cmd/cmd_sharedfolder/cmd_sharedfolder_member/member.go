package cmd_sharedfolder_member

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdSharedFolderMember() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "member",
		CommandDesc: "cmd.sharedfolder.member.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdSharedFolderMemberList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
