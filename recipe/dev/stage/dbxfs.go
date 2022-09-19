package stage

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Dbxfs struct {
	rc_recipe.RemarkSecret
	Peer dbx_conn.ConnScopedIndividual
	Path mo_path.DropboxPath
}

func (z *Dbxfs) Preset() {
	z.Peer.SetScopes(dbx_auth.ScopeFilesContentRead)
}

func (z *Dbxfs) compareEntry(l esl.Logger, base, cached es_filesystem.FileSystem, path es_filesystem.Path) error {
	l = l.With(esl.String("path", path.Path()))
	baseEntry, baseErr := base.Info(path)
	cachedEntry, cachedErr := cached.Info(path)

	if baseErr == nil && cachedErr == nil {
		if es_filesystem.CompareEntry(baseEntry, cachedEntry) {
			l.Debug("Entry matched")
			return nil
		} else {
			l.Warn("Entry mismatch found",
				esl.Any("base", baseEntry.AsData()),
				esl.Any("cached", cachedEntry.AsData()))
			return errors.New("entry mismatch")
		}
	} else if baseErr != nil && cachedErr != nil {
		if baseErr.IsPathNotFound() && cachedErr.IsPathNotFound() {
			l.Debug("Both not found")
			return nil
		}
		// fall though
	}

	l.Warn("Error mismatch found", esl.Any("base", baseErr), esl.Any("cached", cachedErr))
	return errors.New("error mismatch")
}

func (z *Dbxfs) compare(l esl.Logger, base, cached es_filesystem.FileSystem, path es_filesystem.Path) error {
	if err := z.compareEntry(l, base, cached, path); err != nil {
		return err
	}

	baseEntries, baseErr := base.List(path)
	cachedEntries, cachedErr := cached.List(path)

	if baseErr == nil && cachedErr == nil {
		if len(baseEntries) != len(cachedEntries) {
			l.Debug("Num entries mismatch", esl.Any("base", baseEntries), esl.Any("cached", cachedEntries))
			return errors.New("num entries mismatch")
		}

		baseEntryMap := make(map[string]es_filesystem.Entry)
		for _, e := range baseEntries {
			baseEntryMap[e.Path().Path()] = e
		}
		cacheEntryMap := make(map[string]es_filesystem.Entry)
		for _, e := range cachedEntries {
			cacheEntryMap[e.Path().Path()] = e
		}

		// base -> cache
		for k, b := range baseEntryMap {
			if c, ok := cacheEntryMap[k]; ok {
				if es_filesystem.CompareEntry(b, c) {
					l.Debug("Entry matched")
				} else {
					l.Warn("Entry mismatch found",
						esl.Any("base", b.AsData()),
						esl.Any("cached", c.AsData()))
					return errors.New("entry mismatch")
				}
			}
		}

		// cache -> base
		for k, c := range cacheEntryMap {
			if b, ok := baseEntryMap[k]; ok {
				if es_filesystem.CompareEntry(b, c) {
					l.Debug("Entry matched")
				} else {
					l.Warn("Entry mismatch found",
						esl.Any("base", b.AsData()),
						esl.Any("cached", c.AsData()))
					return errors.New("entry mismatch")
				}
			}
		}

		l.Debug("Matched")
		return nil

	} else if baseErr != nil && cachedErr != nil {
		if baseErr.IsPathNotFound() && cachedErr.IsPathNotFound() {
			l.Debug("Both not found")
			return nil
		}
		// fall though
	}

	l.Debug("Error mismatch found", esl.Any("base", baseErr), esl.Any("cached", cachedErr))
	return errors.New("error mismatch")
}

func (z *Dbxfs) Exec(c app_control.Control) error {
	base := dbx_fs.NewFileSystem(z.Peer.Client())
	cached, err := dbx_fs.NewPreScanFileSystem(c, z.Peer.Client(), z.Path)
	if err != nil {
		return err
	}

	return z.compare(c.Log(), base, cached, dbx_fs.NewPath("", z.Path))
}

func (z *Dbxfs) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Dbxfs{}, func(r rc_recipe.Recipe) {
		m := r.(*Dbxfs)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("dbxfs")
	})
}
