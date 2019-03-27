package cmd_member_quota

import "github.com/watermint/toolbox/cmd"

func NewCmdMemberQuota() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "quota",
		CommandDesc: "cmd.member.quota.desc",
		SubCommands: []cmd.Commandlet{
			&CmdMemberQuotaUpdate{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
