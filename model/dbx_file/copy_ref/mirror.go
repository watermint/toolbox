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
	ExecContext      *app.ExecContext
	FromApi          *dbx_api.Context
	FromAccountAlias string
	FromAsMemberId   string
	FromPath         string
	ToApi            *dbx_api.Context
	ToAccountAlias   string
	ToAsMemberId     string
	ToPath           string
}

func (z *Mirror) handleError(annotation dbx_api.ErrorAnnotation, fromPath, toPath string) bool {
	z.ExecContext.Msg("dbx_file.copy_ref.mirror.err.failed_mirror").WithData(struct {
		FromPath    string
		FromAccount string
		ToPath      string
		ToAccount   string
		Error       string
	}{
		FromPath:    fromPath,
		FromAccount: z.FromAccountAlias,
		ToPath:      toPath,
		ToAccount:   z.ToAccountAlias,
		Error:       annotation.Error.Error(),
	}).TellError()

	return true
}

func (z *Mirror) progressFile(file *dbx_file.File, fromPath, toPath string) bool {
	z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.file.done").WithData(struct {
		FromPath    string
		FromAccount string
		ToPath      string
		ToAccount   string
	}{
		FromPath:    fromPath,
		FromAccount: z.FromAccountAlias,
		ToPath:      toPath,
		ToAccount:   z.ToAccountAlias,
	}).Tell()
	return true
}

func (z *Mirror) progressFolder(folder *dbx_file.Folder, fromPath, toPath string) bool {
	z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.folder.done").WithData(struct {
		FromPath    string
		FromAccount string
		ToPath      string
		ToAccount   string
	}{
		FromPath:    fromPath,
		FromAccount: z.FromAccountAlias,
		ToPath:      toPath,
		ToAccount:   z.ToAccountAlias,
	}).Tell()
	return true
}

func (z *Mirror) destToPath(fromPath string) (string, error) {
	pathDiff, err := filepath.Rel(strings.ToLower(z.FromPath), strings.ToLower(fromPath))
	if err != nil {
		z.ExecContext.Log().Debug("unable to calc relative path", zap.String("base", z.FromPath), zap.String("current", fromPath), zap.Error(err))
		z.ExecContext.Msg("dbx_file.copy_ref.mirror.err.failed_mirror").WithData(struct {
			FromPath    string
			FromAccount string
			ToPath      string
			ToAccount   string
			Error       string
		}{
			FromPath:    fromPath,
			FromAccount: z.FromAccountAlias,
			ToPath:      z.ToPath,
			ToAccount:   z.ToAccountAlias,
			Error:       err.Error(),
		}).TellError()
		return "", errors.New("unable to calc relative path")
	}

	// in case of base path
	if pathDiff == "." {
		return z.ToPath, nil
	}

	// should not happen..
	if strings.HasPrefix(pathDiff, "..") {
		err = errors.New("invalid path diff")
		z.ExecContext.Log().Error("invalid path diff", zap.String("diff", pathDiff))
		z.ExecContext.Msg("dbx_file.copy_ref.mirror.err.failed_mirror").WithData(struct {
			FromPath    string
			FromAccount string
			ToPath      string
			ToAccount   string
			Error       string
		}{
			FromPath:    fromPath,
			FromAccount: z.FromAccountAlias,
			ToPath:      z.ToPath,
			ToAccount:   z.ToAccountAlias,
			Error:       err.Error(),
		}).TellError()
		return "", err
	}

	curToPath := filepath.ToSlash(filepath.Join(z.ToPath, pathDiff))

	// preserve case
	curToPathBase := filepath.Base(fromPath)
	curToPathDir := filepath.Dir(curToPath)
	curToPath = filepath.ToSlash(filepath.Join(curToPathDir, curToPathBase))

	z.ExecContext.Log().Debug("list `current` toPath", zap.String("curToPath", curToPath), zap.String("pathDiff", pathDiff))

	return curToPath, nil
}

