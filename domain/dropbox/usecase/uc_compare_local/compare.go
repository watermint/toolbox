package uc_compare_local

import (
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"io/ioutil"
	"os"
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
	Diff(localPath mo_path2.FileSystemPath, dropboxPath mo_path.DropboxPath, onDiff func(diff mo_file_diff.Diff) error) (diffCount int, err error)
}

type CompareOpt func(o *CompareOpts) *CompareOpts
type CompareOpts struct {
	ForceCalcLocalHash bool
}

func New(ctx dbx_context.Context, ui app_ui.UI, opts ...CompareOpt) Compare {
	co := &CompareOpts{}
	for _, o := range opts {
		o(co)
	}
	return &compareImpl{
		ctx:  ctx,
		opts: co,
		ui:   ui,
	}
}

type compareImpl struct {
	ctx  dbx_context.Context
	ui   app_ui.UI
	opts *CompareOpts
}

func (z *compareImpl) cmpLevel(local mo_path2.FileSystemPath, dropbox mo_path.DropboxPath, path string, onDiff func(diff mo_file_diff.Diff) error) (diffCount int, err error) {
	localFiles := make(map[string]os.FileInfo)
	localFolders := make(map[string]os.FileInfo)
	dropboxFiles := make(map[string]*mo_file.File)
	dropboxFolders := make(map[string]*mo_file.Folder)

	l := z.ctx.Log().With(
		es_log.String("local", local.Path()),
		es_log.String("dropbox", dropbox.Path()),
		es_log.String("path", path))

	localPath := func(info os.FileInfo) string {
		if path == "" {
			return filepath.Join(local.Path(), info.Name())
		} else {
			return filepath.Join(local.Path(), path, info.Name())
		}
	}
	relPath := func(info os.FileInfo) string {
		if path == "" {
			return info.Name()
		} else {
			return filepath.Join(path, info.Name())
		}
	}

	z.ui.Progress(MCompare.ProgressScanFolder.With("Path", path))

	// Scan local
	{
		l.Debug("Scan local")
		localPath := filepath.Join(local.Path(), path)
		entries, err := ioutil.ReadDir(localPath)
		if err != nil {
			l.Debug("Unable to read dir")
			return 0, err
		}
		for _, entry := range entries {
			name := strings.ToLower(entry.Name())
			if entry.IsDir() {
				localFolders[name] = entry
			} else {
				localFiles[name] = entry
			}
		}
	}

	// Scan dropbox
	{
		l.Debug("Scan dropbox")
		dropboxPath := dropbox.ChildPath(filepath.ToSlash(path))
		entries, err := sv_file.NewFiles(z.ctx).List(dropboxPath)
		if err != nil {
			l.Debug("unable to list dropbox path", es_log.Error(err))
			return 0, err
		}
		for _, entry := range entries {
			name := strings.ToLower(entry.Name())
			if f, e := entry.File(); e {
				dropboxFiles[name] = f
			}
			if f, e := entry.Folder(); e {
				dropboxFolders[name] = f
			}
		}
	}

	// Compare files local to dropbox
	l.Debug("Compare files local to dropbox")
	for name, lf := range localFiles {
		lfp := localPath(lf)
		calcHash := func(p string) string {
			hash, err := dbx_util.ContentHash(p)
			if err != nil {
				hash = "<failed to calculate content hash>"
				l.Debug("Unable to calculate hash", es_log.String("localPath", p), es_log.Error(err))
			}
			return hash
		}

		if df, e := dropboxFiles[name]; e {
			hash := ""
			if z.opts.ForceCalcLocalHash || (lf.Size() == df.Size) {
				hash = calcHash(lfp)
			}

			switch {
			case lf.Size() != df.Size, hash != df.ContentHash:
				lsz := lf.Size()
				diff := mo_file_diff.Diff{
					DiffType:  mo_file_diff.DiffFileContent,
					LeftPath:  lfp,
					LeftKind:  "file",
					LeftSize:  &lsz,
					LeftHash:  hash,
					RightPath: df.PathDisplay(),
					RightKind: "file",
					RightSize: &df.Size,
					RightHash: df.ContentHash,
				}
				diffCount++
				if err := onDiff(diff); err != nil {
					l.Debug("onDiff returned an error", es_log.Error(err))
					return diffCount, err
				}
			}
		} else {
			lsz := lf.Size()
			hash := ""
			if z.opts.ForceCalcLocalHash {
				hash = calcHash(lfp)
			}
			dt := mo_file_diff.DiffFileMissingRight
			if dbx_util.IsFileNameIgnored(lfp) {
				dt = mo_file_diff.DiffSkipped
			}

			diff := mo_file_diff.Diff{
				DiffType: dt,
				LeftPath: lfp,
				LeftKind: "file",
				LeftSize: &lsz,
				LeftHash: hash,
			}
			diffCount++
			if err := onDiff(diff); err != nil {
				l.Debug("onDiff returned an error", es_log.Error(err))
				return diffCount, err
			}
		}
	}

	// Compare files dropbox to local
	l.Debug("Compare files dropbox to local")
	for name, df := range dropboxFiles {
		if _, e := localFiles[name]; !e {
			diff := mo_file_diff.Diff{
				DiffType:  mo_file_diff.DiffFileMissingLeft,
				RightPath: df.PathDisplay(),
				RightKind: "file",
				RightSize: &df.Size,
				RightHash: df.ContentHash,
			}
			diffCount++
			if err := onDiff(diff); err != nil {
				return diffCount, err
			}
		}
	}

	// Compare folders local to dropbox
	l.Debug("Compare folders local to dropbox")
	for name, lf := range localFolders {
		lfp := localPath(lf)
		if _, e := dropboxFolders[name]; e {
			// proceed to descendants
			rp := relPath(lf)
			l.Debug("Proceed to descendants", es_log.String("rp", rp))
			dc, err := z.cmpLevel(local, dropbox, rp, onDiff)
			if err != nil {
				l.Debug("Descendant returned an error", es_log.Error(err))
				return diffCount + dc, err
			}
			diffCount += dc
		} else {
			dt := mo_file_diff.DiffFolderMissingRight
			if strings.ToLower(name) == ".dropbox.cache" {
				dt = mo_file_diff.DiffSkipped
			}
			diff := mo_file_diff.Diff{
				DiffType: dt,
				LeftPath: lfp,
				LeftKind: "folder",
			}
			diffCount++
			if err := onDiff(diff); err != nil {
				l.Debug("onDiff returned an error", es_log.Error(err))
				return diffCount, err
			}
		}
	}

	// Compare folders dropbox to local
	for name, df := range dropboxFolders {
		if _, e := localFolders[name]; !e {
			diff := mo_file_diff.Diff{
				DiffType:  mo_file_diff.DiffFolderMissingLeft,
				RightPath: df.PathDisplay(),
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

func (z *compareImpl) Diff(localPath mo_path2.FileSystemPath, dropboxPath mo_path.DropboxPath, onDiff func(diff mo_file_diff.Diff) error) (diffCount int, err error) {
	return z.cmpLevel(localPath, dropboxPath, "", onDiff)
}
