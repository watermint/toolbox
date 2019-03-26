package uc_file_mirror

import (
	"errors"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_file_copyref"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
	"time"
)

type Files interface {
	Mirror(pathSrc, pathDst mo_path.Path) (err error)
}

func New(ctxSrc, ctxDst api_context.Context) Files {
	return &filesImpl{
		ctxSrc:       ctxSrc,
		ctxDst:       ctxDst,
		pollInterval: 15 * time.Second,
	}
}

type filesImpl struct {
	ctxSrc       api_context.Context
	ctxDst       api_context.Context
	pollInterval time.Duration
}

func (z *filesImpl) report(metaSrc, metaDst mo_file.Entry) {
	z.ctxSrc.Log().Debug("Mirror complete", zap.String("src", metaSrc.PathDisplay()), zap.String("dst", metaDst.PathDisplay()))
}

func (z *filesImpl) dstPathRelToSrc(pathOrigSrc, pathSrc, pathOrigDst mo_path.Path) (relDst mo_path.Path, err error) {
	log := z.ctxSrc.Log().With(zap.String("origSrc", pathOrigSrc.Path()), zap.String("src", pathSrc.Path()), zap.String("origDst", pathOrigDst.Path()))

	pathDiff, err := filepath.Rel(pathOrigSrc.LogicalPath(), pathSrc.LogicalPath())
	if err != nil {
		log.Warn("Unable to calc rel path", zap.Error(err))
		return nil, err
	}
	pathDiff = filepath.ToSlash(pathDiff)

	if pathDiff == "." {
		return pathOrigDst, nil
	}

	if strings.HasPrefix(pathDiff, "..") {
		log.Warn("Invalid path", zap.String("diff", pathDiff))
		return nil, errors.New("invalid path")
	}

	relDst = pathOrigDst.ChildPath(pathDiff)

	log.Debug("DST: Path rel to src", zap.String("path", relDst.Path()))

	return relDst, nil
}

func (z *filesImpl) mirrorDescendants(pathOrigSrc, pathSrc, pathOrigDst, pathDst mo_path.Path) error {
	log := z.ctxSrc.Log().With(zap.String("origSrc", pathOrigSrc.Path()), zap.String("src", pathSrc.Path()), zap.String("dst", pathDst.Path()))

	// dest descendant files, and folders
	filesDst := make(map[string]*mo_file.File)
	foldersDst := make(map[string]bool)

	// scan dest path
	pathDstRelToSrc, err := z.dstPathRelToSrc(pathOrigSrc, pathSrc, pathOrigDst)
	if err != nil {
		return err
	}
	svfDst := sv_file.NewFiles(z.ctxDst)
	entriesDst, err := svfDst.List(pathDstRelToSrc)
	if err != nil {
		errPrefix := api_util.ErrorSummary(err)
		if !strings.HasPrefix(errPrefix, "path/not_found") {
			log.Debug("DST: Unable to list", zap.Error(err), zap.String("error_summary", errPrefix))
			return err
		}
		log.Debug("DST: Path not found. Proceed to mirror")
	} else {
		for _, entry := range entriesDst {
			name := strings.ToLower(entry.Name())
			if f, e := entry.File(); e {
				filesDst[name] = f
			}
			if _, e := entry.Folder(); e {
				foldersDst[name] = true
			}
		}
	}

	// skipped files (same hash)
	skipped := make([]string, 0)

	// scan src path
	svfSrc := sv_file.NewFiles(z.ctxSrc)
	entriesSrc, err := svfSrc.List(pathSrc)
	if err != nil {
		return err
	}

	for _, entrySrc := range entriesSrc {
		name := strings.ToLower(entrySrc.Name())
		if fileSrc, e := entrySrc.File(); e {
			if fileDst, e := filesDst[name]; e {
				if fileSrc.ContentHash == fileDst.ContentHash {
					skipped = append(skipped, fileSrc.PathDisplay())
					continue
				}
				// otherwise fall back to mirror
			}
			pathFileSrc := mo_path.NewPathDisplay(fileSrc.PathDisplay())
			pathFileDst, err := z.dstPathRelToSrc(pathOrigSrc, pathFileSrc, pathOrigDst)
			if err != nil {
				return err
			}
			if err = z.mirrorCurrent(pathOrigSrc, pathFileSrc, pathOrigDst, pathFileDst); err != nil {
				log.Debug("Mirror failed", zap.String("fileSrc", pathFileSrc.Path()), zap.String("fileDst", pathFileDst.Path()), zap.Error(err))
				// do not fail on file mirroring
			}
		}

		if folderSrc, e := entrySrc.Folder(); e {
			if _, e := foldersDst[name]; e {
				pathFolderSrc := mo_path.NewPathDisplay(folderSrc.PathDisplay())
				pathFolderDst, err := z.dstPathRelToSrc(pathOrigSrc, pathFolderSrc, pathOrigDst)
				if err != nil {
					return err
				}
				if err = z.mirrorCurrent(pathOrigSrc, pathFolderSrc, pathOrigDst, pathFolderDst); err != nil {
					log.Debug("Mirror failed", zap.String("folderSrc", pathFolderSrc.Path()), zap.String("folderDst", pathFolderDst.Path()), zap.Error(err))
					// do not fail on file mirroring
				}
			}
		}
	}

	return nil
}

func (z *filesImpl) mirrorCurrent(pathOrigSrc, pathSrc, pathOrigDst, pathDst mo_path.Path) (err error) {
	log := z.ctxSrc.Log().With(zap.String("origSrc", pathOrigSrc.Path()), zap.String("src", pathSrc.Path()), zap.String("dst", pathDst.Path()))
	scrSrc := sv_file_copyref.New(z.ctxSrc)
	metaSrc, ref, expires, err := scrSrc.Resolve(pathSrc)
	if err != nil {
		log.Debug("SRC: Unable to get copyRef", zap.Error(err))
		return err
	}
	log.Debug("SRC: CopyRef success", zap.String("ref", ref), zap.String("expires", expires))

	scrDst := sv_file_copyref.New(z.ctxDst)
	metaDst, err := scrDst.Save(pathDst, ref)
	if err != nil {
		log.Debug("DST: Unable to save", zap.Error(err))
		return z.handleError(pathOrigSrc, pathSrc, pathOrigDst, pathDst, err)
	}

	z.report(metaSrc, metaDst)
	return nil
}

func (z *filesImpl) handleError(pathOrigSrc, pathSrc, pathOrigDst, pathDst mo_path.Path, apiErr error) error {
	errPrefix := api_util.ErrorSummary(apiErr)
	log := z.ctxSrc.Log().With(zap.String("origSrc", pathOrigSrc.Path()), zap.String("src", pathSrc.Path()), zap.String("dst", pathDst.Path()), zap.String("errorPrefix", errPrefix))
	switch {
	case strings.HasPrefix(errPrefix, "path/conflict"),
		strings.HasPrefix(errPrefix, "too_many_files"):
		log.Debug("Mirror descendants")
		return z.mirrorDescendants(pathOrigSrc, pathSrc, pathOrigDst, pathDst)

	case strings.Contains(errPrefix, "too_many_write_operations"):
		log.Debug("Wait for too many write", zap.Duration("wait", z.pollInterval))
		time.Sleep(z.pollInterval)
		return z.mirrorCurrent(pathOrigSrc, pathSrc, pathOrigDst, pathDst)

	default:
		log.Debug("Unrecoverable error")
		return apiErr
	}
}

func (z *filesImpl) Mirror(pathSrc, pathDst mo_path.Path) (err error) {
	return z.mirrorCurrent(pathSrc, pathSrc, pathDst, pathDst)
}
