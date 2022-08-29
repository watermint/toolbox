package uc_file_mirror

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_copyref"
	"github.com/watermint/toolbox/essentials/log/esl"
	"path/filepath"
	"strings"
	"time"
)

type Files interface {
	Mirror(pathSrc, pathDst mo_path.DropboxPath) (err error)
}

func New(ctxSrc, ctxDst dbx_client.Client) Files {
	return &filesImpl{
		ctxSrc:       ctxSrc,
		ctxDst:       ctxDst,
		pollInterval: 15 * 1000 * time.Millisecond,
	}
}

type filesImpl struct {
	ctxSrc       dbx_client.Client
	ctxDst       dbx_client.Client
	pollInterval time.Duration
}

func (z *filesImpl) report(metaSrc, metaDst mo_file.Entry) {
	z.ctxSrc.Log().Debug("Mirror complete", esl.String("src", metaSrc.PathDisplay()), esl.String("dst", metaDst.PathDisplay()))
}

func (z *filesImpl) dstPathRelToSrc(pathOrigSrc, pathSrc, pathOrigDst mo_path.DropboxPath) (relDst mo_path.DropboxPath, err error) {
	log := z.ctxSrc.Log().With(esl.String("origSrc", pathOrigSrc.Path()), esl.String("src", pathSrc.Path()), esl.String("origDst", pathOrigDst.Path()))

	pathDiff, err := filepath.Rel(pathOrigSrc.LogicalPath(), pathSrc.LogicalPath())
	if err != nil {
		log.Warn("Unable to calc rel path", esl.Error(err))
		return nil, err
	}
	pathDiff = filepath.ToSlash(pathDiff)

	if pathDiff == "." {
		return pathOrigDst, nil
	}

	if strings.HasPrefix(pathDiff, "..") {
		log.Warn("Invalid path", esl.String("diff", pathDiff))
		return nil, errors.New("invalid path")
	}

	relDst = pathOrigDst.ChildPath(pathDiff)

	log.Debug("DST: Path rel to src", esl.String("path", relDst.Path()))

	return relDst, nil
}

func (z *filesImpl) mirrorDescendants(pathOrigSrc, pathSrc, pathOrigDst, pathDst mo_path.DropboxPath) error {
	l := z.ctxSrc.Log().With(esl.String("origSrc", pathOrigSrc.Path()), esl.String("src", pathSrc.Path()), esl.String("dst", pathDst.Path()))
	l.Debug("Start mirroring descendants")

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
		switch {
		case ers.Path().IsNotFound():
			l.Debug("DST: Path not found. Proceed to mirror")

		case ers.Path().IsNotFolder():
			l.Debug("DST: Path is not a folder. Proceed to single file mirror")
			return z.mirrorCurrent(pathOrigSrc, pathSrc, pathOrigDst, pathDst)

		default:
			l.Debug("DST: Unable to list", esl.Error(err), esl.String("error_summary", ers.Summary()))
			return err
		}
		l.Debug("DST: Path not found. Proceed to mirror")
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
		l.Debug("Processing entry", esl.Int("index", i), esl.Int("numEntries", numEntriesSrc), esl.String("entryName", entrySrc.Name()))
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
					l.Warn("Unable to determine fileSrc server modified", esl.Any("fileSrc", fileSrc), esl.Error(err))
					skipped = append(skipped, &Skip{
						Name:   fileSrc.Name(),
						Reason: "src_time_err",
					})
					continue
				}
				dstTime, err := dbx_util.Parse(fileDst.ServerModified)
				if err != nil {
					l.Warn("Unable to determine fileDst server modified", esl.Any("fileDst", fileDst), esl.Error(err))
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
				l.Debug("Mirror failed", esl.String("fileSrc", pathFileSrc.Path()), esl.String("fileDst", pathFileDst.Path()), esl.Error(err))
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
					l.Debug("Mirror failed", esl.String("folderSrc", pathFolderSrc.Path()), esl.String("folderDst", pathFolderDst.Path()), esl.Error(err))
					// do not fail on file mirroring
				}
			} else {
				if err = z.mirrorCurrent(pathOrigSrc, pathFolderSrc, pathOrigDst, pathFolderDst); err != nil {
					l.Debug("Mirror failed", esl.String("folderSrc", pathFolderSrc.Path()), esl.String("folderDst", pathFolderDst.Path()), esl.Error(err))
					// do not fail on file mirroring
				}
			}
		}
	}

	l.Debug("Skipped:", esl.Any("files", skipped))

	return nil
}

func (z *filesImpl) mirrorCurrent(pathOrigSrc, pathSrc, pathOrigDst, pathDst mo_path.DropboxPath) (err error) {
	l := z.ctxSrc.Log().With(esl.String("origSrc", pathOrigSrc.Path()), esl.String("src", pathSrc.Path()), esl.String("dst", pathDst.Path()))
	l.Info("Start mirroring")

	scrSrc := sv_file_copyref.New(z.ctxSrc)
	metaSrc, ref, expires, err := scrSrc.Resolve(pathSrc)
	if err != nil {
		l.Debug("SRC: Unable to get copyRef", esl.Error(err))
		return err
	}
	l.Debug("SRC: CopyRef success", esl.String("ref", ref), esl.String("expires", expires))

	scrDst := sv_file_copyref.New(z.ctxDst)
	metaDst, err := scrDst.Save(pathDst, ref)
	if err != nil {
		l.Debug("DST: Unable to save", esl.Error(err))
		return z.handleError(pathOrigSrc, pathSrc, pathOrigDst, pathDst, err)
	}

	z.report(metaSrc, metaDst)
	return nil
}

func (z *filesImpl) handleError(pathOrigSrc, pathSrc, pathOrigDst, pathDst mo_path.DropboxPath, apiErr error) error {
	de := dbx_error.NewErrors(apiErr)
	log := z.ctxSrc.Log().With(esl.String("origSrc", pathOrigSrc.Path()), esl.String("src", pathSrc.Path()), esl.String("dst", pathDst.Path()), esl.String("errorSummary", de.Summary()))
	switch {
	case de.Path().IsConflictFolder(),
		de.IsTooManyFiles():
		log.Debug("Mirror descendants")
		return z.mirrorDescendants(pathOrigSrc, pathSrc, pathOrigDst, pathDst)

	case de.IsTooManyWriteOperations():
		log.Debug("Wait for too many write", esl.Duration("wait", z.pollInterval))
		time.Sleep(z.pollInterval)
		return z.mirrorCurrent(pathOrigSrc, pathSrc, pathOrigDst, pathDst)

	case de.Path().IsConflictFile():
		log.Info("Skipped: Conflict file found.", esl.String("srcPath", pathSrc.Path()), esl.String("dstPath", pathDst.Path()))
		return nil

	case de.Path().IsNotFound():
		log.Debug("Can't copy file", esl.String("src", pathSrc.Path()), esl.String("dst", pathDst.Path()))
		return nil

	default:
		log.Debug("Unrecoverable error", esl.String("errSummary", de.Summary()), esl.Error(apiErr))
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
