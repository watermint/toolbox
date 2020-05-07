package uc_file_mirror

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_copyref"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"path/filepath"
	"strings"
	"time"
)

type Files interface {
	Mirror(pathSrc, pathDst mo_path.DropboxPath) (err error)
}

func New(ctxSrc, ctxDst dbx_context.Context) Files {
	return &filesImpl{
		ctxSrc:       ctxSrc,
		ctxDst:       ctxDst,
		pollInterval: 15 * 1000 * time.Millisecond,
	}
}

type filesImpl struct {
	ctxSrc       dbx_context.Context
	ctxDst       dbx_context.Context
	pollInterval time.Duration
}

func (z *filesImpl) report(metaSrc, metaDst mo_file.Entry) {
	z.ctxSrc.Log().Debug("Mirror complete", es_log.String("src", metaSrc.PathDisplay()), es_log.String("dst", metaDst.PathDisplay()))
}

func (z *filesImpl) dstPathRelToSrc(pathOrigSrc, pathSrc, pathOrigDst mo_path.DropboxPath) (relDst mo_path.DropboxPath, err error) {
	log := z.ctxSrc.Log().With(es_log.String("origSrc", pathOrigSrc.Path()), es_log.String("src", pathSrc.Path()), es_log.String("origDst", pathOrigDst.Path()))

	pathDiff, err := filepath.Rel(pathOrigSrc.LogicalPath(), pathSrc.LogicalPath())
	if err != nil {
		log.Warn("Unable to calc rel path", es_log.Error(err))
		return nil, err
	}
	pathDiff = filepath.ToSlash(pathDiff)

	if pathDiff == "." {
		return pathOrigDst, nil
	}

	if strings.HasPrefix(pathDiff, "..") {
		log.Warn("Invalid path", es_log.String("diff", pathDiff))
		return nil, errors.New("invalid path")
	}

	relDst = pathOrigDst.ChildPath(pathDiff)

	log.Debug("DST: Path rel to src", es_log.String("path", relDst.Path()))

	return relDst, nil
}

