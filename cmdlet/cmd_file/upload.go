package cmd_file

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"os"
)

type CmdFileUpload struct {
	apiContext   *api.ApiContext
	infraContext *infra.InfraContext
}

func NewCmdFileUpload() *CmdFileUpload {
	c := CmdFileUpload{
		infraContext: &infra.InfraContext{},
	}
	return &c
}

func (c *CmdFileUpload) Name() string {
	return "upload"
}

func (c *CmdFileUpload) Desc() string {
	return "Upload files"
}

func (c *CmdFileUpload) UsageTmpl() string {
	return `
Usage: {{.Command}} LOCAL_PATH DROPBOX_PATH
`
}

func (c *CmdFileUpload) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	c.infraContext.PrepareFlags(f)

	return f
}

func (c *CmdFileUpload) Exec(cc cmdlet.CommandletContext) error {
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
			Instruction: "missing LOCAL_PATH DROPBOX_PATH params",
		}
	}
	paramLocal := remainder[0]
	paramDest := api.NewDropboxPath(remainder[1])
	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	seelog.Debugf("move:%s", util.MarshalObjectToString(c))
	c.apiContext, err = c.infraContext.LoadOrAuthDropboxFull()
	if err != nil {
		seelog.Warnf("Unable to acquire token  : error[%s]", err)
		return cmdlet.NewAuthFailedError(cc, err)
	}

	err = c.execUpload(paramLocal, paramDest)
	if err != nil {
		return c.composeError(cc, err)
	}
	return nil
}

func (c *CmdFileUpload) composeError(cc cmdlet.CommandletContext, err error) error {
	seelog.Warnf("Unable to move file(s) : error[%s]", err)
	return &cmdlet.CommandError{
		Context:     cc,
		ReasonTag:   "file/upload:" + err.Error(),
		Description: fmt.Sprintf("Unable to move file(s) : error[%s].", err),
	}
}

func (c *CmdFileUpload) execUpload(src string, dest *api.DropboxPath) error {
	info, err := os.Lstat(src)
	if err != nil {
		seelog.Warnf("Unable to acquire information about path[%s] : error[%s]", src, err)
		return err
	}

	if info.IsDir() {
		return errors.New("directory upload is not supported")
	} else {
		return c.uploadFile(src, info, dest)
	}
}

func (c *CmdFileUpload) uploadFile(srcPath string, srcInfo os.FileInfo, dest *api.DropboxPath) (err error) {
	//f, err := os.Open(srcPath)
	//if err != nil {
	//	seelog.Warnf("Unable to open file[%s] : error[%s]", srcPath, err)
	//	return err
	//}
	//defer f.Close()
	//ci := files.NewCommitInfo(dest.CleanPath())
	//ci.ClientModified = api.RebaseTimeForAPI(srcInfo.ModTime())
	//
	//_, err = c.apiContext.PatternsFile().Upload(f, srcInfo.Size(), ci)

	return
}
