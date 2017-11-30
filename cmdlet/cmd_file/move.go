package cmd_file

import (
	"flag"
	"fmt"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/infra"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/cihub/seelog"
)

type CmdFileMove struct {
	apiContext   *api.ApiContext
	infraContext *infra.InfraContext
	ParamSrc     *api.DropboxPath
	ParamDest    *api.DropboxPath
}

func NewCmdFileMove() *CmdFileMove {
	c := CmdFileMove{
		infraContext: &infra.InfraContext{},
	}
	return &c
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

	c.infraContext.PrepareFlags(f)

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
	c.ParamSrc = &api.DropboxPath{Path: remainder[0]}
	c.ParamDest = &api.DropboxPath{Path: remainder[1]}
	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	c.apiContext, err = c.infraContext.LoadOrAuthDropboxFull()
	if err != nil {
		seelog.Warnf("Unable to acquire token  : error[%s]", err)
		return &cmdlet.CommandError{
			Context:     cc,
			ReasonTag:   "auth/auth_failed",
			Description: fmt.Sprintf("Unable to acquire token : error[%s].", err),
		}
	}

	return c.move(cc)
}

func (c *CmdFileMove) move(cc cmdlet.CommandletContext) error {
	arg := files.NewRelocationArg(c.ParamSrc.CleanPath(), c.ParamDest.CleanPath())
	arg.Autorename = false
	arg.AllowSharedFolder = true
	arg.AllowOwnershipTransfer = true

	seelog.Debugf("Move from[%s] to[%s] (param src[%s] dest[%s])", arg.FromPath, arg.ToPath, c.ParamSrc, c.ParamDest)

	_, err := c.apiContext.FilesMoveV2(arg)
	if err != nil {
		seelog.Warnf("Unable to move file(s) : error[%s]", err)
		return &cmdlet.CommandError{
			Context:     cc,
			ReasonTag:   "file/move/error",
			Description: fmt.Sprintf("Unable to move file(s) : error[%s].", err),
		}
	}

	return nil
}
