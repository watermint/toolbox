package cmd_group

import "github.com/watermint/toolbox/cmdlet"

type CmdGroup struct {
	*cmdlet.ParentCommandlet
}

func (c *CmdGroup) Name() string {
	return "group"
}

func (c *CmdGroup) Desc() string {
	return "Dropbox Business team group management"
}
