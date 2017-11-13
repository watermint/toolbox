package cmd_file

import (
	"flag"
	"fmt"
	"github.com/watermint/toolbox/cmdlet"
)

type CmdFileCopy struct {
}

func (c *CmdFileCopy) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	var recursive bool
	descRecursive := "Recurse into directories"
	f.BoolVar(&recursive, "recursive", true, descRecursive)
	f.BoolVar(&recursive, "r", true, descRecursive)

	return f
}

func (c *CmdFileCopy) Name() string {
	return "copy"
}

func (c *CmdFileCopy) Desc() string {
	return "Copy files"
}

func (c *CmdFileCopy) UsageTmpl() string {
	return `
Usage: {{.Command}} SRC DEST
`
}

func (c *CmdFileCopy) Exec(cc cmdlet.CommandletContext) error {
	return &cmdlet.CommandShowUsageError{
		Instruction: fmt.Sprintf("The command '%s' is not implemented yet.", c.Name()),
	} //TODO
}
