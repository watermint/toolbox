package uc_file_merge

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_file_relocation"
	"github.com/watermint/toolbox/domain/usecase/uc_file_relocation"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
	"time"
)

type Merge interface {
	Merge(from, to mo_path.Path, opts ...MergeOpt) error
}

func New(ctx api_context.Context, k app_kitchen.Kitchen) Merge {
	return &mergeImpl{
		ctx: ctx,
		k:   k,
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
	ctx       api_context.Context
	k         app_kitchen.Kitchen
	from      mo_path.Path
	fromEntry mo_file.Entry
	to        mo_path.Path
	toEntry   mo_file.Entry
	opts      *MergeOpts
}

func (z *mergeImpl) doOperation(msg app_msg.Message, op func() error) error {
	l := z.k.Log()
	msgParam := make([]app_msg.P, 0)
	msgParam = append(msgParam, msg.Params()...)
	dryRunIndicator := ""
	if z.opts.DryRun {
		dryRunIndicator = "DryRun: "
	}
	msgParam = append(msgParam, app_msg.P{
		"DryRun": dryRunIndicator,
	})

	z.k.UI().Info(msg.Key(), msgParam...)
	if !z.opts.DryRun {
		return op()
	}
	l.Debug("Skip operation")
	return nil
}

func (z *mergeImpl) mergeFile(from, to *mo_file.File) error {
	l := z.k.Log().With(zap.String("from", from.PathDisplay()), zap.String("to", to.PathDisplay()))

	// remove same content hash file
	if from.ContentHash == to.ContentHash {
		l.Debug("Same content hash", zap.String("contentHash", from.ContentHash))
		m := app_msg.M("usecase.uc_file_merge.remove_duplicated_file",
			app_msg.P{
				"FromPath": from.PathDisplay(),
			})
		return z.doOperation(m, func() error {
			entry, err := sv_file.NewFiles(z.ctx).Remove(mo_path.NewPathDisplay(from.PathDisplay()))
			if err != nil {
				l.Debug("Unable to remove file", zap.Error(err))
				return err
			}
			l.Debug("File removed", zap.Any("removedEntry", entry.Concrete()))
			return nil
		})
	}

	fromTs, err := time.Parse(time.RFC3339, from.ServerModified)
	if err != nil {
		l.Warn("Invalid time format", zap.Error(err), zap.String("from.ServerModified", from.ServerModified))
		return err
	}
	toTs, err := time.Parse(time.RFC3339, to.ServerModified)
	if err != nil {
		l.Warn("Invalid time format", zap.Error(err), zap.String("to.ServerModified", to.ServerModified))
		return err
	}

	l = l.With(zap.String("fromServerModified", from.ServerModified), zap.String("toServerModified", to.ServerModified))

	// remove old content
	if toTs.After(fromTs) {
		l.Debug("Remove the old file at from path")
		m := app_msg.M("usecase.uc_file_merge.remove_old_content",
			app_msg.P{
				"FromPath": from.PathDisplay(),
			})
		return z.doOperation(m, func() error {
			entry, err := sv_file.NewFiles(z.ctx).Remove(mo_path.NewPathDisplay(from.PathDisplay()))
			if err != nil {
				l.Debug("Unable to remove file", zap.Error(err))
				return err
			}
			l.Debug("File removed", zap.Any("removedEntry", entry.Concrete()))
			return nil
		})
	}

	// overwrite `from file` to `to path`
	fp := mo_path.NewPathDisplay(from.PathDisplay())
	tp := mo_path.NewPathDisplay(to.PathDisplay())
	m := app_msg.M("usecase.uc_file_merge.move_file",
		app_msg.P{
			"FromPath": from.PathDisplay(),
			"ToPath":   to.PathDisplay(),
		},
	)
	return z.doOperation(m, func() error {
		entry, err := sv_file_relocation.New(z.ctx, sv_file_relocation.AutoRename(false)).Move(fp, tp)
		if err != nil {
			l.Debug("Unable to move file", zap.Error(err))
			return err
		}
		l.Debug("File moved", zap.Any("movedEntry", entry.Concrete()))

		return nil
	})
}

func (z *mergeImpl) moveFile(from *mo_file.File) error {
	l := z.k.Log().With(zap.Any("from", from.Concrete()))
	l.Debug("Move file")
	p, err := filepath.Rel(z.fromEntry.PathLower(), from.PathLower())
	if err != nil {
		l.Warn("Unable to calc relative path", zap.Error(err))
		// TODO: do reporting
		return err
	}
	if strings.HasPrefix(p, "..") {
		l.Warn("Invalid relative path", zap.String("rel", p))
		// TODO; do reporting
		return errors.New("invalid relative path")
	}

	fp := mo_path.NewPathDisplay(from.PathDisplay())
	tp := z.to.ChildPath(filepath.Join(filepath.Dir(p), from.Name()))
	l.Debug("move file")
	m := app_msg.M("usecase.uc_file_merge.move_file",
		app_msg.P{
			"FromPath": from.PathDisplay(),
			"ToPath":   tp.Path(),
		},
	)
	return z.doOperation(m, func() error {
		entry, err := sv_file_relocation.New(z.ctx).Move(fp, tp)
		if err != nil {
			l.Debug("Unable to move", zap.Error(err))
			return err
		}
		l.Debug("file moved", zap.Any("entry", entry.Concrete()))
		return nil
	})
}

func (z *mergeImpl) moveFolder(from *mo_file.Folder) error {
	l := z.k.Log().With(zap.Any("from", from.Concrete()))
	l.Debug("Move folder")
	p, err := filepath.Rel(z.fromEntry.PathLower(), from.PathLower())
	if err != nil {
		l.Warn("Unable to calc relative path", zap.Error(err))
		// TODO: do reporting
		return err
	}
	if strings.HasPrefix(p, "..") {
		l.Warn("Invalid relative path", zap.String("rel", p))
		// TODO; do reporting
		return errors.New("invalid relative path")
	}
	fp := mo_path.NewPathDisplay(from.PathDisplay())
	tp := z.to.ChildPath(filepath.Join(filepath.Dir(p), from.Name()))

	l = l.With(zap.String("toPath", tp.Path()))

	// move
	m := app_msg.M("usecase.uc_file_merge.move_file",
		app_msg.P{
			"FromPath": from.PathDisplay(),
			"ToPath":   tp.Path(),
		},
	)
	return z.doOperation(m, func() error {
		l.Debug("move folder")
		return uc_file_relocation.New(z.ctx).Move(fp, tp)
	})
}

func (z *mergeImpl) validatePaths(from, to mo_file.Entry) error {
	l := z.k.Log()

	ff, e := from.Folder()
	if !e {
		return errors.New("`from` path is not a folder")
	}
	tf, e := to.Folder()
	if !e {
		return errors.New("`to` path is not a folder")
	}
	l = l.With(zap.Any("from", ff.Concrete()), zap.Any("to", tf.Concrete()))

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

	l := z.k.Log().With(zap.String("path", path))
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
		ll := l.With(zap.String("from", ff.EntryPathDisplay))
		if tf, e := toFiles[ffn]; e {
			ll = ll.With(zap.String("to", tf.EntryPathDisplay))
			if err := z.mergeFile(ff, tf); err != nil {
				ll.Debug("Unable to merge", zap.Error(err))
				lastErr = err
			} else {
				ll.Debug("File merged")
			}
		} else {
			if err := z.moveFile(ff); err != nil {
				ll.Debug("Unable to move", zap.Error(err))
				lastErr = err
			} else {
				ll.Debug("File moved")
			}
		}
	}

	// move or merge folders
	for ffn, ff := range fromFolders {
		ll := l.With(zap.String("from", ff.EntryPathDisplay))
		if _, e := toFolders[ffn]; e {
			ll.Debug("Proceed into descendants")
			p, err := filepath.Rel(z.fromEntry.PathLower(), ff.PathLower())
			if err != nil {
				ll.Warn("Unable to calc relative path", zap.Error(err))
				// TODO: do reporting
				continue
			}
			if strings.HasPrefix(p, "..") {
				ll.Warn("Invalid relative path", zap.String("rel", p))
				// TODO; do reporting
				continue
			}

			if err := z.merge(p); err != nil {
				ll.Debug("One or more error in descendant", zap.Error(err))
				lastErr = err
			}

		} else {
			if err := z.moveFolder(ff); err != nil {
				ll.Debug("Unable to move", zap.Error(err))
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
			l.Debug("Unable to list", zap.Error(err))
			return lastErr
		}
		if len(entries) < 1 {
			l.Debug("Try clean up folder")
			m := app_msg.M("usecase.uc_file_merge.remove_empty_folder",
				app_msg.P{
					"FromPath": fromPath.Path(),
				},
			)
			z.doOperation(m, func() error {
				entry, err := sv_file.NewFiles(z.ctx).Remove(fromPath)
				if err != nil {
					l.Debug("Unable to remove folder", zap.Error(err))
					return err
				}
				l.Debug("Removed", zap.Any("removed", entry.Concrete()))
				return nil
			})
		} else {
			l.Debug("Remaining entries", zap.Int("entries", len(entries)))
		}
	}

	return lastErr
}

func (z *mergeImpl) Merge(from, to mo_path.Path, opts ...MergeOpt) (err error) {
	z.opts = &MergeOpts{}
	for _, o := range opts {
		o(z.opts)
	}
	z.from = from
	z.to = to
	l := z.k.Log().With(zap.String("from", from.Path()), zap.String("to", to.Path()))

	z.fromEntry, err = sv_file.NewFiles(z.ctx).Resolve(from)
	if err != nil {
		l.Debug("Unable to resolve fromFolder", zap.Error(err))
		return err
	}

	z.toEntry, err = sv_file.NewFiles(z.ctx).Resolve(to)
	if err != nil {
		l.Debug("Unable to resolve toFolder", zap.Error(err))
		return err
	}

	if err := z.validatePaths(z.fromEntry, z.toEntry); err != nil {
		return err
	}

	return z.merge("")
}
