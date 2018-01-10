package cmd_file

import (
	"errors"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/api/patterns"
	"path/filepath"
	"strings"
)

type CmdRelocation struct {
	OptForce        bool
	OptIgnoreErrors bool
	ApiContext      *api.ApiContext

	RelocationFunc func(arg *files.RelocationArg) error
}

func (c *CmdRelocation) examineSrc(srcParam *api.DropboxPath) (src []*api.DropboxPath, err error) {
	gmaSrc := files.NewGetMetadataArg(srcParam.CleanPath())
	seelog.Tracef("examine src path[%s]", gmaSrc.Path)
	metaSrc, err := c.ApiContext.Files().GetMetadata(gmaSrc)
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
		seelog.Warnf("Unable to relocation file(s): unexpected metadata found for path[%s]", gmaSrc.Path)
		err = errors.New("unexpected_metadata")
		return
	}
}

func (c *CmdRelocation) examineSrcExpand(srcFolder *files.FolderMetadata) (src []*api.DropboxPath, err error) {
	seelog.Tracef("examine src/expand id[%s] path[%s]", srcFolder.Id, srcFolder.PathDisplay)
	lfa := files.NewListFolderArg(srcFolder.PathDisplay)
	lfa.IncludeMountedFolders = true
	lfa.Recursive = false
	entries, err := c.ApiContext.PatternsFile().ListFolder(lfa)
	if err != nil {
		seelog.Warnf("Unable to list folder : error[%s]", err)
		return
	}

	src = make([]*api.DropboxPath, 0)
	for _, e := range entries {
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
	return
}

func (c *CmdRelocation) Dispatch(srcPath *api.DropboxPath, destPath *api.DropboxPath) (err error) {
	srcPaths, err := c.examineSrc(srcPath)
	if err != nil {
		seelog.Debugf("Unable to examine src path[%s] : error[%s]", srcPath.Path, err)
		return err
	}
	gmaDest := files.NewGetMetadataArg(destPath.CleanPath())
	seelog.Tracef("examine dest path[%s]", gmaDest.Path)
	metaDest, err := c.ApiContext.Files().GetMetadata(gmaDest)
	if err != nil && strings.HasPrefix(err.Error(), "path/not_found") {
		for _, s := range srcPaths {
			err = c.relocationWithRecovery(s, s, destPath)
			if err != nil {
				if c.OptIgnoreErrors {
					seelog.Warnf("Skip moving file/folder from [%s] to [%s] due to error [%s]", s.Path, destPath.Path, err)
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
			d := api.NewDropboxPath(filepath.Join(destPath.CleanPath(), fn))

			seelog.Debugf("moving file/folder from [%s] to [%s]", s.Path, d.Path)
			err = c.relocationWithRecovery(s, s, d)
			if err != nil {
				if c.OptIgnoreErrors {
					seelog.Warnf("Skip moving file/folder from [%s] to [%s] due to error [%s]", s.Path, destPath.Path, err)
				} else {
					return err
				}
			}
		}
	}
	return nil
}

func (c *CmdRelocation) destPath(path *api.DropboxPath, srcBase *api.DropboxPath, destBase *api.DropboxPath) (dest *api.DropboxPath, err error) {
	rel, err := filepath.Rel(srcBase.CleanPath(), path.CleanPath())
	if err != nil {
		seelog.Warnf("Unable to compute destination path[%s] : error[%s]", path.CleanPath(), err)
		return
	}
	dest = api.NewDropboxPath(filepath.ToSlash(filepath.Join(destBase.CleanPath(), rel)))
	return
}

func (c *CmdRelocation) execRelocation(src *api.DropboxPath, srcBase *api.DropboxPath, destBase *api.DropboxPath) (err error) {
	seelog.Tracef("Relocation: srcBase[%s] src[%s] dest[%s]", srcBase.Path, src.Path, destBase.Path)
	dest, err := c.destPath(src, srcBase, destBase)
	if err != nil {
		return err
	}
	arg := files.NewRelocationArg(src.CleanPath(), dest.CleanPath())

	seelog.Tracef("Relocation from[%s] to[%s]", arg.FromPath, arg.ToPath)

	err = c.RelocationFunc(arg)
	return
}

func (c *CmdRelocation) relocationWithRecovery(src *api.DropboxPath, srcBase *api.DropboxPath, dest *api.DropboxPath) (err error) {
	seelog.Tracef("Relocation with recoverable opt: srcBase[%s] src[%s] dest[%s]", srcBase.Path, src.Path, dest.Path)

	err = c.execRelocation(src, srcBase, dest)
	if err == nil {
		seelog.Tracef("Relocation success src[%s] -> dest[%s]", src.Path, dest.Path)
		return
	}
	if strings.HasPrefix(err.Error(), "to/conflict/folder") {
		seelog.Tracef("Ignore conflict error[%s]: src[%s] dest[%s]", err, src.Path, dest.Path)

	} else if !c.isRecoverableError(err) {
		seelog.Debugf("Unrecoverable error found[%s] src[%s] dest[%s]", err, src.Path, dest.Path)
		return err
	}

	seelog.Debugf("Recoverable error[%s] src[%s] dest[%s]", err, src.Path, dest.Path)

	lfa := files.NewListFolderArg(src.CleanPath())
	lfa.Recursive = false
	lfa.IncludeDeleted = false
	lfa.IncludeMediaInfo = false
	lfa.IncludeHasExplicitSharedMembers = false
	lfa.IncludeMountedFolders = true

	entries, err := c.ApiContext.PatternsFile().ListFolder(lfa)
	if err != nil {
		seelog.Warnf("Unable to list files, path[%s] : error[%s]", lfa.Path, err)
		return err
	}
	for _, entry := range entries {
		err = nil
		path := ""
		isFolder := false
		switch f := entry.(type) {
		case *files.FileMetadata:
			err = c.execRelocation(api.NewDropboxPath(f.PathDisplay), srcBase, dest)
			path = f.PathDisplay

		case *files.FolderMetadata:
			err = c.relocationWithRecovery(api.NewDropboxPath(f.PathDisplay), srcBase, dest)
			path = f.PathDisplay
			isFolder = true

		default:
			// ignore
		}

		if err != nil {
			if isFolder && strings.HasSuffix(err.Error(), "to/conflict/folder") {
				seelog.Debugf("Conflict folder found[%s]", path)
			} else {
				seelog.Warnf("Unable to relocation path[%s] due to error [%s]", path, err)
				if !c.OptIgnoreErrors {
					return err
				}
			}
		}
	}
	return
}

func (c *CmdRelocation) isRecoverableError(err error) bool {
	if err == nil {
		return true
	}
	if strings.HasPrefix(err.Error(), "too_many_files") {
		return true
	}
	if c.OptForce && strings.HasPrefix(err.Error(), "cant_copy_shared_folder") {
		return true
	}
	if c.OptForce && strings.HasPrefix(err.Error(), "cant_nest_shared_folder") {
		return true
	}
	return false
}
