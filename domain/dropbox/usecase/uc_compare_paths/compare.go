package uc_compare_paths

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"path/filepath"
	"strings"
)

type MsgCompare struct {
	ProgressScanFolder app_msg.Message
}

var (
	MCompare = app_msg.Apply(&MsgCompare{}).(*MsgCompare)
)

type Compare interface {
	Diff(leftPath mo_path.DropboxPath, rightPath mo_path.DropboxPath, onDiff func(diff mo_file_diff.Diff) error) (diffCount int, err error)
}

func New(left, right dbx_context.Context, ui app_ui.UI, opts ...CompareOpt) Compare {
	co := &CompareOpts{}
	for _, o := range opts {
		o(co)
	}

	return &compareImpl{
		ctxLeft:  left,
		ctxRight: right,
		opts:     co,
		ui:       ui,
	}
}

type CompareOpt func(opt *CompareOpts) *CompareOpts

type CompareOpts struct {
}

type compareImpl struct {
	ctxLeft  dbx_context.Context
	ctxRight dbx_context.Context
	opts     *CompareOpts
	ui       app_ui.UI
}

func (z *compareImpl) cmpLevel(left, right mo_path.DropboxPath, path string, onDiff func(diff mo_file_diff.Diff) error) (diffCount int, err error) {
	leftFiles := make(map[string]*mo_file.File)
	leftFolders := make(map[string]*mo_file.Folder)
	rightFiles := make(map[string]*mo_file.File)
	rightFolders := make(map[string]*mo_file.Folder)

	l := z.ctxLeft.Log().With(es_log.String("path", path))

	z.ui.Progress(MCompare.ProgressScanFolder.With("Path", path))

	// Scan left
	{
		l.Debug("Scan left")
		leftPath := left.ChildPath(path)
		entries, err := sv_file.NewFiles(z.ctxLeft).List(leftPath)
		if err != nil {
			l.Debug("unable to list left path", es_log.Error(err))
			return 0, err
		}
		for _, entry := range entries {
			if f, e := entry.File(); e {
				leftFiles[strings.ToLower(f.Name())] = f
			}
			if f, e := entry.Folder(); e {
				leftFolders[strings.ToLower(f.Name())] = f
			}
		}
	}

	// Scan right
	{
		l.Debug("Scan right")
		rightPath := right.ChildPath(path)
		entries, err := sv_file.NewFiles(z.ctxRight).List(rightPath)
		if err != nil {
			l.Debug("unable to list right path", es_log.Error(err))
			return 0, err
		}
		for _, entry := range entries {
			if f, e := entry.File(); e {
				rightFiles[strings.ToLower(f.Name())] = f
			}
			if f, e := entry.Folder(); e {
				rightFolders[strings.ToLower(f.Name())] = f
			}
		}
	}

	// compare files left to right
	l.Debug("Compare files left to right")
	for lfn, lf := range leftFiles {
		if rf, e := rightFiles[lfn]; e {
			if lf.ContentHash != rf.ContentHash {
				diff := mo_file_diff.Diff{
					DiffType:  mo_file_diff.DiffFileContent,
					LeftPath:  lf.PathDisplay(),
					LeftKind:  "file",
					LeftSize:  &lf.Size,
					LeftHash:  lf.ContentHash,
					RightPath: rf.PathDisplay(),
					RightKind: "file",
					RightSize: &rf.Size,
					RightHash: rf.ContentHash,
				}
				diffCount++
				if err := onDiff(diff); err != nil {
					return diffCount, err
				}
			}
		} else {
			diff := mo_file_diff.Diff{
				DiffType: mo_file_diff.DiffFileMissingRight,
				LeftPath: lf.PathDisplay(),
				LeftKind: "file",
				LeftSize: &lf.Size,
				LeftHash: lf.ContentHash,
			}
			diffCount++
			if err := onDiff(diff); err != nil {
				l.Debug("onDiff returned an error", es_log.Error(err))
				return diffCount, err
			}
		}
	}

	// compare files right to left
	l.Debug("Compare files right to left")
	for rfn, rf := range rightFiles {
		if _, e := leftFiles[rfn]; !e {
			diff := mo_file_diff.Diff{
				DiffType:  mo_file_diff.DiffFileMissingLeft,
				RightPath: rf.PathDisplay(),
				RightKind: "file",
				RightSize: &rf.Size,
				RightHash: rf.ContentHash,
			}
			diffCount++
			if err := onDiff(diff); err != nil {
				l.Debug("onDiff returned an error", es_log.Error(err))
				return diffCount, err
			}
		}
	}

	// compare folders left to right
	l.Debug("Compare folders left to right")
	for lfn, lf := range leftFolders {
		if _, e := rightFolders[lfn]; e {
			// proceed to descendants
			lp := strings.ToLower(left.Path())
			if lp == "" {
				lp = "/"
			}
			pd, err := filepath.Rel(lp, lf.PathLower())
			if err != nil {
				l.Warn("unable to calculate relative path", es_log.String("leftPathBase", lp), es_log.String("leftPath", lf.PathLower()), es_log.Error(err))
				continue
			}
			if strings.HasPrefix(pd, "..") {
				l.Error("invalid relative path", es_log.String("pd", pd), es_log.String("zLeftPath", left.Path()), es_log.String("lfPathLower", lf.PathLower()))
				continue
			}
			l.Debug("Proceed into descendants", es_log.String("pathDescendants", pd))
			dc, err := z.cmpLevel(left, right, pd, onDiff)
			if err != nil {
				return dc, err
			}
			diffCount += dc
		} else {
			diff := mo_file_diff.Diff{
				DiffType: mo_file_diff.DiffFolderMissingRight,
				LeftPath: lf.PathDisplay(),
				LeftKind: "folder",
			}
			diffCount++
			if err := onDiff(diff); err != nil {
				l.Debug("onDiff returned an error", es_log.Error(err))
				return diffCount, err
			}
		}
	}

	// compare folders right to left
	l.Debug("Compare folders right to left")
	for rfn, rf := range rightFolders {
		if _, e := leftFolders[rfn]; !e {
			diff := mo_file_diff.Diff{
				DiffType:  mo_file_diff.DiffFolderMissingLeft,
				RightPath: rf.PathDisplay(),
				RightKind: "folder",
			}
			diffCount++
			if err := onDiff(diff); err != nil {
				l.Debug("onDiff returned an error", es_log.Error(err))
				return diffCount, err
			}
		}
	}

	l.Debug("Completed", es_log.Int("diffCount", diffCount))
	return diffCount, nil
}

func (z *compareImpl) Diff(leftPath mo_path.DropboxPath, rightPath mo_path.DropboxPath, onDiff func(diff mo_file_diff.Diff) error) (diffCount int, err error) {
	return z.cmpLevel(leftPath, rightPath, "", onDiff)
}
