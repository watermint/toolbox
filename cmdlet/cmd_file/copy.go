package cmd_file

import (
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
)

type CmdFileCopy struct {
	optForce        bool
	optIgnoreErrors bool
	optVerbose      bool
	optAutoRename   bool
	apiContext      *api.ApiContext
	infraContext    *infra.InfraContext
}

func NewCmdFileCopy() *CmdFileCopy {
	c := CmdFileCopy{
		infraContext: &infra.InfraContext{},
	}
	return &c
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

func (c *CmdFileCopy) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	descForce := "Force copy even if for a shared folder"
	f.BoolVar(&c.optForce, "force", false, descForce)
	f.BoolVar(&c.optForce, "f", false, descForce)

	descContOnError := "Continue operation even if there are API errors"
	f.BoolVar(&c.optIgnoreErrors, "ignore-errors", false, descContOnError)

	descVerbose := "Showing files after they are copied"
	f.BoolVar(&c.optVerbose, "verbose", false, descVerbose)
	f.BoolVar(&c.optVerbose, "v", false, descVerbose)

	descAutoRename := "Auto rename if an existing file found"
	f.BoolVar(&c.optAutoRename, "auto-rename", true, descAutoRename)
	f.BoolVar(&c.optAutoRename, "n", true, descAutoRename)

	c.infraContext.PrepareFlags(f)

	return f
}

func (c *CmdFileCopy) Exec(cc cmdlet.CommandletContext) error {
	remainder, err := cmdlet.ParseFlags(cc, c)
	if err != nil {
		return err
	}
	if len(remainder) != 2 {
		return &cmdlet.CommandShowUsageError{
			Context:     cc,
			Instruction: "missing SRC DEST params",
		}
	}
	//paramSrc := api.NewDropboxPath(remainder[0])
	//paramDest := api.NewDropboxPath(remainder[1])
	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	seelog.Debugf("copy:%s", util.MarshalObjectToString(c))
	c.apiContext, err = c.infraContext.LoadOrAuthDropboxFull()
	if err != nil {
		seelog.Warnf("Unable to acquire token  : error[%s]", err)
		return cmdlet.NewAuthFailedError(cc, err)
	}

	//reloc := CmdRelocation{
	//	OptForce:        c.optForce,
	//	OptIgnoreErrors: c.optIgnoreErrors,
	//	ApiContext:      c.apiContext,
	//	RelocationFunc:  c.execCopy,
	//}
	//err = reloc.Dispatch(paramSrc, paramDest)
	//if err != nil {
	//	return c.composeError(cc, err)
	//}
	return nil
}

func (c *CmdFileCopy) composeError(cc cmdlet.CommandletContext, err error) error {
	seelog.Warnf("Unable to copy file(s) : error[%s]", err)
	return &cmdlet.CommandError{
		Context:     cc,
		ReasonTag:   "file/copy:" + err.Error(),
		Description: fmt.Sprintf("Unable to copy file(s) : error[%s].", err),
	}
}

//func (c *CmdFileCopy) execCopy(reloc *files.RelocationArg) (err error) {
//	reloc.Autorename = c.optAutoRename
//	reloc.AllowSharedFolder = true
//	reloc.AllowOwnershipTransfer = true
//
//	seelog.Tracef("Copy from[%s] to[%s]", reloc.FromPath, reloc.ToPath)
//
//	_, err = c.apiContext.Files().CopyV2(reloc)
//	if c.optVerbose && err == nil {
//		seelog.Infof("copied[%s] -> [%s]", reloc.FromPath, reloc.ToPath)
//	}
//	return
//}
