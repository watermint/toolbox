package cmd_file

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/api/patterns"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"path/filepath"
	"strings"
)

type CmdFileMove struct {
	optForce        bool
	optIgnoreErrors bool
	apiContext      *api.ApiContext
	infraContext    *infra.InfraContext
	ParamSrc        *api.DropboxPath
	ParamDest       *api.DropboxPath
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

	descContOnError := "Continue operation even if there are API errors"
	f.BoolVar(&c.optIgnoreErrors, "ignore-errors", false, descContOnError)

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
	c.ParamSrc = api.NewDropboxPath(remainder[0])
	c.ParamDest = api.NewDropboxPath(remainder[1])
	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	seelog.Debugf("move:%s", util.MarshalObjectToString(c))
	c.apiContext, err = c.infraContext.LoadOrAuthDropboxFull()
	if err != nil {
		seelog.Warnf("Unable to acquire token  : error[%s]", err)
		return &cmdlet.CommandError{
			Context:     cc,
			ReasonTag:   "auth/auth_failed",
			Description: fmt.Sprintf("Unable to acquire token : error[%s].", err),
		}
	}

	if err != nil {
		return c.composeError(cc, err)
	}

	srcPaths, err := c.examineSrc(c.ParamSrc)
	if err != nil {
		return c.composeError(cc, err)
	}

	err = c.dispatch(cc, srcPaths, c.ParamDest)
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

func (c *CmdFileMove) examineSrc(srcParam *api.DropboxPath) (src []*api.DropboxPath, err error) {
	gmaSrc := files.NewGetMetadataArg(srcParam.CleanPath())
	seelog.Tracef("examine src path[%s]", gmaSrc.Path)
	metaSrc, err := c.apiContext.FilesGetMetadata(gmaSrc)
	if err != nil {
		return
	}
	switch ms := metaSrc.(type) {
	case *files.FileMetadata:
		seelog.Tracef("src file id[%s] path[%s] size[%d] hash[%s]", ms.Id, ms.PathDisplay, ms.Size, ms.ContentHash)
		src = make([]*api.DropboxPath, 1)
		src[0] = &api.DropboxPath{Path: ms.PathDisplay}
		return

	case *files.FolderMetadata:
		seelog.Tracef("src folder id[%s] path[%s]", ms.Id, ms.PathDisplay)
		if strings.HasSuffix(srcParam.Path, "/") {
			seelog.Tracef("try expand path[%s]", ms.PathDisplay)
			return c.examineSrcExpand(ms)
		} else {
			src = make([]*api.DropboxPath, 1)
			src[0] = &api.DropboxPath{Path: ms.PathDisplay}
			return
		}

	default:
		seelog.Warnf("Unable to move file(s): unexpected metadata found for path[%s]", gmaSrc.Path)
		err = errors.New("unexpected_metadata")
		return
	}
}

func (c *CmdFileMove) examineSrcExpand(srcFolder *files.FolderMetadata) (src []*api.DropboxPath, err error) {
	seelog.Tracef("examine src/expand id[%s] path[%s]", srcFolder.Id, srcFolder.PathDisplay)
	lfa := files.NewListFolderArg(srcFolder.PathDisplay)
	lfa.IncludeMountedFolders = true
	lfa.Recursive = false
	res, err := c.apiContext.FilesListFolder(lfa)
	if err != nil {
		seelog.Warnf("Unable to list folder : error[%s]", err)
		return
	}

	src = make([]*api.DropboxPath, 0)
	for {
		for _, e := range res.Entries {
			switch f := e.(type) {
			case *files.FileMetadata:
				seelog.Tracef("src file: id[%s] path[%s] size[%d] hash[%s]", f.Id, f.PathDisplay, f.Size, f.ContentHash)
				src = append(src, api.NewDropboxPath(f.PathDisplay))

			case *files.FolderMetadata:
				seelog.Tracef("src folder: id[%s] path[%s]", f.Id, f.PathDisplay)
				src = append(src, api.NewDropboxPath(f.PathDisplay))

			default:
				seelog.Debugf("unexpected metadata found at path[%s] meta[%s]")
			}
		}
		if !res.HasMore {
			return
		}
		seelog.Tracef("continue list folder with cursor[%s]", res.Cursor)
		res, err = c.apiContext.FilesListFolderContinue(files.NewListFolderContinueArg(res.Cursor))
		if err != nil {
			seelog.Warnf("Unable to list folder : error[%s]", err)
			return
		}
	}
}

func (c *CmdFileMove) dispatch(cc cmdlet.CommandletContext, srcPaths []*api.DropboxPath, dest *api.DropboxPath) (err error) {
	gmaDest := files.NewGetMetadataArg(c.ParamDest.CleanPath())
	seelog.Tracef("examine dest path[%s]", gmaDest.Path)
	metaDest, err := c.apiContext.FilesGetMetadata(gmaDest)
	if err != nil && strings.HasPrefix(err.Error(), "path/not_found") {
		for _, s := range srcPaths {
			err = c.moveWithRecovery(cc, s, s, dest)
			if err != nil {
				if c.optIgnoreErrors {
					seelog.Warnf("Skip moving file/folder from [%s] to [%s] due to error [%s]", s.Path, dest.Path, err)
				} else {
					return err
				}
			}
		}
		return
	}

	switch md := metaDest.(type) {
	case *files.FileMetadata:
		seelog.Warnf("File exist on destination path [%s]", md.PathDisplay)
		return errors.New("path_conflict")

	case *files.FolderMetadata:
		for _, s := range srcPaths {
			fn := filepath.Base(s.Path)
			d := api.NewDropboxPath(filepath.Join(dest.CleanPath(), fn))
			seelog.Debugf("moving file/folder from [%s] to [%s]", s.Path, d.Path)
			err = c.moveWithRecovery(cc, s, s, d)
			if err != nil {
				if c.optIgnoreErrors {
					seelog.Warnf("Skip moving file/folder from [%s] to [%s] due to error [%s]", s.Path, dest.Path, err)
				} else {
					return err
				}
			}
		}
	}
	return nil
}

func (c *CmdFileMove) destPath(path *api.DropboxPath, srcBase *api.DropboxPath, destBase *api.DropboxPath) (dest *api.DropboxPath, err error) {
	rel, err := filepath.Rel(srcBase.CleanPath(), path.CleanPath())
	if err != nil {
		seelog.Warnf("Unable to compute destination path[%s] : error[%s]", path.CleanPath(), err)
		return
	}
	dest = api.NewDropboxPath(filepath.ToSlash(filepath.Join(destBase.CleanPath(), rel)))
	return
}

func (c *CmdFileMove) move(cc cmdlet.CommandletContext, src *api.DropboxPath, srcBase *api.DropboxPath, destBase *api.DropboxPath) (err error) {
	seelog.Debugf("Move: srcBase[%s] src[%s] dest[%s]", srcBase.Path, src.Path, destBase.Path)
	dest, err := c.destPath(src, srcBase, destBase)
	if err != nil {
		return err
	}
	arg := files.NewRelocationArg(src.CleanPath(), dest.CleanPath())
	arg.Autorename = false
	arg.AllowSharedFolder = true
	arg.AllowOwnershipTransfer = true

	seelog.Debugf("Move from[%s] to[%s]", arg.FromPath, arg.ToPath)

	_, err = c.apiContext.FilesMoveV2(arg)
	return
}

func (c *CmdFileMove) moveWithRecovery(cc cmdlet.CommandletContext, src *api.DropboxPath, srcBase *api.DropboxPath, dest *api.DropboxPath) (err error) {
	seelog.Debugf("Move with recoverable opt: srcBase[%s] src[%s] dest[%s]", srcBase.Path, src.Path, dest.Path)

	err = c.move(cc, src, srcBase, dest)
	if !c.isRecoverableError(err) {
		return err
	}
	if err != nil {
		seelog.Debugf("Recoverable error[%s] src[%s] dest[%s]", err, src.Path, dest.Path)
	}

	tw := patterns.NewTreeWalk(c.apiContext)
	tw.RecursiveWalk = false
	tw.Recursive = false
	tw.IncludeDeleted = false
	tw.IncludeMediaInfo = false
	tw.IncludeHasExplicitSharedMembers = false
	tw.IncludeMountedFolders = true

	err = tw.Walk(src, func(entry files.IsMetadata) error {
		switch f := entry.(type) {
		case *files.FileMetadata:
			return c.move(cc, api.NewDropboxPath(f.PathDisplay), srcBase, dest)

		case *files.FolderMetadata:
			return c.moveWithRecovery(cc, api.NewDropboxPath(f.PathDisplay), srcBase, dest)

		default:
			// ignore
			return nil
		}
	})
	return
}

func (c *CmdFileMove) isRecoverableError(err error) bool {
	if err == nil {
		return true
	}
	if strings.HasPrefix(err.Error(), "too_many_files") {
		return true
	}
	if c.optForce && strings.HasPrefix(err.Error(), "cant_copy_shared_folder") {
		return true
	}
	if c.optForce && strings.HasPrefix(err.Error(), "cant_nest_shared_folder") {
		return true
	}
	return false
}
