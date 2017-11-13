package cmd_file

import (
	"flag"
	"fmt"
	"github.com/watermint/toolbox/cmdlet"
)

type CmdFileRemove struct {
}

func (c *CmdFileRemove) Name() string {
	return "remove"
}

func (c *CmdFileRemove) Desc() string {
	return "Remove files"
}

func (c *CmdFileRemove) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	var recursive bool
	descRecursive := "Recurse into directories"
	f.BoolVar(&recursive, "recursive", true, descRecursive)
	f.BoolVar(&recursive, "r", true, descRecursive)

	return f
}

func (c *CmdFileRemove) UsageTmpl() string {
	return `
Usage: {{.Command}} PATH
`
}

func (c *CmdFileRemove) Exec(cc cmdlet.CommandletContext) error {
	return &cmdlet.CommandShowUsageError{
		cc,
		fmt.Sprintf("The command '%s' is not implemented yet.", c.Name()),
	} //TODO
}
