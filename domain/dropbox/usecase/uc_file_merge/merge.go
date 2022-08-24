package uc_file_merge

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_relocation"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_relocation"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"path/filepath"
	"strings"
	"time"
)

type MsgMerge struct {
	RemoveEmptyFolder    app_msg.Message
	RemoveDuplicatedFile app_msg.Message
	RemoveOldContent     app_msg.Message
	MoveFile             app_msg.Message
}

var (
	MMerge = app_msg.Apply(&MsgMerge{}).(*MsgMerge)
)

type Merge interface {
	Merge(from, to mo_path.DropboxPath, opts ...MergeOpt) error
}

func New(ctx dbx_client.Client, ctl app_control.Control) Merge {
	return &mergeImpl{
		ctx: ctx,
		ctl: ctl,
	}
}

type MergeOpt func(opts *MergeOpts) *MergeOpts
type MergeOpts struct {
	DryRun              bool
	WithinSameNamespace bool
	CleanEmptyFolder    bool
}

func DryRun() MergeOpt {
	return func(opts *MergeOpts) *MergeOpts {
		opts.DryRun = true
		return opts
	}
}
func WithinSameNamespace() MergeOpt {
	return func(opts *MergeOpts) *MergeOpts {
		opts.WithinSameNamespace = true
		return opts
	}
}
func ClearEmptyFolder() MergeOpt {
	return func(opts *MergeOpts) *MergeOpts {
		opts.CleanEmptyFolder = true
		return opts
	}
}

type mergeImpl struct {
	ctx       dbx_client.Client
	ctl       app_control.Control
	from      mo_path.DropboxPath
	fromEntry mo_file.Entry
	to        mo_path.DropboxPath
	toEntry   mo_file.Entry
	opts      *MergeOpts
}

func (z *mergeImpl) doOperation(msg app_msg.Message, op func() error) error {
	l := z.ctl.Log()
	dryRunIndicator := ""
	if z.opts.DryRun {
		dryRunIndicator = "DryRun: "
	}
	z.ctl.UI().Info(msg.With("DryRun", dryRunIndicator))
	if !z.opts.DryRun {
		return op()
	}
	l.Debug("Skip operation")
	return nil
}

func (z *mergeImpl) mergeFile(from, to *mo_file.File) error {
	l := z.ctl.Log().With(esl.String("from", from.PathDisplay()), esl.String("to", to.PathDisplay()))

	// remove same content hash file
	if from.ContentHash == to.ContentHash {
		l.Debug("Same content hash", esl.String("contentHash", from.ContentHash))
		m := MMerge.RemoveDuplicatedFile.With("FromPath", from.PathDisplay())
		return z.doOperation(m, func() error {
			entry, err := sv_file.NewFiles(z.ctx).Remove(mo_path.NewPathDisplay(from.PathDisplay()))
			if err != nil {
				l.Debug("Unable to remove file", esl.Error(err))
				return err
			}
			l.Debug("File removed", esl.Any("removedEntry", entry.Concrete()))
			return nil
		})
	}

	fromTs, err := time.Parse(time.RFC3339, from.ServerModified)
	if err != nil {
		l.Warn("Invalid time format", esl.Error(err), esl.String("from.ServerModified", from.ServerModified))
		return err
	}
	toTs, err := time.Parse(time.RFC3339, to.ServerModified)
	if err != nil {
		l.Warn("Invalid time format", esl.Error(err), esl.String("to.ServerModified", to.ServerModified))
		return err
	}

	l = l.With(esl.String("fromServerModified", from.ServerModified), esl.String("toServerModified", to.ServerModified))

	// remove old content
	if toTs.After(fromTs) {
		l.Debug("Remove the old file at from path")
		m := MMerge.RemoveOldContent.With("FromPath", from.PathDisplay())
		return z.doOperation(m, func() error {
			entry, err := sv_file.NewFiles(z.ctx).Remove(mo_path.NewPathDisplay(from.PathDisplay()))
			if err != nil {
				l.Debug("Unable to remove file", esl.Error(err))
				return err
			}
			l.Debug("File removed", esl.Any("removedEntry", entry.Concrete()))
			return nil
		})
	}

	// overwrite `from file` to `to path`
	fp := mo_path.NewPathDisplay(from.PathDisplay())
	tp := mo_path.NewPathDisplay(to.PathDisplay())
	m := MMerge.MoveFile.With("FromPath", from.PathDisplay()).With("ToPath", to.PathDisplay())
	return z.doOperation(m, func() error {
		entry, err := sv_file_relocation.New(z.ctx, sv_file_relocation.AutoRename(false)).Move(fp, tp)
		if err != nil {
			l.Debug("Unable to move file", esl.Error(err))
			return err
		}
		l.Debug("File moved", esl.Any("movedEntry", entry.Concrete()))

		return nil
	})
}

