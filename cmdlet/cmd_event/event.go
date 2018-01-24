package cmd_event

import "github.com/watermint/toolbox/cmdlet"

type CmdEvent struct {
	*cmdlet.ParentCommandlet
}

func (c *CmdEvent) Name() string {
	return "event"
}

func (c *CmdEvent) Desc() string {
	return "Dropbox Business events log"
}
