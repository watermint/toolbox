package cmd_member

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_member_mirror2 "github.com/watermint/toolbox/legacy/cmd/cmd_member/cmd_member_mirror"
	cmd_member_quota2 "github.com/watermint/toolbox/legacy/cmd/cmd_member/cmd_member_quota"
	cmd_member_update2 "github.com/watermint/toolbox/legacy/cmd/cmd_member/cmd_member_update"
)

func NewCmdMember() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "member",
		CommandDesc: "cmd.member.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdMemberRemove{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdMemberSync{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			cmd_member_update2.NewCmdMemberUpdate(),
			cmd_member_quota2.NewCmdMemberQuota(),
			cmd_member_mirror2.NewCmdMemberMirror(),
		},
	}
}
