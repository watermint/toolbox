package copy_ref

import (
	"errors"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_file"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

type Mirror struct {
	ExecContext     *app.ExecContext
	SrcApi          *dbx_api.DbxContext
	SrcAccountAlias string
	SrcAsMemberId   string
	SrcPath         string
	SrcNamespaceId  string
	SrcPathRoot     interface{}
	DstApi          *dbx_api.DbxContext
	DstAccountAlias string
	DstAsMemberId   string
	DstPath         string
	DstNamespaceId  string
	DstPathRoot     interface{}
}

func (z *Mirror) handleError(err error, srcPath, dstPath string) bool {
	z.ExecContext.Msg("dbx_file.copy_ref.mirror.err.failed_mirror").WithData(struct {
		FromPath    string
		FromAccount string
		FromNS      string
		ToPath      string
		ToAccount   string
		ToNS        string
		Error       string
	}{
		FromPath:    srcPath,
		FromAccount: z.SrcAccountAlias,
		FromNS:      z.SrcNamespaceId,
		ToPath:      dstPath,
		ToAccount:   z.DstAccountAlias,
		ToNS:        z.DstNamespaceId,
		Error:       err.Error(),
	}).TellError()

	return true
}

func (z *Mirror) progressFile(file *dbx_file.File, srcPath, dstPath string) bool {
	z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.file.done").WithData(struct {
		FromPath    string
		FromAccount string
		FromNS      string
		ToPath      string
		ToAccount   string
		ToNS        string
	}{
		FromPath:    srcPath,
		FromAccount: z.SrcAccountAlias,
		FromNS:      z.SrcNamespaceId,
		ToPath:      dstPath,
		ToAccount:   z.DstAccountAlias,
		ToNS:        z.DstNamespaceId,
	}).Tell()
	return true
}

func (z *Mirror) progressFolder(folder *dbx_file.Folder, srcPath, dstPath string) bool {
	z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.folder.done").WithData(struct {
		FromPath    string
		FromAccount string
		FromNS      string
		ToPath      string
		ToAccount   string
		ToNS        string
	}{
		FromPath:    srcPath,
		FromAccount: z.SrcAccountAlias,
		FromNS:      z.SrcNamespaceId,
		ToPath:      dstPath,
		ToAccount:   z.DstAccountAlias,
		ToNS:        z.DstNamespaceId,
	}).Tell()
	return true
}

func (z *Mirror) relDstPath(srcPath string) (string, error) {
	pathDiff, err := filepath.Rel(z.SrcPath, srcPath)
	if err != nil {
		z.ExecContext.Log().Debug("unable to calc relative path", zap.String("base", z.SrcPath), zap.String("current", srcPath), zap.Error(err))
		z.ExecContext.Msg("dbx_file.copy_ref.mirror.err.failed_mirror").WithData(struct {
			FromPath    string
			FromAccount string
			FromNS      string
			ToPath      string
			ToAccount   string
			ToNS        string
			Error       string
		}{
			FromPath:    srcPath,
			FromAccount: z.SrcAccountAlias,
			FromNS:      z.SrcNamespaceId,
			ToPath:      z.DstPath,
			ToAccount:   z.DstAccountAlias,
			ToNS:        z.DstNamespaceId,
			Error:       err.Error(),
		}).TellError()
		return "", errors.New("unable to calc relative path")
	}

	// in case of base path
	if pathDiff == "." {
		return z.DstPath, nil
	}

	// should not happen..
	if strings.HasPrefix(pathDiff, "..") {
		err = errors.New("invalid path diff")
		z.ExecContext.Log().Error("invalid path diff", zap.String("diff", pathDiff))
		z.ExecContext.Msg("dbx_file.copy_ref.mirror.err.failed_mirror").WithData(struct {
			FromPath    string
			FromAccount string
			FromNS      string
			ToPath      string
			ToAccount   string
			ToNS        string
			Error       string
		}{
			FromPath:    srcPath,
			FromAccount: z.SrcAccountAlias,
			FromNS:      z.SrcNamespaceId,
			ToPath:      z.DstPath,
			ToAccount:   z.DstAccountAlias,
			ToNS:        z.DstNamespaceId,
			Error:       err.Error(),
		}).TellError()
		return "", err
	}

	curDstPath := filepath.ToSlash(filepath.Join(z.DstPath, pathDiff))

	// preserve case
	curToPathBase := filepath.Base(srcPath)
	curToPathDir := filepath.Dir(curDstPath)
	curDstPath = filepath.ToSlash(filepath.Join(curToPathDir, curToPathBase))

	return curDstPath, nil
}

