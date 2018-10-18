package cmd_group

import (
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_group/cmd_group_member"
)

type CmdGroup struct {
	*cmdlet.ParentCommandlet
}

func (c *CmdGroup) Name() string {
	return "group"
}

func (c *CmdGroup) Desc() string {
	return "Dropbox Business team group management"
}

func NewCmdGroup() cmdlet.Commandlet {
	return &CmdGroup{
		ParentCommandlet: &cmdlet.ParentCommandlet{
			SubCommands: []cmdlet.Commandlet{
				NewCmdGroupList(),
				cmd_group_member.NewCmdGroupMember(),
			},
		},
	}
}