func (z *Mirror) mirrorAncestors(fromPath, toPath string) {
	// files in ancestor under `toPath`
	files := make(map[string]*dbx_file.File)
	folders := make(map[string]bool)

	lst := dbx_file.ListFolder{
		AsMemberId: z.ToAsMemberId,

		IncludeMediaInfo:                false,
		IncludeDeleted:                  false,
		IncludeHasExplicitSharedMembers: false,
		IncludeMountedFolders:           true,

		OnError: func(annotation dbx_api.ErrorAnnotation) bool {
			switch e := annotation.Error.(type) {
			case dbx_api.ApiError:
				switch {
				case strings.HasPrefix(e.ErrorSummary, "path/not_found"):
					z.ExecContext.Log().Debug("To path doesn't have this content", zap.Error(e))
					return false
				}
			}
			z.ExecContext.Log().Debug("other error", zap.Error(annotation.Error))
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

	curToPath, err := z.destToPath(fromPath)
	if err != nil {
		return
	}

	if !lst.List(z.ToApi, curToPath) {
		z.ExecContext.Log().Debug("List folder returns false")
		return
	}

	lsf := dbx_file.ListFolder{
		AsMemberId: z.FromAsMemberId,

		IncludeMediaInfo:                false,
		IncludeDeleted:                  false,
		IncludeHasExplicitSharedMembers: false,
		IncludeMountedFolders:           true,

		OnError: func(annotation dbx_api.ErrorAnnotation) bool {
			return z.handleError(annotation, fromPath, toPath)
		},
		OnFolder: func(folder *dbx_file.Folder) bool {
			if _, e := folders[folder.Name]; e {
				z.ExecContext.Log().Debug("Copy ancestors", zap.String("from", folder.PathDisplay), zap.String("to", toPath))
				curToPath, err := z.destToPath(folder.PathDisplay)
				if err != nil {
					return false
				}
				z.mirrorAncestors(folder.PathDisplay, curToPath)
			} else {
				z.ExecContext.Log().Debug("Copy folder", zap.String("from", folder.PathDisplay), zap.String("to", toPath))
				curToPath, err := z.destToPath(folder.PathDisplay)
				if err != nil {
					return false
				}
				z.doMirror(folder.PathDisplay, curToPath)
			}
			return true
		},
		OnFile: func(file *dbx_file.File) bool {
			if tf, e := files[file.Name]; e {
				z.ExecContext.Log().Debug("File exists on toSide", zap.String("fromPath", file.PathDisplay), zap.String("toPath", tf.PathDisplay))
				if tf.ContentHash == file.ContentHash {
					z.ExecContext.Log().Debug("Skip: same content hash", zap.String("fromPath", file.PathDisplay), zap.String("hash", file.ContentHash))
					return true
				}
				// otherwise fallback to mirror
			}
			z.ExecContext.Log().Debug("Copy ancestor file", zap.String("from", file.PathDisplay), zap.String("to", toPath))
			curToPath, err := z.destToPath(file.PathDisplay)
			if err != nil {
				return false
			}
			z.doMirror(file.PathDisplay, curToPath)
			return true
		},
		OnDelete: func(deleted *dbx_file.Deleted) bool {
			// log & return
			z.ExecContext.Log().Debug("deleted", zap.Any("deleted", deleted))
			return true
		},
	}
	lsf.List(z.FromApi, fromPath)

}

func (z *Mirror) handleApiError(ref CopyRef, fromPath, toPath string, apiErr dbx_api.ApiError) bool {
	z.ExecContext.Log().Debug("handle api error", zap.String("from", fromPath), zap.String("to", toPath), zap.String("error_tag", apiErr.ErrorSummary))
	switch {
	case strings.HasPrefix(apiErr.ErrorSummary, "path/conflict"):
		// Copy each ancestors
		z.ExecContext.Log().Debug("conflict found")
		z.mirrorAncestors(fromPath, toPath)
		return true

	case strings.HasPrefix(apiErr.ErrorSummary, "too_many_files"):
		// Copy each ancestors
		z.ExecContext.Log().Debug("too many files")
		z.mirrorAncestors(fromPath, toPath)
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
			FromPath:    fromPath,
			FromAccount: z.FromAccountAlias,
			ToPath:      toPath,
			ToAccount:   z.ToAccountAlias,
			Error:       errMsg,
		}).TellError()

		z.ExecContext.Log().Debug("other error_tag", zap.String("error_tag", apiErr.ErrorSummary))
		return false
	}
}

func (z *Mirror) onEntry(ref CopyRef, fromPath, toPath string) bool {
	crs := CopyRefSave{
		AsMemberId: z.ToAsMemberId,
		OnError: func(annotation dbx_api.ErrorAnnotation) bool {
			return z.handleError(annotation, fromPath, toPath)
		},
		OnFile: func(file *dbx_file.File) bool {
			return z.progressFile(file, fromPath, toPath)
		},
		OnFolder: func(folder *dbx_file.Folder) bool {
			return z.progressFolder(folder, fromPath, toPath)
		},
	}
	z.ExecContext.Log().Debug("Trying to mirror", zap.String("ref", ref.CopyReference), zap.String("from", fromPath), zap.String("toPath", toPath))
	err := crs.Save(z.ToApi, ref, toPath)
	if err == nil {
		z.ExecContext.Log().Debug("Mirror completed", zap.String("from", fromPath), zap.String("to", toPath))
		return true
	}

	switch e := err.(type) {
	case dbx_api.ApiError:
		return z.handleApiError(ref, fromPath, toPath, e)

	default:
		z.ExecContext.Log().Debug("default error handling", zap.Error(err))
		return false
	}
}

func (z *Mirror) doMirror(fromPath, toPath string) {
	//z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.trying").WithData(struct {
	//	FromPath    string
	//	FromAccount string
	//	ToPath      string
	//	ToAccount   string
	//}{
	//	FromPath:    fromPath,
	//	FromAccount: z.FromAccountAlias,
	//	ToPath:      toPath,
	//	ToAccount:   z.ToAccountAlias,
	//}).Tell()

	crg := CopyRefGet{
		AsMemberId: z.FromAsMemberId,
		OnError: func(annotation dbx_api.ErrorAnnotation) bool {
			return z.handleError(annotation, fromPath, toPath)
		},
		OnEntry: func(ref CopyRef) bool {
			return z.onEntry(ref, fromPath, toPath)
		},
	}
	crg.Get(z.FromApi, fromPath)
}

func (z *Mirror) Mirror() {
	z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.start").Tell()
	z.doMirror(z.FromPath, z.ToPath)
	z.ExecContext.Msg("dbx_file.copy_ref.mirror.progress.done").Tell()
}
