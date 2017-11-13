package cmd_member

import (
	"github.com/watermint/toolbox/cmdlet"
)

type CmdMember struct {
	*cmdlet.ParentCommandlet
}

func (c *CmdMember) Name() string {
	return "member"
}

func (c *CmdMember) Desc() string {
	return "Dropbox Business team member management"
}
