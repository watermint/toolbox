package cmd_group_member

import "github.com/watermint/toolbox/cmdlet"

type CmdGroupMember struct {
	*cmdlet.ParentCommandlet
}

func (c *CmdGroupMember) Name() string {
	return "member"
}

func (c *CmdGroupMember) Desc() string {
	return "Dropbox Business team group member management"
}

func NewCmdGroupMember() cmdlet.Commandlet {
	return &CmdGroupMember{
		ParentCommandlet: &cmdlet.ParentCommandlet{
			SubCommands: []cmdlet.Commandlet{
				NewCmdGroupMemberList(),
			},
		},
	}
}
