package cmd_file

import "github.com/watermint/toolbox/cmdlet"

type CmdFile struct {
	*cmdlet.ParentCommandlet
}

func (c *CmdFile) Name() string {
	return "file"
}

func (c *CmdFile) Desc() string {
	return "File operation"
}