func (z *Mirror) mirrorDescendants(srcPath, dstPath string) {
	// files in descendants under `dstPath`
	files := make(map[string]*dbx_file.File)
	folders := make(map[string]bool)

	lsDst := dbx_file.ListFolder{
		AsAdminId: z.DstAsMemberId,
		PathRoot:  z.DstPathRoot,

		IncludeMediaInfo:                false,
		IncludeDeleted:                  false,
		IncludeHasExplicitSharedMembers: false,
		IncludeMountedFolders:           true,

		OnError: func(err error) bool {
			switch e := err.(type) {
			case dbx_api.ApiError:
				switch {
				case strings.HasPrefix(e.ErrorSummary, "path/not_found"):
					z.ExecContext.Log().Debug("DST path doesn't have this content", zap.Error(e))
					return false
				}
			}
			z.ExecContext.Log().Debug("other error", zap.Error(err))
			return true
		},
		OnFile: func(file *dbx_file.File) bool {
			files[file.Name] = file
			return true
		},
		OnFolder: func(folder *dbx_file.Folder) bool {
			folders[folder.Name] = true
			return true // ignore result
		},
		OnDelete: func(deleted *dbx_file.Deleted) bool {
			return true // ignore result
		},
	}

	curDstPath, err := z.relDstPath(srcPath)
	if err != nil {
		z.ExecContext.Log().Debug("unable to calc curDstPath", zap.Error(err))
		return
	}

	if !lsDst.List(z.DstApi, curDstPath) {
		z.ExecContext.Log().Debug("List folder returns false")
	}

	skipped := make([]string, 0)
	lsSrc := dbx_file.ListFolder{
		AsMemberId: z.SrcAsMemberId,
		PathRoot:   z.SrcPathRoot,

		IncludeMediaInfo:                false,
		IncludeDeleted:                  false,
		IncludeHasExplicitSharedMembers: false,
		IncludeMountedFolders:           true,

		OnError: func(err error) bool {
			return z.handleError(err, srcPath, dstPath)
		},
		OnFolder: func(folder *dbx_file.Folder) bool {
			if _, e := folders[folder.Name]; e {
				z.ExecContext.Log().Debug("Mirror descendants", zap.String("src", folder.PathDisplay), zap.String("dst", dstPath))
				curToPath, err := z.relDstPath(folder.PathDisplay)
				if err != nil {
					return false
				}
				z.mirrorDescendants(folder.PathDisplay, curToPath)
			} else {
				z.ExecContext.Log().Debug("Mirror folder", zap.String("src", folder.PathDisplay), zap.String("dst", dstPath))
				curToPath, err := z.relDstPath(folder.PathDisplay)
				if err != nil {
					return false
				}
				z.doMirror(folder.PathDisplay, curToPath)
			}
			return true
		},
		OnFile: func(file *dbx_file.File) bool {
			if tf, e := files[file.Name]; e {
				if tf.ContentHash == file.ContentHash {
					//z.ec.Log().Debug("Skip: same content hash", zap.String("srcPath", file.PathDisplay), zap.String("hash", file.ContentHash))
					skipped = append(skipped, file.PathDisplay)
					return true
				}
				// otherwise fallback to mirror
			}
			z.ExecContext.Log().Debug("Mirror descendants file", zap.String("src", file.PathDisplay), zap.String("dst", dstPath))
			curDstPath, err := z.relDstPath(file.PathDisplay)
			if err != nil {
				return false
			}
			z.doMirror(file.PathDisplay, curDstPath)
			return true
		},
		OnDelete: func(deleted *dbx_file.Deleted) bool {
			// log & return
			z.ExecContext.Log().Debug("deleted", zap.Any("deleted", deleted))
			return true
		},
	}
	lsSrc.List(z.SrcApi, srcPath)
	z.ExecContext.Log().Debug("Skipped files; Reason(Same hash)", zap.Strings("files", skipped))
}

func (z *Mirror) handleApiError(ref CopyRef, srcPath, dstPath string, apiErr dbx_api.ApiError) bool {
	z.ExecContext.Log().Debug("handle api error", zap.String("src", srcPath), zap.String("dst", dstPath), zap.String("error_tag", apiErr.ErrorSummary))
	switch {
	case strings.HasPrefix(apiErr.ErrorSummary, "path/conflict"):
		// Mirror each descendants
		z.ExecContext.Log().Debug("conflict found")
		z.mirrorDescendants(srcPath, dstPath)
		return true

	case strings.HasPrefix(apiErr.ErrorSummary, "too_many_files"):
		// Mirror each descendants
		z.ExecContext.Log().Debug("too many files")
		z.mirrorDescendants(srcPath, dstPath)
		return true

	case strings.HasPrefix(apiErr.ErrorSummary, "path/too_many_write_operations"):
		// Retry
		z.ExecContext.Log().Debug("too many write operations, wait & retry")
		return false

	default:
		// log and return
		errMsg := apiErr.ErrorSummary
		if apiErr.UserMessage != "" {
			errMsg = apiErr.UserMessage
		}
		z.ExecContext.Msg("dbx_file.copy_ref.mirror.err.failed_mirror").WithData(struct {
			FromPath    string
			FromAccount string
			ToPath      string
			ToAccount   string
			Error       string
		}{
			FromPath:    srcPath,
			FromAccount: z.SrcAccountAlias,
			ToPath:      dstPath,
			ToAccount:   z.DstAccountAlias,
			Error:       errMsg,
		}).TellError()

		z.ExecContext.Log().Debug("other error_tag", zap.String("error_tag", apiErr.ErrorSummary))
		return false
	}
}