func (z *mergeImpl) moveFile(from *mo_file.File) error {
	l := z.ctl.Log().With(esl.Any("from", from.Concrete()))
	l.Debug("Move file")
	p, err := filepath.Rel(z.fromEntry.PathLower(), from.PathLower())
	if err != nil {
		l.Warn("Unable to calc relative path", esl.Error(err))
		// TODO: do reporting
		return err
	}
	if strings.HasPrefix(p, "..") {
		l.Warn("Invalid relative path", esl.String("rel", p))
		// TODO; do reporting
		return errors.New("invalid relative path")
	}

	fp := mo_path.NewPathDisplay(from.PathDisplay())
	tp := z.to.ChildPath(filepath.Dir(p), from.Name())
	l.Debug("move file")
	m := MMerge.MoveFile.With("FromPath", from.PathDisplay()).With("ToPath", tp.Path())
	return z.doOperation(m, func() error {
		entry, err := sv_file_relocation.New(z.ctx).Move(fp, tp)
		if err != nil {
			l.Debug("Unable to move", esl.Error(err))
			return err
		}
		l.Debug("file moved", esl.Any("entry", entry.Concrete()))
		return nil
	})
}

func (z *mergeImpl) moveFolder(from *mo_file.Folder) error {
	l := z.ctl.Log().With(esl.Any("from", from.Concrete()))
	l.Debug("Move folder")
	p, err := filepath.Rel(z.fromEntry.PathLower(), from.PathLower())
	if err != nil {
		l.Warn("Unable to calc relative path", esl.Error(err))
		// TODO: do reporting
		return err
	}
	if strings.HasPrefix(p, "..") {
		l.Warn("Invalid relative path", esl.String("rel", p))
		// TODO; do reporting
		return errors.New("invalid relative path")
	}
	fp := mo_path.NewPathDisplay(from.PathDisplay())
	tp := z.to.ChildPath(filepath.Dir(p), from.Name())

	l = l.With(esl.String("toPath", tp.Path()))

	// move
	m := MMerge.MoveFile.With("FromPath", from.PathDisplay()).With("ToPath", tp.Path())
	return z.doOperation(m, func() error {
		l.Debug("move folder")
		return uc_file_relocation.New(z.ctx).Move(fp, tp)
	})
}

func (z *mergeImpl) validatePaths(from, to mo_file.Entry) error {
	l := z.ctl.Log()

	ff, e := from.Folder()
	if !e {
		return errors.New("`from` path is not a folder")
	}
	tf, e := to.Folder()
	if !e {
		return errors.New("`to` path is not a folder")
	}
	l = l.With(esl.Any("from", ff.Concrete()), esl.Any("to", tf.Concrete()))

	if !z.opts.WithinSameNamespace {
		l.Debug("Skip validate namespace")
		return nil
	}

	if ff.EntrySharedFolderId != tf.EntrySharedFolderId ||
		ff.EntryParentSharedFolderId != tf.EntryParentSharedFolderId {
		l.Debug("Different namespace")
		return errors.New("`from` path and `to` path are different namespace")
	}

	l.Debug("Go for merge")
	return nil
}

