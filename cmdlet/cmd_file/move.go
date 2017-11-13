package cmd_file

import (
	"flag"
	"fmt"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/thinsdk"
)

type CmdFileMove struct {
	ParamSrc  *thinsdk.DropboxPath
	ParamDest *thinsdk.DropboxPath
}

func (c *CmdFileMove) Name() string {
	return "move"
}

func (c *CmdFileMove) Desc() string {
	return "Move files"
}

func (c *CmdFileMove) UsageTmpl() string {
	return `
Usage: {{.Command}} SRC DEST
`
}

func (c *CmdFileMove) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	var recursive bool
	descRecursive := "Recurse into directories"
	f.BoolVar(&recursive, "recursive", true, descRecursive)
	f.BoolVar(&recursive, "r", true, descRecursive)

	return f
}

func (c *CmdFileMove) Exec(cc cmdlet.CommandletContext) error {
	remainder, err := cmdlet.ParseFlags(cc, c)
	if err != nil {
		return &cmdlet.CommandShowUsageError{
			cc,
			err.Error(),
		}
	}
	if len(remainder) != 2 {
		return &cmdlet.CommandShowUsageError{
			cc,
			"missing SRC DEST params",
		}
	}
	c.ParamSrc = &thinsdk.DropboxPath{Path: remainder[0]}
	c.ParamDest = &thinsdk.DropboxPath{Path: remainder[1]}

	return &cmdlet.CommandError{
		Context:     cc,
		ReasonTag:   "not_implemented",
		Description: fmt.Sprintf("The command '%s' is not implemented yet.", c.Name()),
	} //TODO
}

func (c *CmdFileMove) move() error {
	//	arg := files.NewRelocationArg(c.ParamSrc.CleanPath(), c.ParamDest.CleanPath())
	return nil
}