func (z *filesImpl) mirrorDescendants(pathOrigSrc, pathSrc, pathOrigDst, pathDst mo_path.DropboxPath) error {
	log := z.ctxSrc.Log().With(es_log.String("origSrc", pathOrigSrc.Path()), es_log.String("src", pathSrc.Path()), es_log.String("dst", pathDst.Path()))
	log.Debug("Start mirroring descendants")

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
		ers := dbx_error.NewErrors(err)
		if !ers.Path().IsNotFound() {
			log.Debug("DST: Unable to list", es_log.Error(err), es_log.String("error_summary", ers.Summary()))
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
	type Skip struct {
		Name   string `json:"name"`
		Reason string `json:"reason"`
	}
	skipped := make([]*Skip, 0)

	// scan src path
	svfSrc := sv_file.NewFiles(z.ctxSrc)
	entriesSrc, err := svfSrc.List(pathSrc)
	if err != nil {
		return err
	}

	numEntriesSrc := len(entriesSrc)
	for i, entrySrc := range entriesSrc {
		log.Debug("Processing entry", es_log.Int("index", i), es_log.Int("numEntries", numEntriesSrc), es_log.String("entryName", entrySrc.Name()))
		name := strings.ToLower(entrySrc.Name())
		if fileSrc, e := entrySrc.File(); e {
			if fileDst, e := filesDst[name]; e {
				// Skip when same hash
				if fileSrc.ContentHash == fileDst.ContentHash {
					skipped = append(skipped, &Skip{
						Name:   fileSrc.Name(),
						Reason: "same_hash",
					})
					continue
				}
				srcTime, err := dbx_util.Parse(fileSrc.ServerModified)
				if err != nil {
					log.Warn("Unable to determine fileSrc server modified", es_log.Any("fileSrc", fileSrc), es_log.Error(err))
					skipped = append(skipped, &Skip{
						Name:   fileSrc.Name(),
						Reason: "src_time_err",
					})
					continue
				}
				dstTime, err := dbx_util.Parse(fileDst.ServerModified)
				if err != nil {
					log.Warn("Unable to determine fileDst server modified", es_log.Any("fileDst", fileDst), es_log.Error(err))
					skipped = append(skipped, &Skip{
						Name:   fileSrc.Name(),
						Reason: "dst_time_err",
					})
					continue
				}
				if dstTime.After(srcTime) {
					skipped = append(skipped, &Skip{
						Name:   fileSrc.Name(),
						Reason: "timestamp",
					})
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
				log.Debug("Mirror failed", es_log.String("fileSrc", pathFileSrc.Path()), es_log.String("fileDst", pathFileDst.Path()), es_log.Error(err))
				// do not fail on file mirroring
			}
		}

		if folderSrc, e := entrySrc.Folder(); e {
			pathFolderSrc := mo_path.NewPathDisplay(folderSrc.PathDisplay())
			pathFolderDst, err := z.dstPathRelToSrc(pathOrigSrc, pathFolderSrc, pathOrigDst)
			if err != nil {
				return err
			}
			if _, e := foldersDst[name]; e {
				if err = z.mirrorDescendants(pathOrigSrc, pathFolderSrc, pathOrigDst, pathFolderDst); err != nil {
					log.Debug("Mirror failed", es_log.String("folderSrc", pathFolderSrc.Path()), es_log.String("folderDst", pathFolderDst.Path()), es_log.Error(err))
					// do not fail on file mirroring
				}
			} else {
				if err = z.mirrorCurrent(pathOrigSrc, pathFolderSrc, pathOrigDst, pathFolderDst); err != nil {
					log.Debug("Mirror failed", es_log.String("folderSrc", pathFolderSrc.Path()), es_log.String("folderDst", pathFolderDst.Path()), es_log.Error(err))
					// do not fail on file mirroring
				}
			}
		}
	}

	log.Debug("Skipped:", es_log.Any("files", skipped))

	return nil
}

func (z *filesImpl) mirrorCurrent(pathOrigSrc, pathSrc, pathOrigDst, pathDst mo_path.DropboxPath) (err error) {
	log := z.ctxSrc.Log().With(es_log.String("origSrc", pathOrigSrc.Path()), es_log.String("src", pathSrc.Path()), es_log.String("dst", pathDst.Path()))
	log.Debug("Start mirroring current path")

	scrSrc := sv_file_copyref.New(z.ctxSrc)
	metaSrc, ref, expires, err := scrSrc.Resolve(pathSrc)
	if err != nil {
		log.Debug("SRC: Unable to get copyRef", es_log.Error(err))
		return err
	}
	log.Debug("SRC: CopyRef success", es_log.String("ref", ref), es_log.String("expires", expires))

	scrDst := sv_file_copyref.New(z.ctxDst)
	metaDst, err := scrDst.Save(pathDst, ref)
	if err != nil {
		log.Debug("DST: Unable to save", es_log.Error(err))
		return z.handleError(pathOrigSrc, pathSrc, pathOrigDst, pathDst, err)
	}

	z.report(metaSrc, metaDst)
	return nil
}

func (z *filesImpl) handleError(pathOrigSrc, pathSrc, pathOrigDst, pathDst mo_path.DropboxPath, apiErr error) error {
	errPrefix := dbx_util.ErrorSummary(apiErr)
	log := z.ctxSrc.Log().With(es_log.String("origSrc", pathOrigSrc.Path()), es_log.String("src", pathSrc.Path()), es_log.String("dst", pathDst.Path()), es_log.String("errorPrefix", errPrefix))
	switch {
	case strings.HasPrefix(errPrefix, "path/conflict"),
		strings.HasPrefix(errPrefix, "too_many_files"):
		log.Debug("Mirror descendants")
		return z.mirrorDescendants(pathOrigSrc, pathSrc, pathOrigDst, pathDst)

	case strings.Contains(errPrefix, "too_many_write_operations"):
		log.Debug("Wait for too many write", es_log.Duration("wait", z.pollInterval))
		time.Sleep(z.pollInterval)
		return z.mirrorCurrent(pathOrigSrc, pathSrc, pathOrigDst, pathDst)

	case strings.Contains(errPrefix, "not_found"):
		log.Debug("Can't copy file", es_log.String("src", pathSrc.Path()), es_log.String("dst", pathDst.Path()))
		return nil

	default:
		log.Debug("Unrecoverable error", es_log.String("errPrefix", errPrefix), es_log.Error(apiErr))
		return apiErr
	}
}

func (z *filesImpl) Mirror(pathSrc, pathDst mo_path.DropboxPath) (err error) {
	if pathSrc.LogicalPath() == "/" {
		return z.mirrorDescendants(pathSrc, pathSrc, pathDst, pathDst)
	} else {
		return z.mirrorCurrent(pathSrc, pathSrc, pathDst, pathDst)
	}
}