func (z *Mirror) onEntry(ref CopyRef, srcPath, dstPath string) bool {
	crs := CopyRefSave{
		AsMemberId: z.DstAsMemberId,
		PathRoot:   z.DstPathRoot,
		OnError: func(err error) bool {
			//return z.handleError(err, srcPath, dstPath)
			return true
		},
		OnFile: func(file *dbx_file.File) bool {
			return z.progressFile(file, srcPath, dstPath)
		},
		OnFolder: func(folder *dbx_file.Folder) bool {
			return z.progressFolder(folder, srcPath, dstPath)
		},
	}
	z.ExecContext.Log().Debug("Trying to mirror", zap.String("ref", ref.CopyReference), zap.String("src", srcPath), zap.String("dst", dstPath))
	err := crs.Save(z.DstApi, ref, dstPath)
	if err == nil {
		z.ExecContext.Log().Debug("Mirror completed", zap.String("src", srcPath), zap.String("dst", dstPath))
		return true
	}

	switch e := err.(type) {
	case dbx_api.ApiError:
		return z.handleApiError(ref, srcPath, dstPath, e)

	default:
		z.ExecContext.Log().Debug("default error handling", zap.Error(err))
		return false
	}
}

func (z *Mirror) doMirror(srcPath, dstPath string) {
	//z.ec.Msg("dbx_file.copy_ref.mirror.progress.trying").WithData(struct {
	//	SrcPath    string
	//	FromAccount string
	//	DstPath      string
	//	ToAccount   string
	//}{
	//	SrcPath:    srcPath,
	//	FromAccount: z.SrcAccountAlias,
	//	DstPath:      dstPath,
	//	ToAccount:   z.DstAccountAlias,
	//}).Tell()

	crg := CopyRefGet{
		AsMemberId: z.SrcAsMemberId,
		PathRoot:   z.SrcPathRoot,
		OnError: func(err error) bool {
			return z.handleError(err, srcPath, dstPath)
		},
		OnEntry: func(ref CopyRef) bool {
			return z.onEntry(ref, srcPath, dstPath)
		},
	}
	crg.Get(z.SrcApi, srcPath)
}

func (z *Mirror) updatePathRoot() {
	if z.SrcNamespaceId != "" {
		z.SrcPathRoot = dbx_api.NewPathRootNamespace(z.SrcNamespaceId)
	}
	if z.DstNamespaceId != "" {
		z.DstPathRoot = dbx_api.NewPathRootNamespace(z.DstNamespaceId)
	}
}

func (z *Mirror) verifyGivenPaths() bool {
	if z.SrcPath == "/" {
		return true
	}

	mdSrc := dbx_file.Metadata{
		Path:                            z.SrcPath,
		PathRoot:                        z.SrcPathRoot,
		IncludeMediaInfo:                false,
		IncludeDeleted:                  false,
		IncludeHasExplicitSharedMembers: false,
		AsMemberId:                      z.SrcAsMemberId,
		OnError: func(err error) bool {
			z.handleError(err, z.SrcPath, z.DstPath)
			return false
		},
		OnFile: func(file *dbx_file.File) bool {
			z.ExecContext.Log().Debug("Updating src path into DisplayPath", zap.String("srcPath", z.SrcPath), zap.String("displayPath", file.PathDisplay))
			z.SrcPath = file.PathDisplay
			return true
		},
		OnFolder: func(folder *dbx_file.Folder) bool {
			z.ExecContext.Log().Debug("Updating src path into DisplayPath", zap.String("srcPath", z.SrcPath), zap.String("displayPath", folder.PathDisplay))
			z.SrcPath = folder.PathDisplay
			return true
		},
		OnDelete: func(deleted *dbx_file.Deleted) bool {
			// ignore
			return true
		},
	}
	return mdSrc.Get(z.SrcApi)
}

func (z *Mirror) Mirror() {
	z.updatePathRoot()
	if !z.verifyGivenPaths() {
		return
	}

	z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.start").Tell()
	z.doMirror(z.SrcPath, z.DstPath)
	z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.done").Tell()
}

func (z *Mirror) MirrorDescendants() {
	z.updatePathRoot()
	if !z.verifyGivenPaths() {
		return
	}

	z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.start").Tell()
	z.mirrorDescendants(z.SrcPath, z.DstPath)
	z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.done").Tell()
}
