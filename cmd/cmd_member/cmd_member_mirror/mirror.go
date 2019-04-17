package cmd_member_mirror

import "github.com/watermint/toolbox/cmd"

func NewCmdMemberMirror() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "mirror",
		CommandDesc: "cmd.member.mirror.desc",
		SubCommands: []cmd.Commandlet{
			&CmdMemberMirrorFiles{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
