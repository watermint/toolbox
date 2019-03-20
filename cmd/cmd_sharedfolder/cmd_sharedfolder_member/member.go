package cmd_sharedfolder_member

import "github.com/watermint/toolbox/cmd"

func NewCmdSharedFolderMember() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "member",
		CommandDesc: "cmd.sharedfolder.member.desc",
		SubCommands: []cmd.Commandlet{
			&CmdSharedFolderMemberList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