// Merge at relative path from `z.from`.
func (z *mergeImpl) merge(path string) error {
	fromFiles := make(map[string]*mo_file.File)
	fromFolders := make(map[string]*mo_file.Folder)
	toFiles := make(map[string]*mo_file.File)
	toFolders := make(map[string]*mo_file.Folder)

	l := z.ctl.Log().With(esl.String("path", path))
	l.Debug("merge")

	fromPath := z.from.ChildPath(path)
	toPath := z.to.ChildPath(path)

	// Scan from
	{
		entries, err := sv_file.NewFiles(z.ctx).List(fromPath)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			if f, e := entry.File(); e {
				fromFiles[strings.ToLower(f.Name())] = f
			}
			if f, e := entry.Folder(); e {
				fromFolders[strings.ToLower(f.Name())] = f
			}
		}
	}

	// Scan to
	{
		entries, err := sv_file.NewFiles(z.ctx).List(toPath)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			if f, e := entry.File(); e {
				toFiles[strings.ToLower(f.Name())] = f
			}
			if f, e := entry.Folder(); e {
				toFolders[strings.ToLower(f.Name())] = f
			}
		}
	}

	var lastErr error

	// move or merge files
	for ffn, ff := range fromFiles {
		ll := l.With(esl.String("from", ff.EntryPathDisplay))
		if tf, e := toFiles[ffn]; e {
			ll = ll.With(esl.String("to", tf.EntryPathDisplay))
			if err := z.mergeFile(ff, tf); err != nil {
				ll.Debug("Unable to merge", esl.Error(err))
				lastErr = err
			} else {
				ll.Debug("File merged")
			}
		} else {
			if err := z.moveFile(ff); err != nil {
				ll.Debug("Unable to move", esl.Error(err))
				lastErr = err
			} else {
				ll.Debug("File moved")
			}
		}
	}

	// move or merge folders
	for ffn, ff := range fromFolders {
		ll := l.With(esl.String("from", ff.EntryPathDisplay))
		if _, e := toFolders[ffn]; e {
			ll.Debug("Proceed into descendants")
			p, err := filepath.Rel(z.fromEntry.PathLower(), ff.PathLower())
			if err != nil {
				ll.Warn("Unable to calc relative path", esl.Error(err))
				// TODO: do reporting
				continue
			}
			if strings.HasPrefix(p, "..") {
				ll.Warn("Invalid relative path", esl.String("rel", p))
				// TODO; do reporting
				continue
			}

			if err := z.merge(p); err != nil {
				ll.Debug("One or more error in descendant", esl.Error(err))
				lastErr = err
			}

		} else {
			if err := z.moveFolder(ff); err != nil {
				ll.Debug("Unable to move", esl.Error(err))
				lastErr = err
			} else {
				ll.Debug("Folder moved")
			}
		}
	}

	// remove if the folder empty
	if z.opts.CleanEmptyFolder {
		entries, err := sv_file.NewFiles(z.ctx).List(fromPath)
		if err != nil {
			l.Debug("Unable to list", esl.Error(err))
			return lastErr
		}
		if len(entries) < 1 {
			l.Debug("Try clean up folder")
			m := MMerge.RemoveEmptyFolder.With("FromPath", fromPath.Path())
			z.doOperation(m, func() error {
				entry, err := sv_file.NewFiles(z.ctx).Remove(fromPath)
				if err != nil {
					l.Debug("Unable to remove folder", esl.Error(err))
					return err
				}
				l.Debug("Removed", esl.Any("removed", entry.Concrete()))
				return nil
			})
		} else {
			l.Debug("Remaining entries", esl.Int("entries", len(entries)))
		}
	}

	return lastErr
}

func (z *mergeImpl) Merge(from, to mo_path.DropboxPath, opts ...MergeOpt) (err error) {
	z.opts = &MergeOpts{}
	for _, o := range opts {
		o(z.opts)
	}
	z.from = from
	z.to = to
	l := z.ctl.Log().With(esl.String("from", from.Path()), esl.String("to", to.Path()))

	z.fromEntry, err = sv_file.NewFiles(z.ctx).Resolve(from)
	if err != nil {
		l.Debug("Unable to resolve fromFolder", esl.Error(err))
		return err
	}

	z.toEntry, err = sv_file.NewFiles(z.ctx).Resolve(to)
	if err != nil {
		l.Debug("Unable to resolve toFolder", esl.Error(err))
		return err
	}

	if err := z.validatePaths(z.fromEntry, z.toEntry); err != nil {
		return err
	}

	return z.merge("")
}
