package cmd_member_mirror

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdMemberMirror() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "mirror",
		CommandDesc: "cmd.member.mirror.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdMemberMirrorFiles{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
