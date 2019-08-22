package cmd_member_quota

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdMemberQuota() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "quota",
		CommandDesc: "cmd.member.quota.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdMemberQuotaUpdate{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
