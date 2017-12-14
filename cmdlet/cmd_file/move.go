package cmd_file

import (
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
)

type CmdFileMove struct {
	optForce        bool
	optIgnoreErrors bool
	optVerbose      bool
	apiContext      *api.ApiContext
	infraContext    *infra.InfraContext
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

	descForce := "Force move even if for a shared folder"
	f.BoolVar(&c.optForce, "force", false, descForce)
	f.BoolVar(&c.optForce, "f", false, descForce)

	descContOnError := "Continue operation even if there are API errors"
	f.BoolVar(&c.optIgnoreErrors, "ignore-errors", false, descContOnError)

	descVerbose := "Showing files after they are moved"
	f.BoolVar(&c.optVerbose, "verbose", false, descVerbose)
	f.BoolVar(&c.optVerbose, "v", false, descVerbose)

	c.infraContext.PrepareFlags(f)

	return f
}

func (c *CmdFileMove) Exec(cc cmdlet.CommandletContext) error {
	remainder, err := cmdlet.ParseFlags(cc, c)
	if err != nil {
		return &cmdlet.CommandShowUsageError{
			Context:     cc,
			Instruction: err.Error(),
		}
	}
	if len(remainder) != 2 {
		return &cmdlet.CommandShowUsageError{
			Context:     cc,
			Instruction: "missing SRC DEST params",
		}
	}
	paramSrc := api.NewDropboxPath(remainder[0])
	paramDest := api.NewDropboxPath(remainder[1])
	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	seelog.Debugf("move:%s", util.MarshalObjectToString(c))
	c.apiContext, err = c.infraContext.LoadOrAuthDropboxFull()
	if err != nil {
		seelog.Warnf("Unable to acquire token  : error[%s]", err)
		return cmdlet.NewAuthFailedError(cc, err)
	}

	reloc := CmdRelocation{
		OptForce:        c.optForce,
		OptIgnoreErrors: c.optIgnoreErrors,
		ApiContext:      c.apiContext,
		RelocationFunc:  c.execMove,
	}
	err = reloc.Dispatch(paramSrc, paramDest)
	if err != nil {
		return c.composeError(cc, err)
	}
	return nil
}

func (c *CmdFileMove) composeError(cc cmdlet.CommandletContext, err error) error {
	seelog.Warnf("Unable to move file(s) : error[%s]", err)
	return &cmdlet.CommandError{
		Context:     cc,
		ReasonTag:   "file/move:" + err.Error(),
		Description: fmt.Sprintf("Unable to move file(s) : error[%s].", err),
	}
}

func (c *CmdFileMove) execMove(reloc *files.RelocationArg) (err error) {
	reloc.Autorename = false
	reloc.AllowSharedFolder = true
	reloc.AllowOwnershipTransfer = true

	seelog.Tracef("Move from[%s] to[%s]", reloc.FromPath, reloc.ToPath)

	_, err = c.apiContext.FilesMoveV2(reloc)
	if c.optVerbose && err == nil {
		seelog.Infof("moved[%s] -> [%s]", reloc.FromPath, reloc.ToPath)
	}
	return
}
